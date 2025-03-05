// Copyright 2025 David Fritz. All rights reserved.
// This software may be modified and distributed under the terms of the BSD
// 2-clause license. See the LICENSE file for details.

package number

import (
	"bytes"
)

// Return the sum of x and y.
func (x *Real) Add(y *Real) *Real {
	x.validate()
	y.validate()

	z := initFrom(x)

	if x.IsInf() && y.IsInf() && x.negative != y.negative {
		z.form = FormNaN
		return z
	} else if x.IsNaN() || y.IsNaN() {
		z.form = FormNaN
		return z
	} else if x.IsInf() {
		z.form = FormInf
		z.negative = x.negative
		return z
	} else if y.IsInf() {
		z.form = FormInf
		z.negative = y.negative
		return z
	}

	// The sum will have the precision of the larger of the two addends
	z.precision = umax(x.precision, y.precision)

	// We only support the common case right now, but we can still do a
	// zero check for performance reasons.
	if x.IsZero() {
		z.CopyValue(y)
		return z
	} else if y.IsZero() {
		z.CopyValue(x)
		return z
	}

	// x + y
	if x.negative == y.negative {
		// x + y == x + y
		// (-x) + (-y) == -(x + y)
		z.negative = x.negative
		z.add(x, y)
	} else {
		// x + (-y) == x - y
		// (-x) + y == y - x
		if !x.negative {
			z.negative = x.negative
			z.sub(x, y)
		} else {
			z.negative = y.negative
			z.sub(y, x)
		}
	}

	z.round()
	return z
}

// Sum the significands of x and y into the significand of r, ignoring the
// sign.
func (z *Real) add(x, y *Real) {
	shiftAmount := x.exponent - y.exponent
	sa := abs(shiftAmount)

	// Special case: if the shift amount is beyond the precision (+1 for
	// rounding) of the precision, there's no work to be done, and we can
	// just return the value with the larger exponent.
	p := umax(x.precision, y.precision)
	if uint(sa) > p+1 {
		if x.exponent > y.exponent {
			z.CopyValue(x)
			return
		} else {
			z.CopyValue(y)
			return
		}
	}

	var addend []byte

	// Shift the larger exponent value to the left, padding zeros to make way for the sum. For example:
	//
	//	xxxxxxxxxx
	//	    yyyyyyyyyy
	//
	// would become
	//	xxxxxxxxxx0000
	//	    yyyyyyyyyy
	//
	// Keep in mind that the number of digits in either value may be less
	// than the precision, so you could end up with something like:
	//
	//	xxxxxxxxxx
	//	   yyy
	//
	// In which case you don't need to pad at all.
	switch {
	case shiftAmount < 0:
		// y's exponent is larger
		z.CopyValue(y)
		addend = x.significand
	case shiftAmount == 0:
		// same exponent, just pad
		z.CopyValue(x)
		addend = y.significand
	case shiftAmount > 0:
		// x's exponent is larger
		z.CopyValue(x)
		addend = y.significand
	}

	// If the addend's significand extends beyond z's, we must pad.
	if sa+len(addend) > len(z.significand) {
		s := make([]byte, sa+len(addend))
		copy(s, z.significand)

		// now we can just copy the addend's tail up to z and chop off
		// the tail from the addend, giving us:
		//
		// 	xxxxxyyy
		//	   yy
		//
		// This saves us the trouble of doing useless adds on zeros.
		//
		// It's also possible that z has trailing zeros, in which case
		// we can opportunistically move up as any overlap there as
		// well:
		//
		//	x0yyyyy

		digits := len(s) - len(z.significand) // Maximum number of digits we can move up, from the right
		// abs(shiftAmount) is the starting position of the addend
		if len(addend) < digits {
			// truncate down to the number of digits in the addend if it's shorter
			digits = len(addend)
		}
		copy(s[len(s)-digits:], addend[len(addend)-digits:])
		addend = addend[:len(addend)-digits]

		z.significand = s
	}

	// Sum from right to left, overflowing if we need. The significands may
	// not be the same length, so we determine the starting point and sum
	// from there.
	//
	// The addends can be in two shapes:
	//
	// 	xxxxxxxxxx
	//	yyyyyyyyyy
	//
	//	xxxxxZZ
	//	  yyy
	i := int(umin(uint(len(z.significand)), uint(len(addend)))) - 1

	for ; i >= 0; i-- {
		z.significand[i+sa] += addend[i]
		if z.significand[i+sa] > 9 {
			z.significand[i+sa] -= 10
			if i+sa == 0 {
				// overflow
				z.exponent++
				s := make([]byte, len(z.significand)+1)
				s[0] = 1
				copy(s[1:], z.significand)
				z.significand = s
			} else {
				z.significand[i+sa-1]++
			}
		}
	}

	// Continue checking for carry from i on z
	for i = i + sa; i >= 0 && z.significand[i] > 9; i-- {
		z.significand[i] -= 10
		if i == 0 {
			// overflow
			z.exponent++
			s := make([]byte, len(z.significand)+1)
			s[0] = 1
			copy(s[1:], z.significand)
			z.significand = s
		} else {
			z.significand[i-1]++
		}
	}
}

