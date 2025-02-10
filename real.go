package real

import (
	"fmt"
	"math"
	"strconv"
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

func initFrom(x *Real) *Real {
	return &Real{
		digits:    []byte{},
		precision: x.precision,
	}
}

func initFrom2(x, y *Real) *Real {
	r := &Real{
		digits: []byte{},
	}
	if x.precision > y.precision {
		r.precision = x.precision
	} else {
		r.precision = y.precision
	}
	return r
}

func (r *Real) SetInt64(x int64) {
	if x < 0 {
		r.SetUint64(uint64((^x) + 1))
		r.negative = true
	} else {
		r.SetUint64(uint64(x))
	}
	r.round()
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
	r.round()
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

	// an efficient binary to decimal algorithm is still a fantasy. Any
	// approach here would be no better than just doing dota() and parsing
	// the string, so we do exactly that...
	s := fmt.Sprintf("%.17e", x)
	if s[0] == '-' {
		r.negative = true
		s = s[1:]
	}

	// digits
	var hasDecimal bool
	for i, v := range s {
		if v == 'e' {
			s = s[i+1:]
			break
		}
		if v == '.' {
			hasDecimal = true
			continue
		}
		r.digits = append(r.digits, byte(v)-0x30)
		if hasDecimal {
			r.decimal++
		}
	}

	// exponent
	var padFront bool
	if s[0] == '-' {
		padFront = true
		s = s[1:]
	}

	exp, err := strconv.Atoi(s)
	if err != nil {
		panic("could not parse exponent")
	}
	pad := make([]byte, exp)
	if padFront {
		r.digits = append(pad, r.digits...)
		r.decimal += uint(exp)
	} else {
		r.digits = append(r.digits, pad...)
	}

	r.round()
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
