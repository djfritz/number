// Copyright 2024 David Fritz. All rights reserved.
// This software may be modified and distributed under the terms of the BSD
// 2-clause license. See the LICENSE file for details.

package realnumber

import (
	"bytes"
	"testing"
)

func TestSetFloat64(t *testing.T) {
	r := new(Real)
	r.SetFloat64(1.23456789)
	if bytes.Compare(r.significand, []byte{1, 2, 3, 4, 5, 6, 7, 8, 8, 9, 9, 9, 9, 9, 9, 9, 8, 9}) != 0 {
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
	if bytes.Compare(r.significand, []byte{1, 2, 4, 1, 3, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 4}) != 0 {
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
	if bytes.Compare(r.significand, []byte{1, 2, 4, 1, 4, 2, 2, 3, 9, 4, 2, 2, 3, 1, 4, 1, 4, 1}) != 0 {
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
	if bytes.Compare(r.significand, []byte{1, 2, 3, 4, 5, 6, 7, 8, 9}) != 0 {
		t.Fatal("SetUint64 failed", r.significand)
	}
	if r.negative {
		t.Fatal("negative flag set")
	}
	if r.exponent != 9 {
		t.Fatal("invalid exponent")
	}
}

func TestSetInt64(t *testing.T) {
	r := new(Real)

	r.SetInt64(9223372036854775807) // largest int64
	if bytes.Compare(r.significand, []byte{9, 2, 2, 3, 3, 7, 2, 0, 3, 6, 8, 5, 4, 7, 7, 5, 8, 0, 7}) != 0 {
		t.Fatal("SetInt64 failed")
	}
	if r.negative {
		t.Fatal("negative flag set")
	}
	if r.exponent != 18 {
		t.Fatal("invalid exponent")
	}

	r.SetInt64(-9223372036854775808) // smallest int64
	if bytes.Compare(r.significand, []byte{9, 2, 2, 3, 3, 7, 2, 0, 3, 6, 8, 5, 4, 7, 7, 5, 8, 0, 8}) != 0 {
		t.Fatal("SetInt64 failed", r.significand)
	}
	if !r.negative {
		t.Fatal("negative flag not set")
	}
	if r.exponent != 18 {
		t.Fatal("invalid exponent")
	}

	r.SetInt64(0)
	if bytes.Compare(r.significand, []byte{}) != 0 {
		t.Fatal("SetInt64 failed", r.significand)
	}
	if r.negative {
		t.Fatal("negative flag set")
	}
	if r.exponent != 0 {
		t.Fatal("invalid exponent", r.exponent)
	}

	r.SetInt64(-1337)
	if bytes.Compare(r.significand, []byte{1, 3, 3, 7}) != 0 {
		t.Fatal("SetInt64 failed", r.significand)
	}
	if !r.negative {
		t.Fatal("negative flag not set")
	}
	if r.exponent != 3 {
		t.Fatal("invalid exponent")
	}
}

func TestString(t *testing.T) {
	r := new(Real)
	r.SetInt64(-9223372036854775808) // smallest int64
	if r.String() != "-9.223372036854775808e18" {
		t.Fatal("invalid string", r.String())
	}
	r.SetInt64(501)
	r.exponent = 1
	if r.String() != "5.01e1" {
		t.Fatal("invalid string", r.String())
	}
}

func TestTrim(t *testing.T) {
	r := new(Real)
	r.significand = []byte{0, 0, 0, 0, 0, 0, 0, 1, 2, 3, 4, 0, 0, 0, 0, 0}

	r.trim()
	if bytes.Compare(r.significand, []byte{1, 2, 3, 4}) != 0 {
		t.Fatal("invalid trim", r.significand)
	}
}

func TestTrim2(t *testing.T) {
	r := new(Real)
	r.significand = []byte{1, 2, 3, 4, 0, 0, 0, 0, 0}

	r.trim()
	if bytes.Compare(r.significand, []byte{1, 2, 3, 4}) != 0 {
		t.Fatal("invalid trim", r.significand)
	}
}

func TestTrim3(t *testing.T) {
	r := new(Real)
	r.significand = []byte{1, 2, 3, 4}

	r.trim()
	if bytes.Compare(r.significand, []byte{1, 2, 3, 4}) != 0 {
		t.Fatal("invalid trim", r.significand)
	}
}

func TestTrim4(t *testing.T) {
	r := new(Real)
	r.significand = []byte{1}

	r.trim()
	if bytes.Compare(r.significand, []byte{1}) != 0 {
		t.Fatal("invalid trim", r.significand)
	}
}

func TestTrim5(t *testing.T) {
	r := new(Real)
	r.significand = []byte{0, 1, 0}

	r.trim()
	if bytes.Compare(r.significand, []byte{1}) != 0 {
		t.Fatal("invalid trim", r.significand)
	}
}

func TestTrim6(t *testing.T) {
	r := new(Real)
	r.significand = []byte{}

	r.trim()
	if bytes.Compare(r.significand, []byte{}) != 0 {
		t.Fatal("invalid trim", r.significand)
	}
}

func TestTrim7(t *testing.T) {
	r := new(Real)
	r.significand = []byte{0, 0, 0, 0, 0, 0, 0, 1, 2, 3, 4, 0, 0, 0, 0, 0}
	r.exponent = 7

	r.trim()
	if bytes.Compare(r.significand, []byte{1, 2, 3, 4}) != 0 {
		t.Fatal("invalid trim", r.significand)
	}
	if r.exponent != 0 {
		t.Fatal("invalid exponent", r.exponent)
	}
}

func TestRound1(t *testing.T) {
	r := NewUint64(12345678900000)
	if r.String() != "1.23456789e13" {
		t.Fatal("invalid NewUint64")
	}
	r.SetPrecision(5)

	if r.String() != "1.2346e13" {
		t.Fatal("invalid round", r.String())
	}
}

func TestRound2(t *testing.T) {
	r := NewUint64(12345678900000)

	if r.String() != "1.23456789e13" {
		t.Fatal("invalid round", r.String())
	}
}

func TestRound3(t *testing.T) {
	r := NewUint64(12345478900000)
	r.SetPrecision(5)

	if r.String() != "1.2346e13" {
		t.Fatal("invalid round", r.String())
	}
}

func TestRound4(t *testing.T) {
	r := NewUint64(12345378900000)
	r.SetPrecision(5)

	if r.String() != "1.2345e13" {
		t.Fatal("invalid round", r.String())
	}
}

func TestRound5(t *testing.T) {
	r := NewUint64(1233999999)
	r.exponent = 0
	r.SetPrecision(7)

	if r.String() != "1.234e0" {
		t.Fatal("invalid round", r.String())
	}
}

func TestRound6(t *testing.T) {
	r := NewUint64(1233999999)
	r.exponent = 0
	r.SetPrecision(8)

	if r.String() != "1.234e0" {
		t.Fatal("invalid round", r.String())
	}
}

func TestRound7(t *testing.T) {
	r := NewUint64(1233999999)
	r.exponent = 0
	r.SetPrecision(4)

	if r.String() != "1.234e0" {
		t.Fatal("invalid round", r.String())
	}
}

func TestInteger1(t *testing.T) {
	x := NewFloat64(1.234)
	z := x.Integer()
	if bytes.Compare(z.significand, []byte{1}) != 0 {
		t.Fatal("invalid significand", z.significand)
	}
}

func TestInteger2(t *testing.T) {
	x := NewFloat64(12.34)
	z := x.Integer()
	if bytes.Compare(z.significand, []byte{1, 2}) != 0 {
		t.Fatal("invalid significand", z.significand)
	}
}

func TestInteger3(t *testing.T) {
	x := NewFloat64(123.4)
	z := x.Integer()
	if bytes.Compare(z.significand, []byte{1, 2, 3}) != 0 {
		t.Fatal("invalid significand", z.significand)
	}
}

func TestInteger4(t *testing.T) {
	x := NewFloat64(1234.0)
	z := x.Integer()
	if bytes.Compare(z.significand, []byte{1, 2, 3, 4}) != 0 {
		t.Fatal("invalid significand", z.significand)
	}
}
