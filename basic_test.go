package real

import (
	"bytes"
	"testing"
)

func TestAdjust(t *testing.T) {
	x := new(Real)
	y := new(Real)

	x.SetInt64(501)
	x.decimal = 1 // 50.1

	y.SetInt64(100001)
	y.decimal = 5 // 1.00001

	a, b, d := adjust(x, y)

	if bytes.Compare(a, []byte{5, 0, 1, 0, 0, 0, 0}) != 0 {
		t.Fatal("invalid adjust buffer", a)
	}
	if bytes.Compare(b, []byte{0, 1, 0, 0, 0, 0, 1}) != 0 {
		t.Fatal("invalid adjust buffer", b)
	}
	if d != 5 {
		t.Fatal("invalid decimal point")
	}
}

func TestAdd1(t *testing.T) {
	x := new(Real)
	y := new(Real)

	x.SetInt64(501)
	x.decimal = 1 // 50.1

	y.SetInt64(100001)
	y.decimal = 5 // 1.00001

	z := x.Add(y)

	if bytes.Compare(z.digits, []byte{5, 1, 1, 0, 0, 0, 1}) != 0 {
		t.Fatal("invalid add", z)
	}
}

func TestAdd2(t *testing.T) {
	x := new(Real)
	y := new(Real)

	x.SetInt64(-501)
	x.decimal = 1 // 50.1

	y.SetInt64(100001)
	y.decimal = 5 // 1.00001

	z := x.Add(y)

	if bytes.Compare(z.digits, []byte{4, 9, 0, 9, 9, 9, 9}) != 0 {
		t.Fatal("invalid add", z)
	}
	if !z.negative {
		t.Fatal("invalid negative flag")
	}
}

func TestAdd3(t *testing.T) {
	x := new(Real)
	y := new(Real)

	x.SetInt64(501)
	x.decimal = 1 // 50.1

	y.SetInt64(-100001)
	y.decimal = 5 // 1.00001

	z := x.Add(y)

	if bytes.Compare(z.digits, []byte{4, 9, 0, 9, 9, 9, 9}) != 0 {
		t.Fatal("invalid add", z)
	}
	if z.negative {
		t.Fatal("invalid negative flag")
	}
}

func TestAdd4(t *testing.T) {
	x := new(Real)
	y := new(Real)

	x.SetInt64(-50)

	y.SetInt64(100)

	z := x.Add(y)

	if bytes.Compare(z.digits, []byte{5, 0}) != 0 {
		t.Fatal("invalid add", z)
	}
	if z.negative {
		t.Fatal("invalid negative flag")
	}
}

func TestAdd5(t *testing.T) {
	x := new(Real)
	y := new(Real)

	x.SetInt64(50)

	y.SetInt64(-100)

	z := x.Add(y)

	if bytes.Compare(z.digits, []byte{5, 0}) != 0 {
		t.Fatal("invalid add", z)
	}
	if !z.negative {
		t.Fatal("invalid negative flag")
	}
}

func TestSub(t *testing.T) {
	x := new(Real)
	y := new(Real)

	x.SetInt64(50)

	y.SetInt64(100)

	z := x.Sub(y)

	if bytes.Compare(z.digits, []byte{5, 0}) != 0 {
		t.Fatal("invalid add", z)
	}
	if !z.negative {
		t.Fatal("invalid negative flag")
	}
}
