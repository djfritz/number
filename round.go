// Copyright 2025 David Fritz. All rights reserved.
// This software may be modified and distributed under the terms of the BSD
// 2-clause license. See the LICENSE file for details.

package number

import "errors"

const (
	ModeNearestEven = iota
	ModeNearest
	ModeUp
	ModeDown
	ModeZero
)

var ErrInvalidMode = errors.New("invalid mode")

func (x *Real) SetMode(m int) error {
	switch m {
	case ModeNearestEven:
		fallthrough
	case ModeNearest:
		fallthrough
	case ModeUp:
		fallthrough
	case ModeDown:
		fallthrough
	case ModeZero:
		x.mode = m
	default:
		return ErrInvalidMode
	}
	return nil
}

func (x *Real) Mode() int {
	return x.mode
}

// Round the value to the set precision and rounding mode, if necessary.
func (x *Real) round() {
	x.validate()
	x.roundTo(x.precision)
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
	case ModeUp:
	case ModeDown:
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
	case d >= 5:
		// round up
		if p == 0 {
			x.significand[0] = 1
			x.exponent++
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
			break
		}
		x.significand[i-1]++
	}
}
