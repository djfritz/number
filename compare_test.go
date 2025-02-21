// Copyright 2025 David Fritz. All rights reserved.
// This software may be modified and distributed under the terms of the BSD
// 2-clause license. See the LICENSE file for details.

package number

import "testing"

func TestCompare1(t *testing.T) {
	x := NewInt64(5)
	y := NewInt64(-5)

	if x.Compare(y) != 1 {
		t.Fatal("invalid compare")
	}
}

func TestCompare2(t *testing.T) {
	x := NewInt64(-5)
	y := NewInt64(100)

	if x.Compare(y) != -1 {
		t.Fatal("invalid compare")
	}
}

func TestCompare3(t *testing.T) {
	x := NewInt64(-5)
	y := NewInt64(-5)

	if x.Compare(y) != 0 {
		t.Fatal("invalid compare")
	}
}

func TestCompare4(t *testing.T) {
	x := NewInt64(-5)
	y := NewFloat64(-5)

	if x.Compare(y) != 0 {
		t.Fatal("invalid compare")
	}
}

func TestCompareInf(t *testing.T) {
	x := new(Real)
	x.form = FormInf
	y := NewInt64(-5)

	if x.Compare(y) != 1 {
		t.Fatal("invalid compare")
	}
}

func TestCompareInfSame(t *testing.T) {
	x := new(Real)
	x.form = FormInf
	y := new(Real)
	y.form = FormInf

	if x.Compare(y) != 0 {
		t.Fatal("invalid compare")
	}
}

func TestCompareInfDifferent(t *testing.T) {
	x := new(Real)
	x.form = FormInf
	y := new(Real)
	y.form = FormInf
	y.negative = true

	if x.Compare(y) != 1 {
		t.Fatal("invalid compare")
	}
}

func TestCompareInfDifferent2(t *testing.T) {
	x := new(Real)
	x.form = FormInf
	x.negative = true
	y := new(Real)
	y.form = FormInf

	if x.Compare(y) != -1 {
		t.Fatal("invalid compare")
	}
}

func TestCompareNaN(t *testing.T) {
	if attempt() {
		t.Fatal("failed compare")
	}
}

func attempt() bool {
	x := new(Real)
	x.form = FormNaN
	y := new(Real)

	defer func() {
		recover()
	}()

	x.Compare(y)
	return true
}
