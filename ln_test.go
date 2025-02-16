// Copyright 2024 David Fritz. All rights reserved.
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

func BenchmarkLn(b *testing.B) {
	x := new(Real)
	x.significand = []byte{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}
	x.validate()
	for b.Loop() {
		x.Ln()
	}
}
