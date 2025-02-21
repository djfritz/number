// Copyright 2025 David Fritz. All rights reserved.
// This software may be modified and distributed under the terms of the BSD
// 2-clause license. See the LICENSE file for details.

package number

const (
	ModeNearestEven = iota
	ModeNearest
	ModeUp
	ModeDown
	ModeZero
)

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
	case ModeUp:
	case ModeDown:
	case ModeZero:
	}
}

func (x *Real) roundToNearestEven(p uint) {
	var hasNonZero bool
	for i := len(x.significand) - 1; i >= int(p); i-- {
		d := x.significand[i]
		switch {
		case d < 5 && d > 0:
			// we have something but we can't decide yet
			hasNonZero = true
		case d > 5 || (hasNonZero && d == 5):
			// round up
			if i == 0 {
				x.significand[0] = 1
				x.exponent++
				return
			} else {
				x.significand[i-1]++
			}
			hasNonZero = false
		case d == 5 && !hasNonZero:
			// round to nearest even
			if x.significand[i-1]%2 != 0 {
				if i == 0 {
					x.significand[0] = 1
					x.exponent++
					return
				} else {
					x.significand[i-1]++
				}
			}
		}
		x.significand[i] = 0
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
