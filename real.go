// Copyright 2025 David Fritz. All rights reserved.
// This software may be modified and distributed under the terms of the BSD
// 2-clause license. See the LICENSE file for details.

package number

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
	form        int    // other forms of an implementation of a real number -- infinity, NaN, etc.
	mode        int    // rounding mode
}

// Number forms
const (
	FormReal = iota // A finite real number
	FormNaN         // Not a number
	FormInf         // Infinity
)

// The default precision for a real number. Expressed in decimal digits. 34
// digits is equivalent IEEE-754-2008 128-bit decimal floating point.
const DefaultPrecision = 34

const (
	internalPrecisionBuffer        = 5  // additional precision to use when performing operations internally
	float64MinimumDecimalPrecision = 15 // minimum number of correct decimal digits in a float64
)

// Copy returns a deep copy of x.
func (x *Real) Copy() *Real {
	z := &Real{
		precision: x.precision,
		mode:      x.mode,
	}
	z.CopyValue(x)
	return z

}

// Copy just the value of y into x, leaving x's precision and mode the same.
// The result will round if needed.
func (x *Real) CopyValue(y *Real) {
	x.negative = y.negative
	x.exponent = y.exponent
	x.significand = make([]byte, len(y.significand))
	copy(x.significand, y.significand)
	x.form = y.form
	x.round()
}

// Create a zero-value real number, copying precision, form, and mode from the
// given real value. Used in internal functions to maintain precision while
// making new values based on operands.
func initFrom(x *Real) *Real {
	return &Real{
		significand: []byte{},
		precision:   x.precision,
		mode:        x.mode,
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
		r.mode = x.mode
	} else {
		r.precision = y.precision
		r.mode = y.mode
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

// Returns the assigned precision of the number.
func (x *Real) Precision() uint {
	x.validate()
	return x.precision
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

	if y == 0 {
		return
	}

	if math.IsNaN(y) {
		x.form = FormNaN
		return
	} else if math.IsInf(y, -1) {
		x.form = FormInf
		x.negative = true
		return
	} else if math.IsInf(y, 1) {
		x.form = FormInf
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

// Return the absolute value of x.
func (x *Real) Abs() *Real {
	z := x.Copy()
	z.negative = false
	return z
}

// Returns true if x == 0.
func (x *Real) IsZero() bool {
	if x.form == FormReal && len(x.significand) == 0 {
		return true
	}
	return false
}

// Returns true if x is ±Inf.
func (x *Real) IsInf() bool {
	return x.form == FormInf
}

// Returns true if x is NaN.
func (x *Real) IsNaN() bool {
	return x.form == FormNaN
}

// Returns true if x is an integer.
func (x *Real) IsInteger() bool {
	if x.exponent < len(x.significand)-1 {
		return false
	}
	return true
}

// Returns the remaining number of iterations required given the known digits
// and given precision. Assumes quadratic convergence.
func estimateConvergence(known, precision uint) int {
	iterations := math.Ceil(math.Log2(float64(precision)) - math.Log2(float64(known)))
	return int(iterations)
}

// Returns the floor of x.
func (x *Real) Floor() *Real {
	return x.Integer()
}

// Returns the ceiling of x.
func (x *Real) Ceiling() *Real {
	if x.IsInteger() {
		return x.Copy()
	}
	return x.Integer().Add(NewInt64(1))
}

// Prepare internal precision -- used to set a sane internal precision before
// performing an operation.
func (x *Real) pip(p uint) {
	if p < DefaultPrecision {
		x.precision = DefaultPrecision
	}
	x.precision += internalPrecisionBuffer
}
