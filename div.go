// Copyright 2024 David Fritz. All rights reserved.
// This software may be modified and distributed under the terms of the BSD
// 2-clause license. See the LICENSE file for details.

package realnumber

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
