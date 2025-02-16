// Copyright 2024 David Fritz. All rights reserved.
// This software may be modified and distributed under the terms of the BSD
// 2-clause license. See the LICENSE file for details.

package number

// Return the quotient of x/y. Precision and rounding rules are the same as
// Add().
func (x *Real) Div(y *Real) *Real {
	x.validate()
	x2 := x.Copy()
	x2.SetPrecision(internalPrecisionBuffer + x.precision)
	y2 := y.Copy()
	y2.SetPrecision(internalPrecisionBuffer + y.precision)
	z := x2.div(y2)
	z.SetPrecision(x.precision)
	return z
}

func (x *Real) div(y *Real) *Real {
	yr := y.reciprocal()
	return x.mul(yr)
}

// Return the modulus x%y. If either x or y are not integers, they will be
// truncated before the operation.
func (x *Real) Mod(y *Real) *Real {
	xi := x.Integer()
	yi := y.Integer()

	m := xi.Sub(xi.Div(yi).Floor().Mul(yi))
	return m
}
