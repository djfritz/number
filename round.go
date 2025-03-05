// Copyright 2025 David Fritz. All rights reserved.
// This software may be modified and distributed under the terms of the BSD
// 2-clause license. See the LICENSE file for details.

package number

import (
	"errors"
)

// Rounding modes.
const (
	ModeNearestEven = iota
	ModeNearest
	ModeZero
)

var ErrInvalidMode = errors.New("invalid mode")

// Set the rounding mode.
func (x *Real) SetMode(m int) error {
	switch m {
	case ModeNearestEven:
		fallthrough
	case ModeNearest:
		fallthrough
	case ModeZero:
		x.mode = m
	default:
		return ErrInvalidMode
	}
	return nil
}

// Return the rounding mode.
func (x *Real) Mode() int {
	return x.mode
}

// Round the value to the set precision and rounding mode, if necessary.
func (x *Real) round() {
	x.validate()
	x.roundTo(x.precision)
	if x.IsZero() {
		x.negative = false
	}
}

// Round the value to the given precision and rounding mode.
func (x *Real) roundTo(p uint) {
	defer x.trim()

	if uint(len(x.significand)) <= p {
		// number is exact, no rounding needed.
		return
	}

	switch x.mode {
	case ModeNearestEven:
		x.roundToNearestEven(p)
	case ModeNearest:
		x.roundToNearest(p)
	case ModeZero:
		// just truncate
	}

	x.significand = x.significand[:p]
}

func (x *Real) roundToNearestEven(p uint) {
	d := x.significand[p]

	switch {
	case d < 5:
		// round down
	case d > 5:
		// round up
		if p == 0 {
			x.significand[0] = 1
			x.exponent++
			return
		} else {
			x.significand[p-1]++
		}
	case d == 5:
		// round up if any of remaining digits are non-zero
		var nonzero bool
		for i := int(p) + 1; i < len(x.significand); i++ {
			if x.significand[i] != 0 {
				nonzero = true
				break
			}
		}
		if nonzero {
			if p == 0 {
				x.significand[0] = 1
				x.exponent++
				return
			} else {
				x.significand[p-1]++
			}
		} else {
			// round to nearest even!
			if p != 0 {
				if x.significand[p-1]%2 != 0 {
					x.significand[p-1]++
				}
			}
		}
	}

	// now unwind to the left to make sure we don't have any lingering carry
	for i := int(p) - 1; i >= 0; i-- {
		if x.significand[i] < 10 {
			break
		}
		x.significand[i] -= 10

		if i == 0 {
			// pad
			x.significand = append([]byte{1}, x.significand...)
			x.exponent++
			break
		}
		x.significand[i-1]++
	}
}

func (x *Real) roundToNearest(p uint) {
	d := x.significand[p]

	switch {
	case d < 5:
		// round down
	case d >= 5:
		// round up
		if p == 0 {
			x.significand[0] = 1
			x.exponent++
			return
		} else {
			x.significand[p-1]++
		}
	}

	// now unwind to the left to make sure we don't have any lingering carry
	for i := int(p) - 1; i >= 0; i-- {
		if x.significand[i] < 10 {
			break
		}
		x.significand[i] -= 10

		if i == 0 {
			// pad
			x.significand = append([]byte{1}, x.significand...)
			x.exponent++
			break
		}
		x.significand[i-1]++
	}
}

// Return the rounded integer part of a real number.
func (x *Real) RoundedInteger() *Real {
	z := x.Copy()

	if z.IsInf() || z.IsNaN() || z.IsInteger() {
		return z
	}

	if z.exponent < 0 {
		z.roundTo(1)
		if z.exponent == -1 {
			d := z.significand[0]
			switch {
			case d < 5:
				z.SetInt64(0)
				z.negative = false
			case d > 5:
				z.SetInt64(1)
				z.negative = x.negative
			case d == 5:
				switch z.mode {
				case ModeNearestEven, ModeZero:
					z.SetInt64(0)
					z.negative = false
				case ModeNearest:
					z.SetInt64(1)
					z.negative = x.negative
				}
			}
		} else if z.exponent != 0 {
			z.SetInt64(0)
			z.negative = false
		}
	} else {
		z.roundTo(uint(z.exponent) + 1)
	}
	return z
}
