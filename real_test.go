package real

import (
	"bytes"
	"testing"
)

func TestSetFloat64(t *testing.T) {
	r := new(Real)
	r.SetFloat64(1.23456789)
	if bytes.Compare(r.digits, []byte{1, 2, 3, 4, 5, 6, 7, 8, 8, 9, 9, 9, 9, 9, 9, 9, 8, 9}) != 0 {
		t.Fatal("SetFloat64 failed", r)
	}
	if r.negative {
		t.Fatal("negative flag set")
	}
	if r.exponent != 0 {
		t.Fatal("invalid exponent", r.exponent)
	}
}

func TestSetFloat642(t *testing.T) {
	r := new(Real)
	r.SetFloat64(.0000000000012414)
	if bytes.Compare(r.digits, []byte{1, 2, 4, 1, 3, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 4}) != 0 {
		t.Fatal("SetFloat64 failed", r)
	}
	if r.negative {
		t.Fatal("negative flag set")
	}
	if r.exponent != -12 {
		t.Fatal("invalid exponent")
	}
}

func TestSetFloat643(t *testing.T) {
	r := new(Real)
	r.SetFloat64(12414223942231414151231231)
	if bytes.Compare(r.digits, []byte{1, 2, 4, 1, 4, 2, 2, 3, 9, 4, 2, 2, 3, 1, 4, 1, 4, 1}) != 0 {
		t.Fatal("SetFloat64 failed", r)
	}
	if r.negative {
		t.Fatal("negative flag set")
	}
	if r.exponent != 25 {
		t.Fatal("invalid exponent", r.exponent)
	}
}

func TestSetUint64(t *testing.T) {
	r := new(Real)

	r.SetUint64(1234567890)
	if bytes.Compare(r.digits, []byte{1, 2, 3, 4, 5, 6, 7, 8, 9}) != 0 {
		t.Fatal("SetUint64 failed", r.digits)
	}
	if r.negative {
		t.Fatal("negative flag set")
	}
	if r.exponent != 9 {
		t.Fatal("non-zero decimal place")
	}
}

func TestSetInt64(t *testing.T) {
	r := new(Real)

	r.SetInt64(9223372036854775807) // largest int64
	if bytes.Compare(r.digits, []byte{9, 2, 2, 3, 3, 7, 2, 0, 3, 6, 8, 5, 4, 7, 7, 5, 8, 0, 7}) != 0 {
		t.Fatal("SetUint64 failed")
	}
	if r.negative {
		t.Fatal("negative flag set")
	}
	if r.exponent != 18 {
		t.Fatal("non-zero decimal place")
	}

	r.SetInt64(-9223372036854775808) // smallest int64
	if bytes.Compare(r.digits, []byte{9, 2, 2, 3, 3, 7, 2, 0, 3, 6, 8, 5, 4, 7, 7, 5, 8, 0, 8}) != 0 {
		t.Fatal("SetUint64 failed", r.digits)
	}
	if !r.negative {
		t.Fatal("negative flag not set")
	}
	if r.exponent != 18 {
		t.Fatal("non-zero decimal place")
	}

	r.SetInt64(0)
	if bytes.Compare(r.digits, []byte{}) != 0 {
		t.Fatal("SetUint64 failed", r.digits)
	}
	if r.negative {
		t.Fatal("negative flag set")
	}
	if r.exponent != 0 {
		t.Fatal("non-zero decimal place", r.exponent)
	}

	r.SetInt64(-1337)
	if bytes.Compare(r.digits, []byte{1, 3, 3, 7}) != 0 {
		t.Fatal("SetUint64 failed", r.digits)
	}
	if !r.negative {
		t.Fatal("negative flag not set")
	}
	if r.exponent != 3 {
		t.Fatal("non-zero decimal place")
	}
}

func TestString(t *testing.T) {
	r := new(Real)
	r.SetInt64(-9223372036854775808) // smallest int64
	if r.String() != "-9223372036854775808" {
		t.Fatal("invalid string", r.String())
	}
	r.SetInt64(501)
	r.exponent = 1
	if r.String() != "50.1" {
		t.Fatal("invalid string", r.String())
	}
}
