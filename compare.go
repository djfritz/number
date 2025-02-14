package real

import "bytes"

func (x *Real) Compare(y *Real) int {
	if !x.negative && y.negative {
		return 1
	} else if x.negative && !y.negative {
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
