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

func (r *Real) fix() {
	if r.precision == 0 {
		r.precision = 64
	}
	// TODO: other rounding modes
	if r.decimal <= r.precision {
		return
	}

	trim := r.decimal - r.precision
	rd := r.digits[uint(len(r.digits))-trim]
	r.digits = r.digits[:uint(len(r.digits))-trim]
	r.decimal = r.precision

	if rd < 5 {
		// round down
	} else if rd > 5 {
		// round up
		r.digits[len(r.digits)-1]++
	} else { // round to nearest even
		if r.digits[len(r.digits)-1]%2 != 0 {
			r.digits[len(r.digits)-1]++
		}
	}
}

func (r *Real) trim() {
	r.fix()
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

func (r *Real) Abs() *Real {
	z := r.Copy()
	z.negative = false
	return z
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

	// TODO: rounding

	return z

}

func (r *Real) Sub(x *Real) *Real {
	y := x.Copy()
	y.negative = !y.negative
	return r.Add(y)
}

func (r *Real) Mul(x *Real) *Real {
	// TODO: inline addition here to reduce allocations

	a, b, d := adjust(r, x)

	product := new(Real)

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
		zr := &Real{digits: z}
		product = product.Add(zr)
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
	xr := x.Reciprocal()
	return r.Mul(xr)
}

func (r *Real) Reciprocal() *Real {
	// scale r to be 0 < r < 1
	rc := r.Copy()
	dshift := uint(len(r.digits)) - r.decimal
	rc.decimal = uint(len(rc.digits))

	five := new(Real)
	five.SetUint64(5)
	four := new(Real)
	four.SetUint64(4)
	x0 := four.Sub(five.Mul(rc))
	x0.SetUint64(2)

	// f(x) = (1/x) - D
	// xi+1 = xi(2-(D*xi))

	x := x0
	two := new(Real)
	two.SetUint64(2)
	for i := 0; i < 10; i++ { // TODO: precision escape
		xn := x.Mul(two.Sub(rc.Mul(x)))
		if xn.Compare(x) == 0 {
			x = xn
			break
		}
		x = xn
	}

	// restore shift
	x.decimal += dshift
	if x.decimal > uint(len(x.digits)) {
		pad := make([]byte, x.decimal-uint(len(x.digits)))
		x.digits = append(pad, x.digits...)
	}

	x.trim()

	return x
}
