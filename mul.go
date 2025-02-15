// Copyright 2024 David Fritz. All rights reserved.
// This software may be modified and distributed under the terms of the BSD
// 2-clause license. See the LICENSE file for details.

package real

// Return the product of x and y.
func (x *Real) Mul(y *Real) *Real {
	x.validate()
	y.validate()
	x2 := x.Copy()
	x2.SetPrecision(2 + x.precision)
	y2 := y.Copy()
	y2.SetPrecision(2 + y.precision)
	z := x2.mul(y2)
	z.SetPrecision(x.precision)
	return z
}

func (x *Real) mul(y *Real) *Real {
	z := initFrom2(x, y)
	for i := len(x.significand) - 1; i >= 0; i-- {
		p := make([]byte, len(y.significand)+1)
		for j := len(y.significand) - 1; j >= 0; j-- {
			p[j+1] += x.significand[i] * y.significand[j]
			if p[j+1] >= 10 {
				p[j] = p[j+1] / 10
				p[j+1] = p[j+1] % 10
			}
		}
		zr := initFrom(z)
		zr.exponent = 1 - i
		zr.significand = p
		zr.round()
		zn := z.Add(zr)
		z = zn
	}

	z.exponent += x.exponent + y.exponent
	if x.negative != y.negative {
		z.negative = true
	}
	z.round()
	return z
}
