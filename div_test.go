// Copyright 2024 David Fritz. All rights reserved.
// This software may be modified and distributed under the terms of the BSD
// 2-clause license. See the LICENSE file for details.

package realnumber

import "testing"

func TestDiv1(t *testing.T) {
	x := NewInt64(10)
	y := NewInt64(5)

	q := x.Div(y)

	if q.Compare(NewInt64(2)) != 0 {
		t.Fatal("invalid div", q)
	}
}

func TestDiv2(t *testing.T) {
	x := NewInt64(2)
	y := NewInt64(50)

	q := x.Div(y)

	expected := NewInt64(4)
	expected.exponent = -2
	if q.Compare(expected) != 0 {
		t.Fatal("invalid div", q)
	}
}

func TestDiv3(t *testing.T) {
	x := NewInt64(23)
	y := NewInt64(5011513)

	q := x.Div(y)

	if q.String() != "4.589432373017889008768409859457613e-6" {
		t.Fatal("invalid div", q)
	}
}

func BenchmarkDiv(b *testing.B) {
	x := new(Real)
	y := new(Real)
	x.significand = []byte{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}
	y.significand = []byte{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}
	x.validate()
	y.validate()
	for b.Loop() {
		x.Div(y)
	}
}
