// Copyright 2025 David Fritz. All rights reserved.
// This software may be modified and distributed under the terms of the BSD
// 2-clause license. See the LICENSE file for details.

package number

// A complex number. Internally stored as a two real numbers. The real and
// imaginary parts can have different precision and rounding modes.
type Complex struct {
	r *Real
	i *Real
}

// Copy returns a deep copy of x.
func (x *Complex) Copy() *Complex {
	z := &Complex{
		r: x.r.Copy(),
		i: x.i.Copy(),
	}
	return z

}

func NewComplex(r, i *Real) *Complex {
	z := &Complex{
		r: r,
		i: i,
	}
	z.validate()
	return z
}

func (x *Complex) validate() {
	if x.r == nil && x.i != nil {
		x.r = initFrom(x.i)
	} else if x.i == nil && x.r != nil {
		x.i = initFrom(x.r)
	} else if x.r == nil && x.i == nil {
		x.r = NewUint64(0)
		x.i = NewUint64(0)
	}
}

func (x *Complex) Add(y *Complex) *Complex {
	x.validate()
	y.validate()
	z := &Complex{
		r: x.r.Add(y.r),
		i: x.i.Add(y.i),
	}
	return z
}

func (x *Complex) Sub(y *Complex) *Complex {
	x.validate()
	y.validate()
	z := &Complex{
		r: x.r.Sub(y.r),
		i: x.i.Sub(y.i),
	}
	return z
}

func (x *Complex) Mul(y *Complex) *Complex {
	x.validate()
	y.validate()
	z := &Complex{
		r: x.r.Mul(y.r).Sub(x.i.Mul(y.i)),
		i: x.r.Mul(y.i).Add(x.i.Mul(y.r)),
	}
	return z
}

func (x *Complex) Conjugate() *Complex {
	x.validate()
	z := x.Copy()
	z.i.negative = !z.i.negative
	return z
}

func (x *Complex) Abs() *Real {
	x.validate()
	r2 := x.r.ipow(2)
	y2 := x.i.ipow(2)
	z := r2.Add(y2).Sqrt()
	return z
}

func (x *Complex) Div(y *Complex) *Complex {
	x.validate()
	z := x.Mul(y.Conjugate())
	d := y.Abs().ipow(2)
	z.r = z.r.Div(d)
	z.i = z.i.Div(d)
	return z
}
