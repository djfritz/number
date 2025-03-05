// Copyright 2025 David Fritz. All rights reserved.
// This software may be modified and distributed under the terms of the BSD
// 2-clause license. See the LICENSE file for details.

/*
Package number implements arbitrary-precision decimal floating point numbers and
associated arithmetic.

Currently the only supported type is `Real`, which represents a real (ℝ)
number. Eventually complex (ℂ) and rational (ℚ) numbers will be supported.

Arithmetic operations do not modify their operands and return values are always
deep copies of underlying data. This simplifies programming patterns, but
causes additional memory usage. Additionally, return values of operations will
have the precision of the operand with the largest precision and the rounding
mode of the receiver operand.

A zero value for a Real represents the number 0, and new values can be used in
this way:

```
x := new(Real) // 0
```

Real currently supports three rounding modes:

- Round to nearest even (the default and the default for IEEE-754 floating point numbers)
- Round to nearest
- Round to zero (truncate)

The default precision is 34, which is equivalent to IEEE-754-2008 128-bit
decimal floating point numbers.
*/
package number
