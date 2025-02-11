package real

func (x *Real) Mul(y *Real) *Real {
	z := initFrom2(x, y)

	for i := len(x.significand) - 1; i >= 0; i-- {
		p := make([]byte, len(y.significand)+1)
		for j := len(y.significand) - 1; j >= 0; j-- {
			p[j+1] += x.significand[i] * y.significand[j]
			if p[j+1] >= 10 {
				p[j] = p[j+1] / 10
				p[j+1] = p[j+1] % 10
			}
		}
		zr := initFrom(z)
		zr.exponent = 1 - i
		zr.significand = p
		zr.trim()
		z = z.Add(zr)
	}

	z.exponent += x.exponent + y.exponent
	if x.negative != y.negative {
		z.negative = true
	}
	z.round()
	return z
}
