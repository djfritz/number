package real

import "testing"

func TestAdd1(t *testing.T) {
	x := NewInt64(10)
	y := NewInt64(5)

	z := x.Add(y)
	if z.Compare(NewInt64(15)) != 0 {
		t.Fatal("invalid add", z)
	}
}

func TestAdd2(t *testing.T) {
	x := NewInt64(10)
	y := NewInt64(-5)

	z := x.Add(y)
	if z.Compare(NewInt64(5)) != 0 {
		t.Fatal("invalid add", z)
	}
}

func TestAdd3(t *testing.T) {
	x := NewInt64(100)
	y := NewFloat64(1.234)

	z := x.Add(y)
	z.SetPrecision(6)
	expected := NewFloat64(101.234)
	expected.SetPrecision(6)
	if z.Compare(expected) != 0 {
		t.Fatal("invalid add", z, expected)
	}
}

func TestAdd4(t *testing.T) {
	x := NewInt64(-100)
	y := NewInt64(-5)

	z := x.Add(y)
	if z.Compare(NewInt64(-105)) != 0 {
		t.Fatal("invalid add", z)
	}
}
