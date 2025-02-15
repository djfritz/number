// Copyright 2024 David Fritz. All rights reserved.
// This software may be modified and distributed under the terms of the BSD
// 2-clause license. See the LICENSE file for details.

package real

import "testing"

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
	y := NewFloat64(1.234)

	z := x.Add(y)
	z.SetPrecision(6)
	expected := NewFloat64(101.234)
	expected.SetPrecision(6)
	if z.Compare(expected) != 0 {
		t.Fatal("invalid add", z, expected)
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
