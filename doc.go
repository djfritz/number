// Copyright 2025 David Fritz. All rights reserved.
// This software may be modified and distributed under the terms of the BSD
// 2-clause license. See the LICENSE file for details.

/*
Package number implements arbitrary precision decimal floating point numbers and
associated arithmetic. Unlike binary floating point numbers, package number
stores decimal digits of the significand as decimal values (stored as a
[]byte). This means that decimal representations can be stored exactly (unlike
many numbers in binary floating point).

Currently only real numbers ℝ are implemented.

A zero value for a Real represents the number 0, and new values can be used in
this way:

	x := new(Real) // 0

Unless specified, real values use the default rounding mode and precision.

Arithmetic operations do not modify their operands.
*/
package number
