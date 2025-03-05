// Copyright 2025 David Fritz. All rights reserved.
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
	// non-real forms
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

	// mismatched negative flags
	if !x.negative && y.negative {
		return 1
	} else if x.negative && !y.negative {
		return -1
	}

	// zeros
	if x.IsZero() && y.IsZero() {
		return 0
	} else if x.IsZero() {
		if y.negative {
			return 1
		} else {
			return -1
		}
	} else if y.IsZero() {
		if x.negative {
			return -1
		} else {
			return 1
		}
	}

	// exponents
	if x.exponent > y.exponent {
		if x.negative {
			return -1
		} else {
			return 1
		}
	} else if x.exponent < y.exponent {
		if x.negative {
			return 1
		} else {
			return -1
		}
	}

	// same exponents, just compare the significand
	if x.negative {
		return bytes.Compare(y.significand, x.significand)
	}
	return bytes.Compare(x.significand, y.significand)
}

// Return a copy of the larger of x and y, or x if the values are equal.
func (x *Real) Max(y *Real) *Real {
	switch x.Compare(y) {
	case 1:
		return x.Copy()
	case -1:
		return y.Copy()
	default:
		return x.Copy()
	}
}

// Return a copy of the smaller of x and y, or x if the values are equal.
func (x *Real) Min(y *Real) *Real {
	switch x.Compare(y) {
	case 1:
		return y.Copy()
	case -1:
		return x.Copy()
	default:
		return x.Copy()
	}
}
