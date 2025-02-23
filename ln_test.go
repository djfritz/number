// Copyright 2025 David Fritz. All rights reserved.
// This software may be modified and distributed under the terms of the BSD
// 2-clause license. See the LICENSE file for details.

package number

import (
	"testing"
)

func TestLn1(t *testing.T) {
	x := NewUint64(5)
	z := x.Ln()

	if z.String() != "1.609437912434100374600759333226188e0" {
		t.Fatal("invalid ln", z)
	}
}

func TestLn2(t *testing.T) {
	x := NewFloat64(2)
	z := x.Ln()

	if z.String() != "6.931471805599453094172321214581766e-1" {
		t.Fatal("invalid ln", z)
	}
}

func TestLn3(t *testing.T) {
	x := NewInt64(2)
	x.exponent = -3
	z := x.Ln()

	if z.String() != "-6.214608098422191742636742242594916e0" {
		t.Fatal("invalid ln", z)
	}
}

func TestLn4(t *testing.T) {
	x := NewUint64(1)
	z := x.Ln()

	if z.String() != "0" {
		t.Fatal("invalid ln", z)
	}
}

func TestLnInf(t *testing.T) {
	x := new(Real)
	x.form = FormInf
	z := x.Ln()

	if z.String() != "∞" {
		t.Fatal("invalid ln", z)
	}
}

func TestLnNaN(t *testing.T) {
	x := new(Real)
	x.form = FormNaN
	z := x.Ln()

	if z.String() != "NaN" {
		t.Fatal("invalid ln", z)
	}
}

func TestLnNeg(t *testing.T) {
	x := NewInt64(-1)
	z := x.Ln()

	if z.String() != "NaN" {
		t.Fatal("invalid ln", z)
	}
}

func TestLnZero(t *testing.T) {
	x := new(Real)
	z := x.Ln()

	if z.String() != "-∞" {
		t.Fatal("invalid ln", z)
	}
}

func TestLnAllTheWayDown(t *testing.T) {
	x := NewUint64(2)
	z := x.Ln()
	z = z.Ln()
	z = z.Ln()

	if z.String() != "NaN" {
		t.Fatal("invalid ln", z)
	}
}

// The natural log of the number in x below is, to far more than 17 digits:
// -8.999999999999987849999179874970475500000000085410167440336139577720... ×10^-8
// The dectest suite has a note about this particular test being a ">.5ulp
// case", but we calculate this to better than .5ulp with the default internal
// precision.
func TestLnDecTestLn116(t *testing.T) {
	x, _ := ParseReal("0.99999991000000405", DefaultPrecision)
	x.SetMode(ModeNearest)
	x.SetPrecision(17)
	z := x.Ln()

	if z.String() != "-8.9999999999999878e-8" { // original test has this ending in 9879
		t.Fatal("invalid ln", z)
	}
}

func BenchmarkLn(b *testing.B) {
	x := new(Real)
	x.significand = []byte{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}
	x.validate()
	for b.Loop() {
		x.Ln()
	}
}
