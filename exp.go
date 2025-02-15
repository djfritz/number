// Copyright 2024 David Fritz. All rights reserved.
// This software may be modified and distributed under the terms of the BSD
// 2-clause license. See the LICENSE file for details.

package realnumber

import "fmt"

const MaxExpIterations = 1000

// Return the exponential of x (e^x).
func (x *Real) Exp() *Real {
	x.validate()
	x2 := x.Copy()
	x2.SetPrecision(internalPrecisionBuffer + x.precision)
	z := x2.exp()
	z.SetPrecision(x.precision)
	return z
}

func (x *Real) exp() *Real {
	// we decompose e^x using associativity to get x into a normalized
	// (1<x<10) value and compute the power series from there. This adds
	// multiplies, but helps with convergence for large values of x.
	if x.negative {
		// e^-x == 1/e^x
		xcopy := x.Copy()
		xcopy.negative = false
		return xcopy.exp().reciprocal()
	} else if x.exponent > 0 {
		xcopy := x.Copy()
		xcopy.exponent--
		z := xcopy.exp().ipow(10)
		return z
	} else if x.exponent < 0 {
		// we'll use e^.x == e^1.x * e^-1
		xcopy := x.Copy()
		xcopy = xcopy.Add(NewInt64(1))
		z := xcopy.exp()
		e1 := initFrom(z)
		e1.SetInt64(-1)
		e1 = e1.exp()
		z = z.mul(e1)
		return z
	}

	// at this point 1<x<10

	z := initFrom(x)

	xscaled := x.Copy()
	xscaled.exponent = 0

	var converged bool
	for i := 0; i < MaxExpIterations; i++ {
		n := xscaled.ipow(i)
		d := initFrom(xscaled)
		d.SetUint64(uint64(i))
		d = d.Factorial()
		q := n.div(d)
		zn := z.Add(q)
		if z.Compare(zn) == 0 {
			z = zn
			converged = true
			break
		}
		z = zn
	}
	if !converged {
		panic(fmt.Sprintf("failed to converge exp(%v)", x))
	}

	return z
}
