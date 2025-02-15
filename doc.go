// Copyright 2024 David Fritz. All rights reserved.
// This software may be modified and distributed under the terms of the BSD
// 2-clause license. See the LICENSE file for details.

/*
Package realnumber implements arbitrary precision decimal floating point numbers and
associated arithmetic. Unlike binary floating point numbers, package real
stores decimal digits of the significand as decimal values (stored as a
[]byte). This means that decimal representations can be stored exactly (unlike
many numbers in binary floating point).

A zero value for a Real represents the number 0, and new values can be used in
this way:

	x := new(Real) // 0

Unless specified, real values use the default rounding mode and precision.

Arithmetic operations do not modify their operands.
*/
package realnumber
