package real

func (x *Real) Pow(y *Real) *Real {
	// x^y == e^(y*ln(x))
	a := x.Ln()
	b := y.Mul(a)
	return b.Exp()
}

func (x *Real) ipow(y int) *Real {
	if y < 0 {
		return x.ipow(y * -1).Reciprocal()
	}

	if y == 0 {
		z := initFrom(x)
		z.SetUint64(1)
		return z
	} else if y == 1 {
		return x.Copy()
	} else if y == 2 {
		return x.Mul(x)
	}

	if y%2 == 0 {
		return x.ipow(y / 2).ipow(2)
	} else {
		return x.ipow(y - 1).Mul(x)
	}
}

func (x *Real) Sqrt() *Real {
	half := initFrom(x)
	half.SetUint64(5)
	half.exponent = -1
	return x.Pow(half)
}
