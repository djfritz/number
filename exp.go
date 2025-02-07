package real

import "fmt"

func (r *Real) Exp() *Real {
	z := r.Copy()
	z.SetUint64(0)

	// TODO: escape logic
	for i := uint64(0); i < 50; i++ {
		n := r.ipow(i)
		d := factorial(i)
		q := n.Div(d)
		z = z.Add(q)

		fmt.Println(n, d, q)
	}
	z.fix()
	return z
}

func (r *Real) ipow(x uint64) *Real {
	z := r.Copy()
	z.SetUint64(1)
	if x == 0 {
		return z
	} else if x == 1 {
		return r.Copy()
	}

	z = r.Copy()
	for i := uint64(1); i < x; i++ {
		z = z.Mul(r)
	}
	z.fix()
	return z
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
	z.fix()
	return z
}
