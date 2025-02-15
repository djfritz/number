// Copyright 2024 David Fritz. All rights reserved.
// This software may be modified and distributed under the terms of the BSD
// 2-clause license. See the LICENSE file for details.

package real

import (
	"fmt"
	"math"
	"strconv"
)

// A real number. Internally stored as a real number in decimal scientific notation.
type Real struct {
	significand []byte // decimal significand -- only valid values are 0-9
	negative    bool   // true if the number is negative
	exponent    int    // exponent
	precision   uint   // maximum allowed precision of the significand in decimal digits
}

const (
	DefaultPrecision               = 34 // The default precision for a real number. Expressed in decimal digits.
	internalPrecisionBuffer        = 10
	float64MinimumDecimalPrecision = 15 // minimum number of correct decimal digits in a float64
)

// Copy returns a deep copy of the real value.
func (x *Real) Copy() *Real {
	z := &Real{
		negative:  x.negative,
		precision: x.precision,
	}
	z.CopyValue(x)
	return z

}

// Copy just the value of y into x, leaving x's precision and mode the same.
// The result will round if needed.
func (x *Real) CopyValue(y *Real) {
	x.exponent = y.exponent
	x.significand = make([]byte, len(y.significand))
	copy(x.significand, y.significand)
	x.round()
}

// Create a zero-value real number, copying precision, form, and mode from the
// given real value. Used in internal functions to maintain precision while
// making new values based on operands.
func initFrom(x *Real) *Real {
	return &Real{
		significand: []byte{},
		precision:   x.precision,
	}
}

// Same as initFrom(), but takes the maximum precision of x,y. Mode and form
// always copy from x.
func initFrom2(x, y *Real) *Real {
	r := &Real{
		significand: []byte{},
	}
	if x.precision > y.precision {
		r.precision = x.precision
	} else {
		r.precision = y.precision
	}
	return r
}

// Return a new real number set to the given signed int64, with the default
// rounding mode and precision.
func NewInt64(x int64) *Real {
	r := new(Real)
	r.SetInt64(x)
	return r
}

// Return a new real number set to the given unsigned uint64, with the default
// rounding mode and precision.
func NewUint64(x uint64) *Real {
	r := new(Real)
	r.SetUint64(x)
	return r
}

// Return a new real number set to the given float64, with the default
// rounding mode and precision.
func NewFloat64(x float64) *Real {
	r := new(Real)
	r.SetFloat64(x)
	return r
}

// Set the precision of the given number and round if necessary.
func (x *Real) SetPrecision(y uint) {
	x.precision = y
	x.round()
}

// Set a real number to the given signed int64. Rounding mode and precision are
// left unchanged. If precision is lower than the given value, rounding occurs.
func (x *Real) SetInt64(y int64) {
	if y < 0 {
		x.SetUint64(uint64((^y) + 1))
		x.negative = true
	} else {
		x.SetUint64(uint64(y))
	}
}

// Set a real number to the given unsigned uint64. Rounding mode and precision
// are left unchanged. If precision is lower than the given value, rounding
// occurs.
func (x *Real) SetUint64(y uint64) {
	x.significand = []byte{}
	x.negative = false
	x.exponent = 0
	if y == 0 {
		return
	}
	for y != 0 {
		x.significand = append([]byte{byte(y % 10)}, x.significand...)
		y /= 10
	}
	x.exponent = len(x.significand) - 1
	x.round()
}

// Set a real number to the given float64. Rounding mode and precision are left
// unchanged. If precision is lower than the given value, rounding occurs.
func (x *Real) SetFloat64(y float64) {
	x.significand = []byte{}
	x.negative = false

	// TODO: forms
	//	if math.IsInf(x, 1) {
	//		r.form = INF
	//		return
	//	} else if math.IsInf(x, -1) {
	//		r.form = NINF
	//		return
	//	} else if math.IsNaN(x) {
	//		r.form = NAN
	//		return
	if y == 0 {
		return
	}

	// an efficient binary to decimal algorithm is still a fantasy. Any
	// approach here would be no better than just doing dtoa() and parsing
	// the string, so we do exactly that...
	s := fmt.Sprintf("%.17e", y)
	if s[0] == '-' {
		x.negative = true
		s = s[1:]
	}

	// significand
	for i, v := range s {
		if v == 'e' {
			s = s[i+1:]
			break
		}
		if v == '.' {
			continue
		}
		x.significand = append(x.significand, byte(v)-0x30)
	}

	// exponent
	exp, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("could not parse exponent %v", s))
	}
	x.exponent = exp
	x.round()
}

