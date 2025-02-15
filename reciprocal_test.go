// Copyright 2024 David Fritz. All rights reserved.
// This software may be modified and distributed under the terms of the BSD
// 2-clause license. See the LICENSE file for details.

package real

import "testing"

func TestReciprocal1(t *testing.T) {
	x := NewInt64(10)
	z := x.Reciprocal()

	expected := NewInt64(1)
	expected.exponent = -1
	if z.Compare(expected) != 0 {
		t.Fatal("invalid reciprocal", z)
	}
}

func TestReciprocal2(t *testing.T) {
	x := NewInt64(1234)
	x.exponent = 1
	z := x.Reciprocal()

	if z.String() != "8.103727714748784440842787682333874e-2" {
		t.Fatal("invalid reciprocal", z)
	}
}

func TestReciprocal3(t *testing.T) {
	x := NewInt64(-10)
	z := x.Reciprocal()

	expected := NewInt64(-1)
	expected.exponent = -1
	if z.Compare(expected) != 0 {
		t.Fatal("invalid reciprocal", z)
	}
}
