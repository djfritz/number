package real

// Factorial returns the integer factorial of x. If x is not an integer, the
// integer portion of x is used.
func (x *Real) Factorial() *Real {
	z := initFrom(x)
	z.SetUint64(1)
	if x.Compare(NewUint64(2)) == -1 {
		return z
	}

	i := initFrom(x)
	i.SetUint64(2)
	ipart := x.Integer()
	for i.Compare(ipart) != 1 {
		z = z.Mul(i)
		i = i.Add(NewUint64(1))
	}
	return z
}
