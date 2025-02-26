// Copyright 2025 David Fritz. All rights reserved.
// This software may be modified and distributed under the terms of the BSD
// 2-clause license. See the LICENSE file for details.

package number

import "testing"

func TestRemainder(t *testing.T) {
	x := NewInt64(102)
	y := NewInt64(3)

	q := x.Remainder(y)

	if q.String() != "0" {
		t.Fatal("invalid remainder", q)
	}
}
