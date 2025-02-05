package real

import (
	"bytes"
	"testing"
)

func TestSetUint64(t *testing.T) {
	r := new(Real)

	r.SetUint64(1234567890)
	if bytes.Compare(r.digits, []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}) != 0 {
		t.Fatal("SetUint64 failed")
	}
	if r.negative {
		t.Fatal("negative flag set")
	}
	if r.decimal != 0 {
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
	if r.decimal != 0 {
		t.Fatal("non-zero decimal place")
	}

	r.SetInt64(-9223372036854775808) // smallest int64
	if bytes.Compare(r.digits, []byte{9, 2, 2, 3, 3, 7, 2, 0, 3, 6, 8, 5, 4, 7, 7, 5, 8, 0, 8}) != 0 {
		t.Fatal("SetUint64 failed", r.digits)
	}
	if !r.negative {
		t.Fatal("negative flag not set")
	}
	if r.decimal != 0 {
		t.Fatal("non-zero decimal place")
	}

	r.SetInt64(0)
	if bytes.Compare(r.digits, []byte{}) != 0 {
		t.Fatal("SetUint64 failed", r.digits)
	}
	if r.negative {
		t.Fatal("negative flag set")
	}
	if r.decimal != 0 {
		t.Fatal("non-zero decimal place")
	}

	r.SetInt64(-1337)
	if bytes.Compare(r.digits, []byte{1, 3, 3, 7}) != 0 {
		t.Fatal("SetUint64 failed", r.digits)
	}
	if !r.negative {
		t.Fatal("negative flag not set")
	}
	if r.decimal != 0 {
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
	r.decimal = 1
	if r.String() != "50.1" {
		t.Fatal("invalid string", r.String())
	}
}
