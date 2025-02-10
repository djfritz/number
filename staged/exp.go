package real

func (r *Real) Exp() *Real {
	z := r.Copy()
	z.SetUint64(0)
	for i := int64(0); i < 100; i++ {
		n := r.ipow(i)
		d := factorial(uint64(i))
		d.precision = r.precision
		q := n.div(d)
		z = z.add(q)
		z.roundTo(2 * z.precision)
	}
	z.round()
	return z
}

func (r *Real) ipow(x int64) *Real {
	if x < 0 {
		return r.ipow(x * -1).Reciprocal()
	}
	z := r.Copy()
	z.SetUint64(1)
	if x == 0 {
		return z
	} else if x == 1 {
		return r.Copy()
	} else if x == 2 {
		return r.Mul(r)
	}

	if x%2 == 0 {
		return r.ipow(x / 2).ipow(2)
	} else {
		return r.ipow(x - 1).Mul(r)
	}
}

func factorial(x uint64) *Real {
	z := new(Real)
	z.SetUint64(1)
	if x < 2 {
		return z
	}

	m := new(Real)
	for i := uint64(2); i <= x; i++ {
		m.SetUint64(i)
		z = z.Mul(m)
	}
	return z
}

func (r *Real) Ln() *Real {
	// z1 = z0 * 2*((x-exp(z0))/(x+exp(z0)))

	z := initFrom(r)
	z.SetUint64(1) // TODO: linear approximation

	two := initFrom(r)
	two.SetUint64(2)

	for i := 0; i < 100; i++ {
		ez := z.Exp()
		n := r.sub(ez)
		d := r.add(ez)
		q := n.div(d)
		q2 := two.mul(q)
		znext := z.add(q2)
		znext.roundTo(2 * znext.precision)
		if znext.Compare(z) == 0 {
			z = znext
			break
		}
		z = znext
	}
	z.round()
	return z
}
