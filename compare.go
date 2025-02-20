// Copyright 2024 David Fritz. All rights reserved.
// This software may be modified and distributed under the terms of the BSD
// 2-clause license. See the LICENSE file for details.

package number

import "bytes"

// Compare x with y, returing an integer representing:
//
//	1  : x > y
//	0  : x == y
//	-1 : x < y
func (x *Real) Compare(y *Real) int {
	if x.IsInf() && y.IsInf() && x.negative == y.negative {
		return 0
	} else if x.IsNaN() || y.IsNaN() {
		panic("cannot compare NaN")
	} else if x.IsInf() {
		if x.negative {
			return -1
		}
		return 1
	} else if y.IsInf() {
		if y.negative {
			return 1
		}
		return -1
	}

	if !x.negative && y.negative {
		return 1
	} else if x.negative && !y.negative {
		return -1
	}

	if !x.negative && y.IsZero() {
		return 1
	} else if x.negative && y.IsZero() {
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
