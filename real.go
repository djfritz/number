package real

import (
	"fmt"
	"math"
)

type Real struct {
	digits    []byte // decimal digits -- only valid values are 0-9
	negative  bool   // true if the number is negative
	decimal   uint   // decimal place offset from the right. 0 mean the number is an integer
	form      int    // ±∞, NaN, or a real number
	mode      int
	precision uint
}

const (
	REAL = iota
	INF
	NINF
	NAN
)

const (
	EVEN = iota
)

func (r *Real) Copy() *Real {
	return &Real{
		digits:    append([]byte{}, r.digits...),
		negative:  r.negative,
		decimal:   r.decimal,
		form:      r.form,
		mode:      r.mode,
		precision: r.precision,
	}
}

func (r *Real) SetInt64(x int64) {
	if x < 0 {
		r.SetUint64(uint64((^x) + 1))
		r.negative = true
	} else {
		r.SetUint64(uint64(x))
	}
	r.fix()
}

func (r *Real) SetUint64(x uint64) {
	r.digits = []byte{}
	r.negative = false
	r.decimal = 0
	r.form = REAL
	for x != 0 {
		r.digits = append([]byte{byte(x % 10)}, r.digits...)
		x /= 10
	}
	r.fix()
}

func (r *Real) SetFloat64(x float64) {
	r.digits = []byte{}
	r.decimal = 0
	r.negative = false
	r.form = REAL

	if math.IsInf(x, 1) {
		r.form = INF
		return
	} else if math.IsInf(x, -1) {
		r.form = NINF
		return
	} else if math.IsNaN(x) {
		r.form = NAN
		return
	} else if x == 0 {
		return
	}

	fmant, _ := math.Frexp(x)
	r.SetUint64(1<<63 | math.Float64bits(fmant)<<11)
	r.decimal = uint(len(r.digits))

	// r = r.Mul(r.Pow(2,exp))

	// TODO: round precision
}

func (r *Real) String() string {
	// TODO: forms
	var s string
	if r.negative {
		s = "-"
	}
	for i, v := range r.digits {
		if uint(len(r.digits))-r.decimal == uint(i) {
			s += "."
		}
		s += fmt.Sprintf("%c", v+0x30)
	}
	return s
}
