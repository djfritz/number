// Copyright 2024 David Fritz. All rights reserved.
// This software may be modified and distributed under the terms of the BSD
// 2-clause license. See the LICENSE file for details.

package number

import "testing"

func TestIpow(t *testing.T) {
	x := NewInt64(5)
	z := x.ipow(8)

	if z.Compare(NewInt64(390625)) != 0 {
		t.Fatal("invalid power", z)
	}
}

func TestIpow2(t *testing.T) {
	x := NewInt64(51)
	x.exponent = 0
	z := x.ipow(8)

	if z.String() != "4.5767944570401e5" {
		t.Fatal("invalid power", z)
	}
}

func TestIpow3(t *testing.T) {
	x := NewInt64(51)
	x.exponent = 0
	z := x.ipow(-2)

	if z.String() != "3.844675124951941560938100730488276e-2" {
		t.Fatal("invalid power", z)
	}
}

func TestIpowInf(t *testing.T) {
	x := new(Real)
	x.form = FormInf
	z := x.ipow(8)

	if z.String() != "∞" {
		t.Fatal("invalid power", z)
	}
}

func TestIpowNegInf(t *testing.T) {
	x := new(Real)
	x.form = FormInf
	x.negative = true
	z := x.ipow(8)

	if z.String() != "∞" {
		t.Fatal("invalid power", z)
	}
}

func TestIpowZero(t *testing.T) {
	x := new(Real)
	z := x.ipow(8)

	if z.String() != "0" {
		t.Fatal("invalid power", z)
	}
}

func TestPow1(t *testing.T) {
	x := NewInt64(5)
	y := NewInt64(8)
	z := x.Pow(y)

	if z.String() != "3.90625e5" {
		t.Fatal("invalid power", z)
	}
}

func TestPow2(t *testing.T) {
	x := NewInt64(9)
	y := NewFloat64(.5)
	z := x.Pow(y)

	if z.String() != "3e0" {
		t.Fatal("invalid power", z)
	}
}

func TestPowInfInf(t *testing.T) {
	x := new(Real)
	x.form = FormInf
	y := new(Real)
	y.form = FormInf
	z := x.Pow(y)

	if z.String() != "NaN" {
		t.Fatal("invalid power", z)
	}
}

func TestPowInfZero(t *testing.T) {
	x := new(Real)
	x.form = FormInf
	y := new(Real)
	z := x.Pow(y)

	if z.String() != "1e0" {
		t.Fatal("invalid power", z)
	}
}

func TestPowInf(t *testing.T) {
	x := new(Real)
	x.form = FormInf
	y := NewUint64(1)
	z := x.Pow(y)

	if z.String() != "∞" {
		t.Fatal("invalid power", z)
	}
}

func TestPowInfAExactly1(t *testing.T) {
	x := NewUint64(1)
	y := new(Real)
	y.form = FormInf
	z := x.Pow(y)

	if z.String() != "NaN" {
		t.Fatal("invalid power", z)
	}
}

func TestPowInfA1(t *testing.T) {
	x := NewUint64(5)
	y := new(Real)
	y.form = FormInf
	z := x.Pow(y)

	if z.String() != "∞" {
		t.Fatal("invalid power", z)
	}
}

func TestPowNegInfA1(t *testing.T) {
	x := NewUint64(5)
	y := new(Real)
	y.form = FormInf
	y.negative = true
	z := x.Pow(y)

	if z.String() != "0" {
		t.Fatal("invalid power", z)
	}
}

func TestPowInfA0(t *testing.T) {
	x := NewFloat64(.5)
	y := new(Real)
	y.form = FormInf
	z := x.Pow(y)

	if z.String() != "0" {
		t.Fatal("invalid power", z)
	}
}

func TestPowInfANeg(t *testing.T) {
	x := NewFloat64(-.5)
	y := new(Real)
	y.form = FormInf
	z := x.Pow(y)

	if z.String() != "NaN" {
		t.Fatal("invalid power", z)
	}
}

func TestPowZero(t *testing.T) {
	x := new(Real)
	y := NewUint64(1)
	z := x.Pow(y)

	if z.String() != "0" {
		t.Fatal("invalid power", z)
	}
}

func TestSqrt1(t *testing.T) {
	x := NewInt64(9)
	z := x.Sqrt()

	if z.String() != "3e0" {
		t.Fatal("invalid sqrt", z)
	}
}

func TestSqrt2(t *testing.T) {
	x := NewInt64(2)
	z := x.Sqrt()

	if z.String() != "1.414213562373095048801688724209698e0" {
		t.Fatal("invalid sqrt", z)
	}
}

func TestSqrt3(t *testing.T) {
	x := NewInt64(2000)
	z := x.Sqrt()

	if z.String() != "4.472135954999579392818347337462552e1" {
		t.Fatal("invalid sqrt", z)
	}
}

func TestSqrt4(t *testing.T) {
	x := NewInt64(2)
	x.exponent = -3
	z := x.Sqrt()

	if z.String() != "4.472135954999579392818347337462552e-2" {
		t.Fatal("invalid sqrt", z)
	}
}

func TestSqrtNeg1(t *testing.T) {
	x := NewInt64(-1)
	z := x.Sqrt()

	if z.String() != "NaN" {
		t.Fatal("invalid sqrt", z)
	}
}

func TestSqrtInf(t *testing.T) {
	x := new(Real)
	x.form = FormInf
	z := x.Sqrt()

	if z.String() != "∞" {
		t.Fatal("invalid sqrt", z)
	}
}

func TestSqrtNegInf(t *testing.T) {
	x := new(Real)
	x.form = FormInf
	x.negative = true
	z := x.Sqrt()

	if z.String() != "NaN" {
		t.Fatal("invalid sqrt", z)
	}
}

func TestSqrtNaN(t *testing.T) {
	x := new(Real)
	x.form = FormNaN
	z := x.Sqrt()

	if z.String() != "NaN" {
		t.Fatal("invalid sqrt", z)
	}
}
