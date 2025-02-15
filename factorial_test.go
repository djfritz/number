// Copyright 2024 David Fritz. All rights reserved.
// This software may be modified and distributed under the terms of the BSD
// 2-clause license. See the LICENSE file for details.

package real

import "testing"

func TestFactorial1(t *testing.T) {
	x := NewUint64(5)
	z := x.Factorial()

	if z.String() != "1.2e2" {
		t.Fatal("invalid factorial", z)
	}
}

func TestFactorial2(t *testing.T) {
	x := NewUint64(20) // largest factorial that can be represented with a uint64
	z := x.Factorial()

	if z.String() != "2.43290200817664e18" {
		t.Fatal("invalid factorial", z)
	}
}

func TestFactorial3(t *testing.T) {
	x := NewUint64(2)
	z := x.Factorial()

	if z.String() != "2e0" {
		t.Fatal("invalid factorial", z)
	}
}

func TestFactorial4(t *testing.T) {
	x := NewUint64(1)
	z := x.Factorial()

	if z.String() != "1e0" {
		t.Fatal("invalid factorial", z)
	}
}

func TestFactorial5(t *testing.T) {
	x := NewUint64(0)
	z := x.Factorial()

	if z.String() != "1e0" {
		t.Fatal("invalid factorial", z)
	}
}
