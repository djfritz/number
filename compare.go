// Copyright 2024 David Fritz. All rights reserved.
// This software may be modified and distributed under the terms of the BSD
// 2-clause license. See the LICENSE file for details.

package real

import "bytes"

// Compare x with y, returing an integer representing:
//
//	1  : x > y
//	0  : x == y
//	-1 : x < y
func (x *Real) Compare(y *Real) int {
	if !x.negative && y.negative {
		return 1
	} else if x.negative && !y.negative {
		return -1
	}

	if x.exponent > y.exponent {
		return 1
	} else if x.exponent < y.exponent {
		return -1
	}

	// same exponents, just compare the significand
	return bytes.Compare(x.significand, y.significand)
}
