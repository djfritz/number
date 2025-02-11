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
	DefaultPrecision = 100 // the default precision for a real number. Expressed in decimal digits.
)

// Copy returns a deep copy of the real value
func (r *Real) Copy() *Real {
	return &Real{
		significand: append([]byte{}, r.significand...),
		negative:    r.negative,
		exponent:    r.exponent,
		precision:   r.precision,
	}
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
		panic("could not parse exponent")
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
	if r.precision == 0 {
		r.precision = DefaultPrecision
	}
	r.roundTo(r.precision)
}

// Round the value to the given precision and rounding mode.
func (r *Real) roundTo(p uint) {
	defer r.trim()

	if uint(len(r.significand)) <= p {
		// number is exact, no rounding needed.
		return
	}

	for i := uint(len(r.significand)) - 1; i >= p; i-- {
		d := r.significand[i]
		switch {
		case d < 5:
			// round down
		case d > 5:
			// round up
			r.significand[i-1]++
		case d == 5:
			// round to nearest even
			if r.significand[i-1]%2 != 0 {
				r.significand[i-1]++
			}
		}
		r.significand[i] = 0
	}

	// now unwind to the left to make sure we don't have any lingering carry
	for i := p - 1; i >= 0; i-- {
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

// Adjust shifts x and y to the same exponent and returns the significand of x and y
// as byte slices, as well as the exponent the slices have. The precision of
// the returned byte slices equals that of the higher precision operand, and
// padding is added to ensure the byte slices are the same length. Shifted
// values are rounded according to the rounding mode set in that operand.
func adjust(x, y *Real) ([]byte, []byte, int) {
	p := x.precision
	if y.precision > p {
		p = y.precision
	}

	e := x.exponent
	if y.exponent > e {
		e = y.exponent
	}

	ar := x
	br := y
	if uint(len(x.significand)) > p {
		ar = x.Copy()
		ar.SetPrecision(p)
	}
	if uint(len(y.significand)) > p {
		br = y.Copy()
		br.SetPrecision(p)
	}

	a := shift(ar.significand, x.exponent-e, p)
	b := shift(br.significand, y.exponent-e, p)
	return a, b, e
}

// Shift the byte slice by e bytes, keeping the size of the byte slice in p.
// Pad to p bytes if needed. Positive e is a left shift.
func shift(x []byte, e int, p uint) []byte {
	z := append([]byte{}, x...)
	if uint(len(z)) > p {
		z = z[:p]
	} else {
		pad := make([]byte, p-uint(len(z)))
		z = append(z, pad...)
	}

	if e != 0 {
		eabs := e
		if eabs < 0 {
			eabs *= -1
		}
		if eabs > len(z) {
			// the entire slice will be shifted off, just return zeros
			z = make([]byte, p)
		} else {
			pad := make([]byte, eabs)
			if e < 0 {
				z = z[:len(z)-eabs]
				z = append(pad, z...)
			} else if e > 0 {
				z = z[eabs:]
				z = append(z, pad...)
			}
		}
	}

	return z
}
