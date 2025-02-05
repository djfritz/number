package real

func adjust(x, y *Real) ([]byte, []byte, uint) {
	var a, b []byte
	var d uint

	a = append(a, x.digits[uint(len(x.digits))-x.decimal:]...)
	b = append(b, y.digits[uint(len(y.digits))-y.decimal:]...)

	if len(a) > len(b) {
		pad := make([]byte, len(a)-len(b))
		b = append(b, pad...)
	} else if len(a) < len(b) {
		pad := make([]byte, len(b)-len(a))
		a = append(a, pad...)
	}
	d = uint(len(a))

	ia := append([]byte{}, x.digits[:uint(len(x.digits))-x.decimal]...)
	ib := append([]byte{}, y.digits[:uint(len(y.digits))-y.decimal]...)

	a = append(ia, a...)
	b = append(ib, b...)

	if len(a) > len(b) {
		diff := len(a) - len(b)
		pad := make([]byte, diff, len(a))
		b = append(pad, b...)
	} else if len(a) < len(b) {
		diff := len(b) - len(a)
		pad := make([]byte, diff, len(b))
		a = append(pad, a...)
	}

	return a, b, d
}

func (r *Real) trim() {
	for _, v := range r.digits {
		if v != 0 {
			break
		}
		r.digits = r.digits[1:]
	}
	for i := len(r.digits) - 1; uint(i) > uint(len(r.digits))-r.decimal; i-- {
		if r.digits[i] != 0 {
			break
		}
		r.digits = r.digits[:i]
	}
}

func (r *Real) Compare(x *Real) int {
	// TODO: deal with forms
	if !r.negative && x.negative {
		return 1
	} else if r.negative && !x.negative {
		return -1
	}

	a, b, _ := adjust(r, x)

	for i := range a {
		if a[i] > b[i] {
			if r.negative {
				return -1
			}
			return 1
		} else if a[i] < b[i] {
			if r.negative {
				return 1
			}
			return -1
		}
	}

	return 0
}

// slices should be adjusted first
func absCompare(a, b []byte) int {
	for i := range a {
		if a[i] > b[i] {
			return 1
		} else if a[i] < b[i] {
			return -1
		}
	}

	return 0
}

func (r *Real) Add(x *Real) *Real {
	// TODO: deal with forms

	a, b, d := adjust(r, x)
	a = append([]byte{0}, a...)
	b = append([]byte{0}, b...)

	z := &Real{
		decimal: d,
	}

	if r.negative == x.negative {
		z.negative = r.negative

		for i := len(a) - 1; i >= 0; i-- {
			a[i] += b[i]
			if a[i] > 10 {
				a[i] -= 10
				a[i-1]++
			}
		}
	} else {
		z.negative = r.negative
		if absCompare(a, b) == -1 {
			a, b = b, a
			z.negative = !r.negative
		}

		for i := len(a) - 1; i >= 0; i-- {
			if a[i] >= b[i] {
				a[i] -= b[i]
			} else {
				a[i] = (10 + a[i]) - b[i]
				b[i-1]++
			}
		}
	}

	z.digits = a
	z.trim()

	// TODO: rounding

	return z

}

func (r *Real) Sub(x *Real) *Real {
	y := x.Copy()
	y.negative = !y.negative
	return r.Add(y)
}
