// Copyright 2024 David Fritz. All rights reserved.
// This software may be modified and distributed under the terms of the BSD
// 2-clause license. See the LICENSE file for details.

package realnumber

// Factorial returns the integer factorial of x. If x is not an integer, the
// integer portion of x is used.
func (x *Real) Factorial() *Real {
	x.validate()
	z := initFrom(x)
	z.SetUint64(1)
	if x.Compare(NewUint64(2)) == -1 {
		return z
	}

	i := initFrom(x)
	i.SetUint64(2)
	ipart := x.Integer()
	for i.Compare(ipart) != 1 {
		z = z.mul(i)
		i = i.Add(NewUint64(1))
	}
	return z
}