// Subtract the significands of x and y into the significand of z, ignoring the
// sign.
func (z *Real) sub(x, y *Real) {
	// check for aliasing
	if (y.exponent > x.exponent) || (x.exponent == y.exponent && bytes.Compare(x.significand, y.significand) == -1) {
		z.negative = !z.negative
		z.sub(y, x)
		return
	}

	shiftAmount := x.exponent - y.exponent
	sa := abs(shiftAmount)

	// Special case: if the shift amount is beyond the precision (+1 for
	// rounding) of the precision, there's no work to be done, and we can
	// just return the value with the larger exponent.
	p := umax(x.precision, y.precision)
	if uint(sa) > p+1 {
		if x.exponent > y.exponent {
			z.CopyValue(x)
			return
		} else {
			z.CopyValue(y)
			return
		}
	}

	var addend []byte

	// Shift the larger exponent value to the left, padding zeros to make way for the sum. For example:
	//
	//	xxxxxxxxxx
	//	    yyyyyyyyyy
	//
	// would become
	//	xxxxxxxxxx0000
	//	    yyyyyyyyyy
	//
	// Keep in mind that the number of digits in either value may be less
	// than the precision, so you could end up with something like:
	//
	//	xxxxxxxxxx
	//	   yyy
	//
	// In which case you don't need to pad at all.
	switch {
	case shiftAmount < 0:
		// y's exponent is larger
		z.CopyValue(y)
		addend = make([]byte, len(x.significand))
		copy(addend, x.significand)
	case shiftAmount == 0:
		// same exponent, just pad
		z.CopyValue(x)
		addend = make([]byte, len(y.significand))
		copy(addend, y.significand)
	case shiftAmount > 0:
		// x's exponent is larger
		z.CopyValue(x)
		addend = make([]byte, len(y.significand))
		copy(addend, y.significand)
	}

	// If the addend's significand extends beyond z's, we must pad.
	if sa+len(addend) > len(z.significand) {
		s := make([]byte, sa+len(addend))
		copy(s, z.significand)
		z.significand = s
	}

	i := int(umin(uint(len(z.significand)), uint(len(addend)))) - 1

	var carry bool
	for ; i >= 0; i-- {
		if z.significand[i+sa] >= addend[i] {
			z.significand[i+sa] -= addend[i]
		} else {
			z.significand[i+sa] += 10 - addend[i]
			if i == 0 {
				carry = true
			} else {
				addend[i-1]++
			}
		}
	}

	if carry {
		for i = i + sa; i >= 0; i-- {
			if z.significand[i] == 0 {
				z.significand[i] = 9
			} else {
				z.significand[i]--
				break
			}
		}
	}
}

// Return the subtraction of y from x.
func (x *Real) Sub(y *Real) *Real {
	x.validate()
	y.validate()
	yn := y.Copy()
	yn.negative = !yn.negative
	return x.Add(yn)
}
