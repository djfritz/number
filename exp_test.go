// Copyright 2025 David Fritz. All rights reserved.
// This software may be modified and distributed under the terms of the BSD
// 2-clause license. See the LICENSE file for details.

package number

import "testing"

func TestExp1(t *testing.T) {
	x := NewUint64(5)
	z := x.Exp()

	if z.String() != "1.484131591025766034211155800405523e2" {
		t.Fatal("invalid exp", z)
	}
}

func TestExp2(t *testing.T) {
	x := NewUint64(1)
	x.exponent = -2 // .01
	z := x.Exp()

	if z.String() != "1.01005016708416805754216545690286e0" {
		t.Fatal("invalid exp", z)
	}
}

func TestExp3(t *testing.T) {
	x := NewUint64(100)
	z := x.Exp()

	if z.String() != "2.688117141816135448412625551580014e43" {
		t.Fatal("invalid exp", z)
	}
}

func TestExp4(t *testing.T) {
	x := NewUint64(200)
	z := x.Exp()

	if z.String() != "7.225973768125749258177477042189306e86" {
		t.Fatal("invalid exp", z)
	}
}

func TestExp5(t *testing.T) {
	x := NewUint64(500)
	z := x.Exp()

	if z.String() != "1.403592217852837410739770332840912e217" {
		t.Fatal("invalid exp", z)
	}
}

func TestExpInf(t *testing.T) {
	x := new(Real)
	x.form = FormInf

	z := x.Exp()
	if z.String() != "∞" {
		t.Fatal("invalid exp", z)
	}
}

func TestExpNegInf(t *testing.T) {
	x := new(Real)
	x.form = FormInf
	x.negative = true

	z := x.Exp()
	if z.String() != "0" {
		t.Fatal("invalid exp", z)
	}
}

func TestExpNaN(t *testing.T) {
	x := new(Real)
	x.form = FormNaN

	z := x.Exp()
	if z.String() != "NaN" {
		t.Fatal("invalid exp", z)
	}
}

func TestExp0(t *testing.T) {
	x := new(Real)

	z := x.Exp()
	if z.String() != "1e0" {
		t.Fatal("invalid exp", z)
	}
}

func BenchmarkExp(b *testing.B) {
	x := new(Real)
	x.significand = []byte{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}
	x.validate()
	for b.Loop() {
		x.Exp()
	}
}
