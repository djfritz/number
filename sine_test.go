// Copyright 2025 David Fritz. All rights reserved.
// This software may be modified and distributed under the terms of the BSD
// 2-clause license. See the LICENSE file for details.

package number

import "testing"

func TestSine1(t *testing.T) {
	x := NewUint64(5)
	z := x.Sin()

	if z.String() != "-9.58924274663138468893154406155994e-1" {
		t.Fatal("invalid sin", z.String())
	}
}

func TestSine2(t *testing.T) {
	x := NewUint64(5000)
	z := x.Sin()

	if z.String() != "-9.879664387667768472476423570861315e-1" {
		t.Fatal("invalid sin", z)
	}
}

func TestSine3(t *testing.T) {
	x := NewUint64(1)
	x.exponent = 22
	x.SetPrecision(50)
	z := x.Sin()

	if z.String() != "-8.522008497671888017727058937530294e-1" {
		t.Fatal("invalid sin", z)
	}
}

func TestSine4(t *testing.T) {
	x := NewUint64(0)
	z := x.Sin()

	if z.String() != "0" {
		t.Fatal("invalid sin", z)
	}
}

// sin(2π) == 0, but even the intel decimal arithmetic library gives us 2.316e-34.
func TestSine5(t *testing.T) {
	x, _ := ParseReal("6.28318530717958647692528676655900576839433879875021164194988918461563281257", DefaultPrecision) // 2π
	z := x.Sin()

	if z.String() != "2.316e-34" {
		t.Fatal("invalid sin", z.String())
	}
}
