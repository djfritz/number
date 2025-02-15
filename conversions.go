// Copyright 2024 David Fritz. All rights reserved.
// This software may be modified and distributed under the terms of the BSD
// 2-clause license. See the LICENSE file for details.

package real

import (
	"bytes"
	"fmt"
)

const (
	asciiOffset           = 0x30 // offset to add to bytes to get their ASCII representation
	sensibleSize          = 40   // sensible number of digits to print before engaging scientific notation for the %v verb
	defaultPrintPrecision = DefaultPrecision
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
		}
		trailing := printable.exponent - len(printable.significand) + 1
		for i := 0; i < trailing; i++ {
			o.WriteString("0")
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
