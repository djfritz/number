// Copyright 2025 David Fritz. All rights reserved.
// This software may be modified and distributed under the terms of the BSD
// 2-clause license. See the LICENSE file for details.

package number

// Return the power of y and base x (x^y).
func (x *Real) Pow(y *Real) *Real {
	x.validate()
	y.validate()
	x2 := x.Copy()
	y2 := y.Copy()
	x2.pip(x.precision)
	y2.pip(y.precision)
	z := x2.pow(y2)
	z.SetPrecision(x.precision)
	return z
}

func (x *Real) pow(y *Real) *Real {
	// Exponentiation has a lot of edge cases around infinity.
	if x.IsNaN() || y.IsNaN() {
		z := initFrom2(x, y)
		z.form = FormNaN
		return z
	} else if x.IsInf() && y.IsInf() {
		// inf^inf == NaN
		z := initFrom2(x, y)
		z.form = FormNaN
		return z
	} else if y.IsZero() {
		// z^0 == 1
		z := initFrom2(x, y)
		z.SetUint64(1)
		return z
	} else if x.IsZero() {
		z := initFrom2(x, y)
		return z
	} else if y.Compare(NewUint64(1)) == 0 {
		return x.Copy()
	} else if x.IsInf() && y.Abs().Compare(NewUint64(1)) >= 0 {
		// inf^Z == inf for all other |Z| >=1
		z := initFrom2(x, y)
		z.form = FormInf
		z.negative = x.negative
		return z
	} else if x.Compare(NewUint64(1)) == 0 && y.IsInf() {
		// a^inf == NaN for a == 1
		z := initFrom2(x, y)
		z.form = FormNaN
		return z
	} else if x.Compare(NewUint64(1)) == 1 && y.IsInf() && !y.negative {
		// a^inf == inf for a > 1
		z := initFrom2(x, y)
		z.form = FormInf
		return z
	} else if x.Compare(NewUint64(1)) == 1 && y.IsInf() && y.negative {
		// a^-inf == 0 for a > 1
		z := initFrom2(x, y)
		return z
	} else if x.Compare(NewUint64(1)) == -1 && x.Compare(new(Real)) == 1 && y.IsInf() {
		// a^±inf == 0 for 0 < a < 1
		z := initFrom2(x, y)
		return z
	} else if x.negative && y.IsInf() {
		// a^±inf == NaN for a < 0
		z := initFrom2(x, y)
		z.form = FormNaN
		return z
	} else if x.negative && y.Abs().Compare(NewUint64(1)) == -1 {
		z := initFrom2(x, y)
		z.form = FormNaN
		return z
	}

	// Positive integer exponents can be calculated faster by decomposing
	// the exponent. Additionally it allows for things like -3^2.
	if y.IsInteger() {
		if y.negative {
			yi := y.Copy()
			yi.negative = false
			return x.pow(yi).reciprocal()
		} else {
			two := initFrom2(x, y)
			two.SetUint64(2)
			if y.Compare(two) == 0 {
				return x.mul(x)
			} else if y.mod(two).Compare(new(Real)) == 0 {
				return x.pow(y.div(two).Integer()).pow(two)
			} else {
				return x.Pow(y.Sub(NewUint64(1))).mul(x)
			}
		}
	}

	p := umax(x.precision, y.precision)

	// x^y == e^(y*ln(x))
	x2 := x.Copy()
	x2.pip(p)
	y2 := y.Copy()
	y2.pip(p)

	z := y2.mul(x2.ln()).exp()
	z.SetPrecision(p)

	return z
}

func (x *Real) ipow(y int) *Real {
	x.validate()
	if y < 0 {
		return x.ipow(y * -1).reciprocal()
	}

	if y == 0 {
		z := initFrom(x)
		z.SetUint64(1)
		return z
	} else if y == 1 {
		return x.Copy()
	} else if y == 2 {
		return x.mul(x)
	}

	if y%2 == 0 {
		return x.ipow(y / 2).ipow(2)
	} else {
		return x.ipow(y - 1).mul(x)
	}
}

// Return the square root of x.
func (x *Real) Sqrt() *Real {
	x.validate()

	x2 := x.Copy()
	x2.pip(x.precision)

	half := initFrom(x2)
	half.SetUint64(5)
	half.exponent = -1
	z := x2.pow(half)
	z.SetPrecision(x.precision)
	return z
}
