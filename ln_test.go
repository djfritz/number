package real

import "testing"

func TestLn1(t *testing.T) {
	x := NewUint64(5)
	z := x.Ln()
	z.SetPrecision(10)

	if z.String() != "1.609437912e0" {
		t.Fatal("invalid ln", z)
	}
}

func TestLn2(t *testing.T) {
	x := NewFloat64(2)
	z := x.Ln()
	z.SetPrecision(10)

	if z.String() != "6.931471806e-1" {
		t.Fatal("invalid ln", z)
	}
}

func TestLn3(t *testing.T) {
	x := NewFloat64(.002)
	z := x.Ln()
	z.SetPrecision(10)

	if z.String() != "-6.214608098e0" {
		t.Fatal("invalid ln", z)
	}
}

func TestLn4(t *testing.T) {
	x := NewUint64(1)
	z := x.Ln()
	z.SetPrecision(10)

	if z.String() != "0" {
		t.Fatal("invalid ln", z)
	}
}
