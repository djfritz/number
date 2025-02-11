package real

func (x *Real) Add(y *Real) *Real {
	a, b, e := adjust(x, y)
	a = append([]byte{0}, a...)
	b = append([]byte{0}, b...)
	e++

	z := initFrom(x)
	z.exponent = e

	if x.negative == y.negative {
		z.negative = x.negative

		for i := len(a) - 1; i >= 0; i-- {
			a[i] += b[i]
			if a[i] >= 10 {
				a[i] -= 10
				a[i-1]++
			}
		}
	} else {
		z.negative = x.negative
		if compareSignificand(a, b) == -1 {
			a, b = b, a
			z.negative = !z.negative
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

	z.significand = a
	z.round()

	return z
}
