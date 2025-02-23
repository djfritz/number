[![Go Reference](https://pkg.go.dev/badge/github.com/djfritz/number.svg)](https://pkg.go.dev/github.com/djfritz/number)

Package number implements arbitrary precision decimal floating point numbers and
associated arithmetic. Unlike binary floating point numbers, package number 
stores decimal digits of the significand as decimal values (stored as a
[]byte). This means that decimal representations can be stored exactly (unlike
many numbers in binary floating point).

Currently the only type in this package is `Real`, which is meant to represent
a real (ℝ) number. Eventually complex (ℂ) and rational (ℚ) numbers will be
supported.

## Real numbers 

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

Unless specified, real values use the default rounding mode and precision.

Arithmetic operations do not modify their operands.

### Tests

Beyond the unit tests in this package, Real is tested against Mike Cowlishaw's
excellent [dectest](https://speleotrove.com/decimal/) tests. Those tests are
kept in [another package](https://github.com/djfritz/numbertests), mostly to
avoid embedding an ICU license in this package.

Currently, 8185 of the subset dectest tests are run against Real. Six of those
currently fail, but only because of inexact rounding expected in the test
suite. Real computes to better than .5ulp in those cases and provides a more
accurate answer.

