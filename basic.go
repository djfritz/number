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

func (r *Real) round() {
	if r.precision == 0 {
		r.precision = 100
	}
	r.roundTo(r.precision)
}

func (r *Real) roundTo(p uint) {
	defer r.trim()

	// TODO: other rounding modes
	if r.decimal <= p {
		return
	}

	// rounding position
	rp := uint(len(r.digits)) - (r.decimal - p) - 1

	// attempt correct rounding to the precision we have

	for i := uint(len(r.digits)) - 1; i > rp; i-- {
		d := r.digits[i]
		switch {
		case d < 5:
			// round down
		case d > 5:
			// round up
			r.digits[i-1]++
		case d == 5:
			// round to nearest even
			if r.digits[i-1]%2 != 0 {
				r.digits[i-1]++
			}
		}
		r.digits[i] = 0
	}

	// now unwind to the left to make sure we don't have any lingering carry
	for i := rp; i >= 0; i-- {
		if r.digits[i] < 10 {
			break
		}
		r.digits[i] -= 10

		if i == 0 {
			// pad
			r.digits = append([]byte{1}, r.digits...)
			break
		}
		r.digits[i-1]++
	}

	r.digits = r.digits[:rp+1]
	r.decimal = p
}

func (r *Real) trim() {
	for _, v := range r.digits[:uint(len(r.digits))-r.decimal] {
		if v != 0 {
			break
		}
		r.digits = r.digits[1:]
	}
	if len(r.digits) == 0 {
		return
	}
	for i := len(r.digits) - 1; uint(i) > uint(len(r.digits))-r.decimal; i-- {
		if r.digits[i] != 0 {
			break
		}
		r.digits = r.digits[:i]
		r.decimal--
	}
}

func (r *Real) Compare(x *Real) int {
	// TODO: deal with forms
	if !r.negative && x.negative {
		return 1
	} else if r.negative && !x.negative {
		return -1
	}

	p := r.precision
	if x.precision < p {
		p = x.precision
	}

	a, b, d := adjust(r, x)

	if d > p {
		// we only care about the digits up to the given precision
		a = a[:uint(len(a))-(d-p)-1]
		b = b[:uint(len(b))-(d-p)-1]
	}

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

func (r *Real) Abs() *Real {
	z := r.Copy()
	z.negative = false
	return z
}

func (r *Real) Add(x *Real) *Real {
	z := r.add(x)
	z.round()
	return z
}

func (r *Real) add(x *Real) *Real {
	// TODO: deal with forms

	a, b, d := adjust(r, x)
	a = append([]byte{0}, a...)
	b = append([]byte{0}, b...)

	z := initFrom(r)
	z.decimal = d

	if r.negative == x.negative {
		z.negative = r.negative

		for i := len(a) - 1; i >= 0; i-- {
			a[i] += b[i]
			if a[i] >= 10 {
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

	return z

}

func (r *Real) Sub(x *Real) *Real {
	z := r.sub(x)
	z.round()
	return z
}

func (r *Real) sub(x *Real) *Real {
	y := x.Copy()
	y.negative = !y.negative
	return r.add(y)
}

func (r *Real) Mul(x *Real) *Real {
	z := r.mul(x)
	z.round()
	return z
}

func (r *Real) mul(x *Real) *Real {
	// TODO: inline addition here to reduce allocations

	a, b, d := adjust(r, x)

	product := initFrom2(r, x)

	for i := len(a) - 1; i >= 0; i-- {
		z := make([]byte, len(b)+1)
		for j := len(b) - 1; j >= 0; j-- {
			z[j+1] += a[i] * b[j]
			if z[j+1] > 10 {
				z[j] = (z[j+1] / 10)
				z[j+1] = z[j+1] % 10
			}
		}
		shift := len(a) - 1 - i
		pad := make([]byte, shift)
		z = append(z, pad...)
		zr := initFrom(product)
		zr.digits = z
		product = product.add(zr)
	}

	product.decimal = d * 2
	if product.decimal > uint(len(product.digits)) {
		pad := make([]byte, product.decimal-uint(len(product.digits)))
		product.digits = append(pad, product.digits...)
	}
	product.negative = r.negative != x.negative
	product.trim()
	return product
}

func (r *Real) Div(x *Real) *Real {
	z := r.div(x)
	z.round()
	return z
}

func (r *Real) div(x *Real) *Real {
	xr := x.reciprocal()
	return r.mul(xr)
}

func (r *Real) Reciprocal() *Real {
	z := r.reciprocal()
	z.round()
	return z
}
