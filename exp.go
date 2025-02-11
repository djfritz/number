package real

const MaxExpIterations = 1000

func (x *Real) Exp() *Real {
	z := initFrom(x)

	for i := 0; i < MaxExpIterations; i++ {
		n := x.ipow(i)
		d := initFrom(x)
		d.SetUint64(uint64(i))
		d = d.Factorial()
		q := n.Div(d)
		zn := z.Add(q)
		if z.Compare(zn) == 0 {
			z = zn
			break
		}
		z = zn
	}
	return z
}
