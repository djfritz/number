// Copyright 2025 David Fritz. All rights reserved.
// This software may be modified and distributed under the terms of the BSD
// 2-clause license. See the LICENSE file for details.

package number

import (
	"testing"
)

func TestAdd1(t *testing.T) {
	x := NewInt64(10)
	y := NewInt64(5)

	z := x.Add(y)
	if z.Compare(NewInt64(15)) != 0 {
		t.Fatal("invalid add", z)
	}
}

func TestAdd2(t *testing.T) {
	x := NewInt64(10)
	y := NewInt64(-5)

	z := x.Add(y)
	if z.Compare(NewInt64(5)) != 0 {
		t.Fatal("invalid add", z)
	}
}

func TestAdd3(t *testing.T) {
	x := NewInt64(100)
	y := NewInt64(1234)
	y.exponent = 0

	z := x.Add(y)
	if z.String() != "1.01234e2" {
		t.Fatal("invalid add", z)
	}
}

func TestAdd4(t *testing.T) {
	x := NewInt64(-100)
	y := NewInt64(-5)

	z := x.Add(y)
	if z.Compare(NewInt64(-105)) != 0 {
		t.Fatal("invalid add", z)
	}
}

func TestAddInf(t *testing.T) {
	x := new(Real)
	x.form = FormInf
	y := NewInt64(-5)

	z := x.Add(y)
	if z.String() != "∞" {
		t.Fatal("invalid add", z)
	}
}

func TestAddNegInf(t *testing.T) {
	x := new(Real)
	x.form = FormInf
	x.negative = true
	y := NewInt64(-5)

	z := x.Add(y)
	if z.String() != "-∞" {
		t.Fatal("invalid add", z)
	}
}

func TestAddBothInf(t *testing.T) {
	x := new(Real)
	x.form = FormInf
	x.negative = true
	y := new(Real)
	y.form = FormInf
	y.negative = false

	z := x.Add(y)
	if z.String() != "NaN" {
		t.Fatal("invalid add", z)
	}
}

func TestAddNaN(t *testing.T) {
	x := new(Real)
	x.form = FormNaN
	y := NewUint64(1)

	z := x.Add(y)
	if z.String() != "NaN" {
		t.Fatal("invalid add", z)
	}
}

func TestSub1(t *testing.T) {
	x := NewInt64(-100)
	y := NewInt64(-5)

	z := x.Sub(y)
	if z.Compare(NewInt64(-95)) != 0 {
		t.Fatal("invalid sub", z)
	}
}

func TestSub2(t *testing.T) {
	x := NewInt64(1)
	y := NewInt64(5)

	z := x.Sub(y)
	if z.Compare(NewInt64(-4)) != 0 {
		t.Fatal("invalid sub", z)
	}
}

func TestSub3(t *testing.T) {
	x := NewInt64(2)
	y := NewInt64(15)
	y.exponent = 0

	z := x.Sub(y)
	if z.String() != "5e-1" {
		t.Fatal("invalid sub", z)
	}
}

func BenchmarkAdd(b *testing.B) {
	x := new(Real)
	y := new(Real)
	x.significand = []byte{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}
	y.significand = []byte{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}
	x.validate()
	y.validate()
	for b.Loop() {
		x.Add(y)
	}
}

func TestAddRoundEven(t *testing.T) {
	x, _ := ParseReal("0.4444444444", DefaultPrecision)
	y, _ := ParseReal("0.5555555555", DefaultPrecision)
	x.SetMode(ModeNearest)
	y.SetMode(ModeNearest)

	z := x.Add(y)
	z.SetPrecision(9)
	if z.Compare(NewInt64(1)) != 0 {
		t.Fatal("invalid add", z)
	}
}

func TestAddRoundEven2(t *testing.T) {
	x, _ := ParseReal("12e3", DefaultPrecision)
	y, _ := ParseReal("3446", DefaultPrecision)
	x.SetMode(ModeNearest)
	y.SetMode(ModeNearest)
	x.SetPrecision(3)
	y.SetPrecision(3)

	z := x.Add(y)
	if z.Compare(NewInt64(15500)) != 0 {
		t.Fatal("invalid add", z)
	}
}
