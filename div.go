package real

func (x *Real) Div(y *Real) *Real {
	x.validate()
	x2 := x.Copy()
	x2.SetPrecision(2 + x.precision)
	y2 := y.Copy()
	y2.SetPrecision(2 + y.precision)
	z := x2.div(y2)
	z.SetPrecision(x.precision)
	return z
}

func (x *Real) div(y *Real) *Real {
	yr := y.reciprocal()
	return x.mul(yr)
}
