package real

import (
	"bytes"
	"fmt"
	"testing"
)

func TestReciprocal(t *testing.T) {
	x := new(Real)

	x.SetInt64(1234)
	x.decimal = 2 // 12.34

	z := x.Reciprocal()
	z.precision = 8
	z.round()

	if bytes.Compare(z.digits, []byte{0, 8, 1, 0, 3, 7, 2, 8}) != 0 {
		t.Fatal("invalid reciprocal", z)
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
	z.precision = 8
	z.round()

	if bytes.Compare(z.digits, []byte{0, 0, 1}) != 0 {
		t.Fatal("invalid reciprocal", z)
	}
	if z.negative {
		t.Fatal("invalid negative flag")
	}
	if z.decimal != 3 {
		t.Fatal("invalid decimal point")
	}
}

func TestReciprocal3(t *testing.T) {
	x := new(Real)

	x.digits = []byte{2, 3, 5, 7, 2, 3, 6, 4, 2, 6, 7, 3, 4, 6, 7, 8, 3, 2, 1, 4, 5, 6, 3, 3, 5, 7, 2, 1}

	z := x.Reciprocal()

	fmt.Println(x)
	if bytes.Compare(z.digits, []byte{0, 0, 1}) != 0 {
		t.Fatal("invalid reciprocal", z)
	}
}
