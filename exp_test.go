package real

import "testing"

func TestExp1(t *testing.T) {
	x := NewUint64(5)
	z := x.Exp()
	z.SetPrecision(10)

	if z.String() != "1.484131591e2" {
		t.Fatal("invalid exp", z)
	}
}

func TestExp2(t *testing.T) {
	x := NewUint64(1)
	x.exponent = -2 // .01
	z := x.Exp()
	z.SetPrecision(10)

	if z.String() != "1.010050167e0" {
		t.Fatal("invalid exp", z)
	}
}

func TestExp3(t *testing.T) {
	x := NewUint64(100)
	z := x.Exp()
	z.SetPrecision(10)

	if z.String() != "2.688117142e43" {
		t.Fatal("invalid exp", z)
	}
}

func TestExp4(t *testing.T) {
	x := NewUint64(200)
	z := x.Exp()
	z.SetPrecision(10)

	if z.String() != "7.225973768e86" {
		t.Fatal("invalid exp", z)
	}
}

func TestExp5(t *testing.T) {
	x := NewUint64(500)
	z := x.Exp()
	z.SetPrecision(10)

	if z.String() != "1.403592218e217" {
		t.Fatal("invalid exp", z)
	}
}
