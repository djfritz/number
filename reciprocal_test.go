package real

import "testing"

func TestReciprocal1(t *testing.T) {
	x := NewInt64(10)
	z := x.Reciprocal()

	expected := NewInt64(1)
	expected.exponent = -1
	if z.Compare(expected) != 0 {
		t.Fatal("invalid reciprocal", z)
	}
}

func TestReciprocal2(t *testing.T) {
	x := NewFloat64(12.34)
	z := x.Reciprocal()
	z.SetPrecision(8)

	expected := NewInt64(81037277)
	expected.exponent = -2
	if z.Compare(expected) != 0 {
		t.Fatal("invalid reciprocal", z)
	}
}

func TestReciprocal3(t *testing.T) {
	x := NewInt64(-10)
	z := x.Reciprocal()

	expected := NewInt64(-1)
	expected.exponent = -1
	if z.Compare(expected) != 0 {
		t.Fatal("invalid reciprocal", z)
	}
}
