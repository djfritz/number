package real

func (x *Real) Compare(y *Real) int {
	if !x.negative && y.negative {
		return 1
	} else if x.negative && !y.negative {
		return -1
	}

	p := x.precision
	if y.precision < p {
		p = y.precision
	}

	a, b, _ := adjust(x, y)

	for i := range a {
		if a[i] > b[i] {
			if x.negative {
				return -1
			}
			return 1
		} else if a[i] < b[i] {
			if x.negative {
				return 1
			}
			return -1
		}
	}

	return 0
}

func compareSignificand(a, b []byte) int {
	if len(a) != len(b) {
		sc := len(a)
		if len(b) < sc {
			sc = len(b)
		}
		c := compareSignificand(a[:sc], b[:sc])
		if c != 0 {
			return c
		}

		// A normal significand never has trailing zeros, so we just
		// check length here.
		if len(a) > len(b) {
			return 1
		} else {
			return -1
		}
	}
	for i := range a {
		if a[i] > b[i] {
			return 1
		} else if a[i] < b[i] {
			return -1
		}
	}

	return 0
}
