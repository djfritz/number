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
