package real

func (x *Real) Div(y *Real) *Real {
	yr := y.Reciprocal()
	return x.Mul(yr)
}
