// Copyright 2025 David Fritz. All rights reserved.
// This software may be modified and distributed under the terms of the BSD
// 2-clause license. See the LICENSE file for details.

package number

import (
	"testing"
)

func TestNewComplex(t *testing.T) {
	r := NewUint64(5)
	i := NewUint64(10)

	z := NewComplex(r, i)

	if z.r.Compare(r) != 0 {
		t.Fatal("invalid real part")
	} else if z.i.Compare(i) != 0 {
		t.Fatal("invalid imaginary part")
	}
}

func TestComplexAdd(t *testing.T) {
	a := &Complex{
		r: NewUint64(5),
		i: NewInt64(-10),
	}
	b := &Complex{
		r: NewUint64(100),
		i: NewInt64(10),
	}
	expected := &Complex{
		r: NewUint64(105),
		i: NewInt64(0),
	}
	z := a.Add(b)

	if z.r.Compare(expected.r) != 0 {
		t.Fatal("invalid real part")
	} else if z.i.Compare(expected.i) != 0 {
		t.Fatal("invalid imaginary part")
	}
}

func TestComplexSub(t *testing.T) {
	a := &Complex{
		r: NewUint64(5),
		i: NewInt64(-10),
	}
	b := &Complex{
		r: NewUint64(100),
		i: NewInt64(10),
	}
	expected := &Complex{
		r: NewInt64(-95),
		i: NewInt64(-20),
	}
	z := a.Sub(b)

	if z.r.Compare(expected.r) != 0 {
		t.Fatal("invalid real part")
	} else if z.i.Compare(expected.i) != 0 {
		t.Fatal("invalid imaginary part")
	}
}

func TestComplexMul(t *testing.T) {
	a := &Complex{
		r: NewUint64(5),
		i: NewInt64(-10),
	}
	b := &Complex{
		r: NewUint64(100),
		i: NewInt64(10),
	}
	expected := &Complex{
		r: NewInt64(600),
		i: NewInt64(-950),
	}
	z := a.Mul(b)

	if z.r.Compare(expected.r) != 0 {
		t.Fatal("invalid real part")
	} else if z.i.Compare(expected.i) != 0 {
		t.Fatal("invalid imaginary part")
	}
}

func TestComplexDiv(t *testing.T) {
	a := &Complex{
		r: NewUint64(5),
		i: NewInt64(-10),
	}
	b := &Complex{
		r: NewUint64(100),
		i: NewInt64(10),
	}
	r, _ := ParseReal("0.03960396039603960396039603960396028", DefaultPrecision)
	i, _ := ParseReal("-0.1039603960396039603960396039603957", DefaultPrecision)
	expected := &Complex{
		r: r,
		i: i,
	}
	z := a.Div(b)

	if z.r.Compare(expected.r) != 0 {
		t.Fatal("invalid real part", z.r, expected.r)
	} else if z.i.Compare(expected.i) != 0 {
		t.Fatal("invalid imaginary part", z.i)
	}
}

func TestComplexPolar1(t *testing.T) {
	x := NewComplex(NewInt64(1), NewInt64(1))
	ρ, φ := x.Polar()

	if ρ.String() != "1.414213562373095048801688724209698e0" {
		t.Fatal("invalid polar radius", ρ.String())
	}
	if φ.String() != "7.853981633974483096156608458198757e-1" {
		t.Fatal("invalid polar angle", φ.String())
	}
}

func TestComplexPolar2(t *testing.T) {
	x := NewComplex(NewInt64(-123), NewInt64(-2))
	ρ, φ := x.Polar()

	if ρ.String() != "1.230162590879758463158120036190364e2" {
		t.Fatal("invalid polar radius", ρ.String())
	}
	if φ.String() != "-3.125333923784663650770045485291427e0" {
		t.Fatal("invalid polar angle", φ.String())
	}
}
