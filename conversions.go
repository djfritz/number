// Copyright 2025 David Fritz. All rights reserved.
// This software may be modified and distributed under the terms of the BSD
// 2-clause license. See the LICENSE file for details.

package number

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Printing specific constants
const (
	asciiOffset           = 0x30 // offset to add to bytes to get their ASCII representation
	sensibleSize          = 40   // sensible number of digits to print before engaging scientific notation for the %v verb
	defaultPrintPrecision = DefaultPrecision
)

var (
	ErrInvalidCharacter = errors.New("invalid character")
)

// Return the string form of the real number in scientific notation.
func (x *Real) String() string {
	return fmt.Sprintf("%e", x)
}

// Format implements [fmt.Formatter]. It acccepts 'd', 'e', 'f', and 'v' verbs.
// The 'd' verb will return an integer with trailing zeros. 'e' is the same as
// String(), and returns the value in scientific notation. 'f' returns a
// floating point value, and accepts precision modifiers. 'v' will choose an
// appropriate form based on the magnitude of the number.
func (x *Real) Format(s fmt.State, verb rune) {
	printable := x.Copy()
	p, precisionSet := s.Precision()
	if !precisionSet {
		p = defaultPrintPrecision
	}

	var o bytes.Buffer
	if printable.negative {
		o.WriteString("-")
	}

	if x.form == FormInf {
		o.WriteString("∞")
		s.Write(o.Bytes())
		return
	} else if x.form == FormNaN {
		o.WriteString("NaN")
		s.Write(o.Bytes())
		return
	}

	switch verb {
	case 'd':
		// don't change the precision -- if they want a giant integer and we have it...

		// decimal -- rounds to the precision of digits left of the decimal place
		if printable.exponent < 0 {
			printable.SetUint64(0)
		} else if len(printable.significand)-1 > printable.exponent {
			printable.SetPrecision(uint(printable.exponent) + 1)
		}
		if len(printable.significand) == 0 {
			o.WriteString("0")
		} else {
			for _, v := range printable.significand {
				o.WriteString(fmt.Sprintf("%c", v+asciiOffset))
			}
			trailing := printable.exponent - len(printable.significand) + 1
			for i := 0; i < trailing; i++ {
				o.WriteString("0")
			}
		}
	case 'e':
		printable.SetPrecision(uint(p))

		// scientific notation
		if len(printable.significand) == 0 {
			o.WriteString("0")
		} else {
			o.WriteString(fmt.Sprintf("%c", printable.significand[0]+asciiOffset))

			if len(printable.significand) > 1 {
				o.WriteString(".")

				for _, v := range printable.significand[1:] {
					o.WriteString(fmt.Sprintf("%c", v+asciiOffset))
				}
			}
			o.WriteString(fmt.Sprintf("e%v", printable.exponent))
		}
	case 'f':
		printable.SetPrecision(uint(p))

		// floating point notation
		if len(printable.significand) == 0 {
			o.WriteString("0.0")
		} else {
			if printable.exponent < 0 {
				o.WriteString("0.")
				for printable.exponent < -1 {
					o.WriteString("0")
					printable.exponent++
				}
				for _, v := range printable.significand {
					o.WriteString(fmt.Sprintf("%c", v+asciiOffset))
				}
			} else {
				for printable.exponent >= 0 && len(printable.significand) > 0 {
					o.WriteString(fmt.Sprintf("%c", printable.significand[0]+asciiOffset))
					printable.significand = printable.significand[1:]
					printable.exponent--
				}
				if printable.exponent >= 0 {
					// trailing zeros in the integer part
					trailing := printable.exponent - len(printable.significand) + 1
					for i := 0; i < trailing; i++ {
						o.WriteString("0")
					}
				}
				o.WriteString(".")
				if len(printable.significand) != 0 {
					for _, v := range printable.significand {
						o.WriteString(fmt.Sprintf("%c", v+asciiOffset))
					}
				} else {
					o.WriteString("0")
				}
			}
		}
	case 'v':
		o.Reset()
		// attempt a natural notation based on the value
		if !precisionSet && abs(printable.exponent)+len(printable.significand) > sensibleSize {
			// scientific notation
			printable.SetPrecision(uint(p))
			o.WriteString(fmt.Sprintf("%.*e", printable.precision, printable))
		} else if printable.IsInteger() {
			// integer
			o.WriteString(fmt.Sprintf("%.*d", printable.precision, printable))
		} else {
			// floating point
			printable.SetPrecision(uint(p))
			o.WriteString(fmt.Sprintf("%.*f", printable.precision, printable))
		}
	}

	// TODO: width
	//width, hasWidth := s.Width()
	s.Write(o.Bytes())
}

// Return the integer part of a real number by truncating.
func (x *Real) Integer() *Real {
	z := x.Copy()

	if z.IsInf() {
		return z
	} else if z.IsNaN() {
		return z
	}

	if z.exponent < 0 {
		z.SetUint64(0)
	} else if z.exponent < len(x.significand)-1 {
		z.significand = z.significand[:z.exponent+1]
	}
	return z
}

// Return a uint64 representation of the number, if possible. If not possible,
// err will be non-nil.
func (x *Real) Uint64() (uint64, error) {
	s := fmt.Sprintf("%d", x)
	u, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return u, nil
}

// Return an int64 representation of the number, if possible. If not possible,
// err will be non-nil.
func (x *Real) Int64() (int64, error) {
	s := fmt.Sprintf("%d", x)
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// Return a float64 representation of the number, if possible. If not possible,
// err will be non-nil.
func (x *Real) Float64() (float64, error) {
	if x.IsInf() {
		sign := 1
		if x.negative {
			sign = -1
		}
		return math.Inf(sign), nil
	} else if x.IsNaN() {
		return math.NaN(), nil
	}

	s := fmt.Sprintf("%e", x)
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}
	return f, nil
}

// ParseReal converts a string s to a Real with the given precision p.
//
// Input beyond the given precision is ignored but not considered an error.
//
// Input can be as a fixed precision number or in scientific notation, using a
// lower case 'e' for the exponent.
func ParseReal(s string, p uint) (*Real, error) {
	s = strings.ToLower(s)

	x := new(Real)

	// negative sign
	if len(s) > 0 && s[0] == '-' {
		x.negative = true
		s = s[1:]
	} else if len(s) > 0 && s[0] == '+' {
		s = s[1:]
	}

	if s == "inf" {
		x.form = FormInf
		return x, nil
	} else if s == "nan" {
		x.form = FormNaN
		return x, nil
	}

	// significand
	var radixSet bool
	var oneDigit bool
	for len(s) > 0 {
		if s[0] == '.' {
			radixSet = true
			x.exponent = len(x.significand) - 1
		} else if s[0] >= '0' && s[0] <= '9' {
			x.significand = append(x.significand, byte(s[0])-asciiOffset)
		} else if s[0] == 'e' {
			// exponent
			if len(x.significand) == 0 {
				return nil, ErrInvalidCharacter
			}
			break
		} else {
			return nil, ErrInvalidCharacter
		}
		s = s[1:]
		oneDigit = true
	}

	if !oneDigit {
		return nil, ErrInvalidCharacter
	}

	if !radixSet {
		x.exponent = len(x.significand) - 1
	}

	// optional exponent
	if len(s) > 0 {
		if s[0] != 'e' {
			return nil, ErrInvalidCharacter
		}
		s = s[1:]
		exp, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil, err
		}
		x.exponent += int(exp)
	}

	x.SetPrecision(p)
	return x, nil
}
