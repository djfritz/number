// Copyright 2025 David Fritz. All rights reserved.
// This software may be modified and distributed under the terms of the BSD
// 2-clause license. See the LICENSE file for details.

package number

import "testing"

func TestDiv1(t *testing.T) {
	x := NewInt64(10)
	y := NewInt64(5)

	q := x.Div(y)

	if q.Compare(NewInt64(2)) != 0 {
		t.Fatal("invalid div", q)
	}
}

func TestDiv2(t *testing.T) {
	x := NewInt64(2)
	y := NewInt64(50)

	q := x.Div(y)

	expected := NewInt64(4)
	expected.exponent = -2
	if q.Compare(expected) != 0 {
		t.Fatal("invalid div", q)
	}
}

func TestDiv3(t *testing.T) {
	x := NewInt64(23)
	y := NewInt64(5011513)

	q := x.Div(y)

	if q.String() != "4.589432373017889008768409859457613e-6" {
		t.Fatal("invalid div", q)
	}
}

func TestDivInf(t *testing.T) {
	x := new(Real)
	x.form = FormInf
	y := NewInt64(-5)

	z := x.Div(y)
	if z.String() != "∞" {
		t.Fatal("invalid div", z)
	}
}

func TestDivNegInf(t *testing.T) {
	x := new(Real)
	x.form = FormInf
	x.negative = true
	y := NewInt64(-5)

	z := x.Div(y)
	if z.String() != "-∞" {
		t.Fatal("invalid div", z)
	}
}

func TestDivInf2(t *testing.T) {
	y := new(Real)
	y.form = FormInf
	x := NewInt64(5)

	z := x.Div(y)
	if z.String() != "0" {
		t.Fatal("invalid div", z)
	}
}

func TestDivNegInf2(t *testing.T) {
	y := new(Real)
	y.form = FormInf
	y.negative = true
	x := NewInt64(5)

	z := x.Div(y)
	if z.String() != "0" {
		t.Fatal("invalid div", z)
	}
}

func TestDivBothInf(t *testing.T) {
	x := new(Real)
	x.form = FormInf
	x.negative = true
	y := new(Real)
	y.form = FormInf
	y.negative = false

	z := x.Div(y)
	if z.String() != "NaN" {
		t.Fatal("invalid div", z)
	}
}

func TestDivNaN(t *testing.T) {
	x := new(Real)
	x.form = FormNaN
	y := NewUint64(1)

	z := x.Div(y)
	if z.String() != "NaN" {
		t.Fatal("invalid div", z)
	}
}

func TestDivZero(t *testing.T) {
	x := NewUint64(1)
	y := new(Real)

	z := x.Div(y)
	if z.String() != "∞" {
		t.Fatal("invalid div", z)
	}
}

func BenchmarkDiv(b *testing.B) {
	x := new(Real)
	y := new(Real)
	x.significand = []byte{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}
	y.significand = []byte{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}
	x.validate()
	y.validate()
	for b.Loop() {
		x.Div(y)
	}
}

func TestMod1(t *testing.T) {
	x := NewInt64(23)
	y := NewInt64(2)

	m := x.Mod(y)

	if m.String() != "1e0" {
		t.Fatal("invalid mod", m)
	}
}

func TestMod2(t *testing.T) {
	x := NewInt64(23)
	y := NewInt64(8)

	m := x.Mod(y)

	if m.String() != "7e0" {
		t.Fatal("invalid mod", m)
	}
}

func TestModInf(t *testing.T) {
	x := new(Real)
	x.form = FormInf
	y := NewInt64(8)

	m := x.Mod(y)

	if m.String() != "NaN" {
		t.Fatal("invalid mod", m)
	}
}
