// Copyright 2025 David Fritz. All rights reserved.
// This software may be modified and distributed under the terms of the BSD
// 2-clause license. See the LICENSE file for details.

package number

import (
	"strconv"
)

// Return the reciprocal of x.
func (x *Real) Reciprocal() *Real {
	x.validate()
	x2 := x.Copy()
	x2.pip(x.precision)
	z := x2.reciprocal()
	z.SetPrecision(x.precision)
	return z
}

func (x *Real) reciprocal() *Real {
	if x.IsInf() {
		z := initFrom(x)
		return z
	} else if x.IsNaN() {
		z := initFrom(x)
		z.form = FormNaN
		return z
	} else if x.IsZero() {
		z := initFrom(x)
		z.form = FormInf
		z.negative = x.negative
		return z
	} else if x.Compare(NewUint64(1)) == 0 {
		z := initFrom(x)
		z.SetUint64(1)
		return z
	}

	xscaled := x.Copy()
	xscaled.exponent = 0

	s := xscaled.String()
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic("could not parse float")
	}
	f = 1 / f

	z0 := initFrom(x)
	z0.SetFloat64(f)

	z := z0
	two := initFrom(x)
	two.SetInt64(2)

	for i := 0; i < estimateConvergence(float64MinimumDecimalPrecision, x.precision); i++ {
		zn := z.mul(two.Sub(xscaled.mul(z)))
		z = zn
	}

	z.exponent += x.exponent * -1
	z.round()
	return z
}
