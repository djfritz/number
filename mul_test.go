package real

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
