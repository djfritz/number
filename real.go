package real

import (
	"fmt"
	"math"
	"strconv"
)

type Real struct {
	digits    []byte // decimal digits -- only valid values are 0-9
	negative  bool   // true if the number is negative
	exponent  int
	form      int // ±∞, NaN, or a real number
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
		exponent:  r.exponent,
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
}

func (r *Real) SetUint64(x uint64) {
	r.digits = []byte{}
	r.negative = false
	r.exponent = 0
	r.form = REAL
	if x == 0 {
		return
	}
	for x != 0 {
		r.digits = append([]byte{byte(x % 10)}, r.digits...)
		x /= 10
	}
	r.exponent = len(r.digits) - 1
	r.trim()
}

func (r *Real) SetFloat64(x float64) {
	r.digits = []byte{}
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
	// approach here would be no better than just doing dtoa() and parsing
	// the string, so we do exactly that...
	s := fmt.Sprintf("%.17e", x)
	if s[0] == '-' {
		r.negative = true
		s = s[1:]
	}

	// digits
	for i, v := range s {
		if v == 'e' {
			s = s[i+1:]
			break
		}
		if v == '.' {
			continue
		}
		r.digits = append(r.digits, byte(v)-0x30)
	}

	// exponent
	exp, err := strconv.Atoi(s)
	if err != nil {
		panic("could not parse exponent")
	}
	r.exponent = exp
	r.trim()
}

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
		for _, v := range r.digits {
			s += fmt.Sprintf("%c", v+0x30)
		}
	} else {
		for i, v := range r.digits {
			s += fmt.Sprintf("%c", v+0x30)
			if i == r.exponent && i != len(r.digits)-1 {
				s += "."
			}
		}
		if r.exponent > len(r.digits)-1 {
			for i := 0; i < r.exponent-(len(r.digits)-1); i++ {
				s += "0"
			}
		}
	}
	return s
}

func (r *Real) trim() {
	for i := 0; i < len(r.digits); i++ {
		if r.digits[i] != 0 {
			break
		}
		r.digits = r.digits[1:]
	}
	for i := len(r.digits) - 1; i >= 0; i-- {
		if r.digits[i] != 0 {
			break
		}
		r.digits = r.digits[:i]
	}
}
