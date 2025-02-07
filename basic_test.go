package real

import (
	"bytes"
	"testing"
)

func TestTrim(t *testing.T) {
	x := new(Real)

	x.digits = []byte{0, 5, 0, 1}
	x.decimal = 4 // .0501

	x.trim()

	if x.String() != ".0501" {
		t.Fatal("invalid trim", x.String())
	}
}

func TestTrim2(t *testing.T) {
	x := new(Real)

	x.digits = []byte{0, 5, 0, 1, 0, 0, 0}
	x.decimal = 7 // .0501

	x.trim()

	if x.String() != ".0501" {
		t.Fatal("invalid trim", x.String())
	}
}

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

func TestMul1(t *testing.T) {
	x := new(Real)
	y := new(Real)

	x.SetInt64(1234)

	y.SetInt64(567)

	z := x.Mul(y)

	if bytes.Compare(z.digits, []byte{6, 9, 9, 6, 7, 8}) != 0 {
		t.Fatal("invalid mul", z)
	}
	if z.negative {
		t.Fatal("invalid negative flag")
	}
}

func TestMul2(t *testing.T) {
	x := new(Real)
	y := new(Real)

	x.SetInt64(1234)
	x.decimal = 2 // 12.34

	y.SetInt64(5671)
	y.decimal = 1 // 567.1

	z := x.Mul(y)

	if bytes.Compare(z.digits, []byte{6, 9, 9, 8, 0, 1, 4}) != 0 {
		t.Fatal("invalid mul", z)
	}
	if z.negative {
		t.Fatal("invalid negative flag")
	}
	if z.decimal != 3 {
		t.Fatal("invalid decimal point", z)
	}
}

func TestMul3(t *testing.T) {
	x := new(Real)
	y := new(Real)

	x.SetInt64(1234)
	x.decimal = 2 // 12.34

	y.SetInt64(5)
	y.decimal = 1 // .5

	z := x.Mul(y)

	if bytes.Compare(z.digits, []byte{6, 1, 7}) != 0 {
		t.Fatal("invalid mul", z)
	}
	if z.negative {
		t.Fatal("invalid negative flag")
	}
	if z.decimal != 2 {
		t.Fatal("invalid decimal point", z.decimal)
	}
	if z.String() != "6.17" {
		t.Fatal("invalid string", z.String())
	}
}

func TestAddCatchOverflow(t *testing.T) {
	x := new(Real)
	y := new(Real)

	x.SetInt64(1928140)
	y.SetInt64(11342000)

	z := x.Add(y)

	if bytes.Compare(z.digits, []byte{1, 3, 2, 7, 0, 1, 4, 0}) != 0 {
		t.Fatal("invalid add", z)
	}
	if z.negative {
		t.Fatal("invalid negative flag")
	}
	if z.decimal != 0 {
		t.Fatal("invalid decimal point")
	}
}

func TestReciprocal(t *testing.T) {
	x := new(Real)

	x.SetInt64(1234)
	x.decimal = 2 // 12.34

	z := x.Reciprocal()
	z.precision = 8
	z.fix()

	if bytes.Compare(z.digits, []byte{0, 8, 1, 0, 3, 7, 2, 8}) != 0 {
		t.Fatal("invalid mul", z)
	}
	if z.negative {
		t.Fatal("invalid negative flag")
	}
	if z.decimal != 8 {
		t.Fatal("invalid decimal point")
	}
}

func TestReciprocal2(t *testing.T) {
	x := new(Real)

	x.SetInt64(1000)

	z := x.Reciprocal()

	if bytes.Compare(z.digits, []byte{0, 0, 1}) != 0 {
		t.Fatal("invalid mul", z)
	}
	if z.negative {
		t.Fatal("invalid negative flag")
	}
	if z.decimal != 3 {
		t.Fatal("invalid decimal point")
	}
}

func TestDiv(t *testing.T) {
	x := new(Real)
	y := new(Real)

	x.SetInt64(5)
	y.SetInt64(4)

	z := x.Div(y)

	if bytes.Compare(z.digits, []byte{1, 2, 5}) != 0 {
		t.Fatal("invalid div", z)
	}
	if z.negative {
		t.Fatal("invalid negative flag")
	}
	if z.decimal != 2 {
		t.Fatal("invalid decimal point")
	}
}

func TestDiv2(t *testing.T) {
	x := new(Real)
	y := new(Real)

	x.SetInt64(302875106592253)
	y.SetInt64(87178291200)

	z := x.Div(y)

	if bytes.Compare(z.digits, []byte{1, 2, 5}) != 0 {
		t.Fatal("invalid add", z)
	}
	if z.negative {
		t.Fatal("invalid negative flag")
	}
	if z.decimal != 2 {
		t.Fatal("invalid decimal point")
	}
}
