// Copyright 2024 David Fritz. All rights reserved.
// This software may be modified and distributed under the terms of the BSD
// 2-clause license. See the LICENSE file for details.

package number

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
)

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

func (x *Real) Format(s fmt.State, verb rune) {
	printable := x.Copy()
	if p, ok := s.Precision(); ok {
		printable.SetPrecision(uint(p))
	} else {
		printable.SetPrecision(defaultPrintPrecision)
	}

	var o bytes.Buffer
	if printable.negative {
		o.WriteString("-")
	}

	switch verb {
	case 'd':
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
		if abs(printable.exponent)-len(printable.significand) > sensibleSize {
			// scientific notation
			o.WriteString(fmt.Sprintf("%e", printable))
		} else if printable.IsInteger() {
			// integer
			o.WriteString(fmt.Sprintf("%d", printable))
		} else {
			// floating point
			o.WriteString(fmt.Sprintf("%f", printable))
		}
	}

	// TODO: width
	//width, hasWidth := s.Width()
	s.Write(o.Bytes())
}

// Return the integer part of a real number.
func (x *Real) Integer() *Real {
	z := x.Copy()
	if z.exponent < 0 {
		z.SetUint64(0)
	} else if z.exponent < len(x.significand)-1 {
		z.significand = z.significand[:z.exponent+1]
	}
	return z
}

func (x *Real) Uint64() (uint64, bool) {
	s := fmt.Sprintf("%d", x)
	u, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, false
	}
	return u, true
}

func (x *Real) Int64() (int64, bool) {
	s := fmt.Sprintf("%d", x)
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, false
	}
	return i, true
}

func (x *Real) Float64() (float64, bool) {
	s := fmt.Sprintf("%e", x)
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, false
	}
	return f, true
}

// ParseReal converts a string s to a Real with the given precision p.
//
// Input beyond the given precision is ignored but not considered an error.
//
// Input can be as a fixed precision number or in scientific notation, using a
// lower case 'e' for the exponent.
func ParseReal(s string, p uint) (*Real, error) {
	x := new(Real)

	// negative sign
	if len(s) > 0 && s[0] == '-' {
		x.negative = true
		s = s[1:]
	}

	// significand
	var radixSet bool
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
