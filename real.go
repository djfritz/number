package real

import (
	"fmt"
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
	DefaultPrecision = 34 // the default precision for a real number. Expressed in decimal digits.
)

// Copy returns a deep copy of the real value
func (r *Real) Copy() *Real {
	z := &Real{
		negative:  r.negative,
		precision: r.precision,
	}
	z.CopyValue(r)
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
func (r *Real) SetPrecision(x uint) {
	r.precision = x
	r.round()
}

// Set a real number to the given signed int64. Rounding mode and precision are
// left unchanged. If precision is lower than the given value, rounding occurs.
func (r *Real) SetInt64(x int64) {
	if x < 0 {
		r.SetUint64(uint64((^x) + 1))
		r.negative = true
	} else {
		r.SetUint64(uint64(x))
	}
}

// Set a real number to the given unsigned uint64. Rounding mode and precision
// are left unchanged. If precision is lower than the given value, rounding
// occurs.
func (r *Real) SetUint64(x uint64) {
	r.significand = []byte{}
	r.negative = false
	r.exponent = 0
	if x == 0 {
		return
	}
	for x != 0 {
		r.significand = append([]byte{byte(x % 10)}, r.significand...)
		x /= 10
	}
	r.exponent = len(r.significand) - 1
	r.round()
}

// Set a real number to the given float64. Rounding mode and precision are left
// unchanged. If precision is lower than the given value, rounding occurs.
func (r *Real) SetFloat64(x float64) {
	r.significand = []byte{}
	r.negative = false

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
	if x == 0 {
		return
	}

	// an efficient binary to decimal algorithm is still a fantasy. Any
	// approach here would be no better than just doing dtoa() and parsing
	// the string, so we do exactly that...
	s := fmt.Sprintf("%.17e", x)
	if s[0] == '-' {
		r.negative = true
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
		r.significand = append(r.significand, byte(v)-0x30)
	}

	// exponent
	exp, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("could not parse exponent %v", s))
	}
	r.exponent = exp
	r.round()
}

// Return the string form of the real number in scientific notation.
func (r *Real) String() string {
	var s string
	if r.negative {
		s = "-"
	}
	if len(r.significand) == 0 {
		s += "0"
		return s
	}
	s += fmt.Sprintf("%c", r.significand[0]+0x30)

	if len(r.significand) > 1 {
		s += "."

		for _, v := range r.significand[1:] {
			s += fmt.Sprintf("%c", v+0x30)
		}
	}
	s += fmt.Sprintf("e%v", r.exponent)
	return s
}

// Return the string form of the real number in ??? notation.
func (r *Real) StringRegular() string {
	var s string
	if r.negative {
		s = "-"
	}
	if r.exponent < 0 {
		s += "0."
		for i := 0; i < (r.exponent*-1)-1; i++ {
			s += "0"
		}
		for _, v := range r.significand {
			s += fmt.Sprintf("%c", v+0x30)
		}
	} else {
		for i, v := range r.significand {
			s += fmt.Sprintf("%c", v+0x30)
			if i == r.exponent && i != len(r.significand)-1 {
				s += "."
			}
		}
		if r.exponent > len(r.significand)-1 {
			for i := 0; i < r.exponent-(len(r.significand)-1); i++ {
				s += "0"
			}
		}
	}
	return s
}

// Trim removes leading and trailing zeros from a normalized value.
func (r *Real) trim() {
	var i int
	for i = 0; i < len(r.significand); i++ {
		if r.significand[i] != 0 {
			break
		}
	}
	r.significand = r.significand[i:]
	r.exponent -= i
	for i := len(r.significand) - 1; i >= 0; i-- {
		if r.significand[i] != 0 {
			break
		}
		r.significand = r.significand[:i]
	}
}

// Round the value to the set precision and rounding mode, if necessary.
func (r *Real) round() {
	r.validate()
	r.roundTo(r.precision)
}

// Round the value to the given precision and rounding mode.
func (r *Real) roundTo(p uint) {
	defer r.trim()

	if uint(len(r.significand)) <= p {
		// number is exact, no rounding needed.
		return
	}

	for i := len(r.significand) - 1; i >= int(p); i-- {
		d := r.significand[i]
		switch {
		case d < 5:
			// round down
		case d > 5:
			// round up
			if i == 0 {
				r.significand[0] = 1
				r.exponent++
				return
			} else {
				r.significand[i-1]++
			}
		case d == 5:
			// round to nearest even
			if r.significand[i-1]%2 != 0 {
				if i == 0 {
					r.significand[0] = 1
					r.exponent++
					return
				} else {
					r.significand[i-1]++
				}
			}
		}
		r.significand[i] = 0
	}

	// now unwind to the left to make sure we don't have any lingering carry
	for i := int(p) - 1; i >= 0; i-- {
		if r.significand[i] < 10 {
			break
		}
		r.significand[i] -= 10

		if i == 0 {
			// pad
			r.significand = append([]byte{1}, r.significand...)
			break
		}
		r.significand[i-1]++
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

func (r *Real) IsZero() bool {
	if len(r.significand) == 0 {
		return true
	}
	return false
}
