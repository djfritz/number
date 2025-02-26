// Copyright 2025 David Fritz. All rights reserved.
// This software may be modified and distributed under the terms of the BSD
// 2-clause license. See the LICENSE file for details.

package number

func (x *Real) Remainder(y *Real) *Real {
	x.validate()
	x2 := x.Copy()
	x2.pip(x.precision)
	y2 := y.Copy()
	y2.pip(y.precision)
	z := x2.remainder(y2)
	z.SetPrecision(x.precision)
	return z
}

func (x *Real) remainder(y *Real) *Real {
	// r = x - y*round(x/y)
	z := initFrom2(x, y)
	if x.IsInf() || y.IsInf() {
		z.form = FormNaN
		return z
	} else if x.IsNaN() || y.IsNaN() {
		z.form = FormNaN
		return z
	} else if x.IsZero() {
		return z
	} else if y.IsZero() {
		z.form = FormInf
		return z
	}

	rd := x.Div(y) // do the full Div() function here to make sure we expand and round first
	rd = rd.Integer()

	z = x.Sub(y.mul(rd))

	return z
}