// Return the string form of the real number in scientific notation.
func (x *Real) String() string {
	var s string
	if x.negative {
		s = "-"
	}
	if len(x.significand) == 0 {
		s += "0"
		return s
	}
	s += fmt.Sprintf("%c", x.significand[0]+0x30)

	if len(x.significand) > 1 {
		s += "."

		for _, v := range x.significand[1:] {
			s += fmt.Sprintf("%c", v+0x30)
		}
	}
	s += fmt.Sprintf("e%v", x.exponent)
	return s
}

// Return the string form of the real number in ??? notation.
func (x *Real) StringRegular() string {
	var s string
	if x.negative {
		s = "-"
	}
	if x.exponent < 0 {
		s += "0."
		for i := 0; i < (x.exponent*-1)-1; i++ {
			s += "0"
		}
		for _, v := range x.significand {
			s += fmt.Sprintf("%c", v+0x30)
		}
	} else {
		for i, v := range x.significand {
			s += fmt.Sprintf("%c", v+0x30)
			if i == x.exponent && i != len(x.significand)-1 {
				s += "."
			}
		}
		if x.exponent > len(x.significand)-1 {
			for i := 0; i < x.exponent-(len(x.significand)-1); i++ {
				s += "0"
			}
		}
	}
	return s
}

// Trim removes leading and trailing zeros from a normalized value.
func (x *Real) trim() {
	var i int
	for i = 0; i < len(x.significand); i++ {
		if x.significand[i] != 0 {
			break
		}
	}
	x.significand = x.significand[i:]
	x.exponent -= i
	for i := len(x.significand) - 1; i >= 0; i-- {
		if x.significand[i] != 0 {
			break
		}
		x.significand = x.significand[:i]
	}
}

// Round the value to the set precision and rounding mode, if necessary.
func (x *Real) round() {
	x.validate()
	x.roundTo(x.precision)
}

// Round the value to the given precision and rounding mode.
func (x *Real) roundTo(p uint) {
	defer x.trim()

	if uint(len(x.significand)) <= p {
		// number is exact, no rounding needed.
		return
	}

	for i := len(x.significand) - 1; i >= int(p); i-- {
		d := x.significand[i]
		switch {
		case d < 5:
			// round down
		case d > 5:
			// round up
			if i == 0 {
				x.significand[0] = 1
				x.exponent++
				return
			} else {
				x.significand[i-1]++
			}
		case d == 5:
			// round to nearest even
			if x.significand[i-1]%2 != 0 {
				if i == 0 {
					x.significand[0] = 1
					x.exponent++
					return
				} else {
					x.significand[i-1]++
				}
			}
		}
		x.significand[i] = 0
	}

	// now unwind to the left to make sure we don't have any lingering carry
	for i := int(p) - 1; i >= 0; i-- {
		if x.significand[i] < 10 {
			break
		}
		x.significand[i] -= 10

		if i == 0 {
			// pad
			x.significand = append([]byte{1}, x.significand...)
			break
		}
		x.significand[i-1]++
	}
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

func (x *Real) validate() {
	if x.precision == 0 {
		x.precision = DefaultPrecision
	}
}

func umax(a, b uint) uint {
	if a > b {
		return a
	}
	return b
}

func umin(a, b uint) uint {
	if a < b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

// Returns true if x == 0.
func (x *Real) IsZero() bool {
	if len(x.significand) == 0 {
		return true
	}
	return false
}

// Returns the remaining number of iterations required given the known digits
// and given precision. Assumes quadratic convergence.
func estimateConvergence(known, precision uint) int {
	iterations := math.Ceil(math.Log2(float64(precision)) - math.Log2(float64(known)))
	return int(iterations)
}
