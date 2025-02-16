// Copyright 2024 David Fritz. All rights reserved.
// This software may be modified and distributed under the terms of the BSD
// 2-clause license. See the LICENSE file for details.

package number

import "testing"

func TestMul1(t *testing.T) {
	x := NewInt64(5)
	y := NewInt64(12)

	z := x.Mul(y)
	if z.Compare(NewInt64(60)) != 0 {
		t.Fatal("invalid mul", z)
	}
}

func TestMul2(t *testing.T) {
	x := NewInt64(12)
	y := NewInt64(5)

	z := x.Mul(y)
	if z.Compare(NewInt64(60)) != 0 {
		t.Fatal("invalid mul", z)
	}
}

func TestMul3(t *testing.T) {
	x := NewInt64(-12)
	y := NewInt64(5)

	z := x.Mul(y)
	if z.Compare(NewInt64(-60)) != 0 {
		t.Fatal("invalid mul", z)
	}
}

func TestMul4(t *testing.T) {
	x := NewInt64(15)
	x.exponent = 0
	y := NewInt64(5)

	z := x.Mul(y)
	if z.Compare(NewFloat64(7.5)) != 0 {
		t.Fatal("invalid mul", z)
	}
}

func TestMul5(t *testing.T) {
	x := NewInt64(5)
	y := NewInt64(4)

	z := x.Mul(y)
	if z.Compare(NewInt64(20)) != 0 {
		t.Fatal("invalid mul", z)
	}
}

func TestMul6(t *testing.T) {
	x := new(Real)
	x.significand = []byte{8, 1, 0, 3, 7, 2, 7, 7, 1, 4, 7, 4, 8, 7, 8, 4, 0, 6}
	x.exponent = -1
	y := new(Real)
	y.significand = []byte{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 6, 9, 9, 6}

	z := x.Mul(y)
	if z.String() != "8.103727714748784440842787682333856e-1" {
		t.Fatal("invalid mul", z)
	}
}

func BenchmarkMul(b *testing.B) {
	x := new(Real)
	y := new(Real)
	x.significand = []byte{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}
	y.significand = []byte{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}
	x.validate()
	y.validate()
	for b.Loop() {
		x.Mul(y)
	}
}
