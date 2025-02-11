package real

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
	z.SetPrecision(8)

	if z.String() != "3.8446751e-2" {
		t.Fatal("invalid power", z)
	}
}

func TestPow1(t *testing.T) {
	x := NewInt64(5)
	y := NewInt64(8)
	z := x.Pow(y)
	z.SetPrecision(10)

	if z.String() != "3.90625e5" {
		t.Fatal("invalid power", z)
	}
}

func TestPow2(t *testing.T) {
	x := NewInt64(9)
	y := NewFloat64(.5)
	z := x.Pow(y)
	z.SetPrecision(10)

	if z.String() != "3e0" {
		t.Fatal("invalid power", z)
	}
}

func TestSqrt1(t *testing.T) {
	x := NewInt64(9)
	z := x.Sqrt()
	z.SetPrecision(10)

	if z.String() != "3e0" {
		t.Fatal("invalid sqrt", z)
	}
}

func TestSqrt2(t *testing.T) {
	x := NewInt64(2)
	z := x.Sqrt()
	z.SetPrecision(10)

	if z.String() != "1.414213562e0" {
		t.Fatal("invalid sqrt", z)
	}
}

func TestSqrt3(t *testing.T) {
	x := NewInt64(2000)
	z := x.Sqrt()
	z.SetPrecision(10)

	if z.String() != "4.472135955e1" {
		t.Fatal("invalid sqrt", z)
	}
}

func TestSqrt4(t *testing.T) {
	x := NewFloat64(.002)
	z := x.Sqrt()
	z.SetPrecision(10)

	if z.String() != "4.472135955e-2" {
		t.Fatal("invalid sqrt", z)
	}
}
