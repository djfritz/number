// Copyright 2025 David Fritz. All rights reserved.
// This software may be modified and distributed under the terms of the BSD
// 2-clause license. See the LICENSE file for details.

package number

import "fmt"

// Return the inverse tangent of x, in radians.
func (x *Real) Arctan() *Real {
	x.validate()
	x2 := x.Copy()
	x2.pip(x.precision)
	z := x2.arctan()
	z.SetPrecision(x.precision)
	return z
}

func (x *Real) arctan() *Real {
	if x.IsInf() || x.IsNaN() {
		z := initFrom(x)
		z.form = FormNaN
		return z
	} else if x.IsZero() {
		z := initFrom(x)
		return z
	}

	// arctan(-x) == -arctan(x)
	if x.negative {
		z := x.Abs().arctan()
		z.negative = true
		return z
	}

	// at this point x > 0
	two := initFrom(x)
	two.SetInt64(2)
	one := initFrom(x)
	one.SetInt64(1)

	if x.Compare(NewUint64(1)) == 1 {
		// arctan(x) = π/2 - arctan(1/x)
		pio2 := initFrom(x)
		pio2.significand = make([]byte, len(π))
		copy(pio2.significand, π)
		pio2.round()
		pio2 = pio2.div(two)
		z := pio2.Sub(x.reciprocal().arctan())
		return z
	}

	z := initFrom(x)

	// we set the first sum term to 1 because at n=0, the empty product is
	// 1.
	z.SetInt64(1)

	x2 := x.ipow(2)

	var converged bool
	for i := 1; i < MaxTrigIterations; i++ {
		p := initFrom(x)
		p.SetInt64(1)
		for k := 1; k <= i; k++ {
			kr := initFrom(x)
			kr.SetInt64(int64(k))
			n := two.mul(kr).mul(x2)
			d1 := two.mul(kr).Add(one)
			d2 := one.Add(x2)

			p = p.mul(n.div(d1.mul(d2)))
		}

		zn := z.Add(p)
		if z.Compare(zn) == 0 {
			z = zn
			converged = true
			break
		}
		z = zn
	}

	if !converged {
		panic(fmt.Sprintf("failed to converge sin(%v)", x))
	}

	z = z.mul(x.div(one.Add(x2)))

	return z
}

// Return the 2-argument inverse tangent of x, in radians.
func Atan2(y, x *Real) *Real {
	y.validate()
	x.validate()
	zero := new(Real)
	z := initFrom2(y, x)

	switch x.Compare(zero) {
	case 1:
		z = y.Div(x).Arctan()
	case 0:
		switch y.Compare(zero) {
		case 1:
			two := initFrom(z)
			two.SetInt64(2)
			z.significand = make([]byte, len(π))
			copy(z.significand, π)
			z.round()
			z = z.div(two)
		case 0:
			// undefined
			z.form = FormNaN
		case -1:
			two := initFrom(z)
			two.SetInt64(2)
			z.significand = make([]byte, len(π))
			copy(z.significand, π)
			z.round()
			z = z.div(two)
			z.negative = true
		}
	case -1:
		z = y.Div(x).Arctan()
		pi := initFrom(z)
		pi.significand = make([]byte, len(π))
		copy(pi.significand, π)
		pi.round()
		if y.Compare(zero) == -1 {
			pi.negative = true
		}
		z = z.Add(pi)
	}
	return z
}
