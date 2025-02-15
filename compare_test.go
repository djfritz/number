// Copyright 2024 David Fritz. All rights reserved.
// This software may be modified and distributed under the terms of the BSD
// 2-clause license. See the LICENSE file for details.

package real

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
