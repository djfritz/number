// Copyright 2025 David Fritz. All rights reserved.
// This software may be modified and distributed under the terms of the BSD
// 2-clause license. See the LICENSE file for details.

package number

import (
	"testing"
)

func TestRoundHalfEven1(t *testing.T) {
	x := NewInt64(12346)
	y := NewInt64(5)
	y.exponent = -1
	z := x.Add(y)
	z.SetPrecision(5)

	if z.String() != "1.2346e4" {
		t.Fatal("invalid round", z)
	}
}
