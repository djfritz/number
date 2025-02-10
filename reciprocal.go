package real

import (
	"strconv"
)

func (r *Real) reciprocal() *Real {
	// scale r to be 0 < r < 1
	rc := r.Copy()
	dshift := uint(len(r.digits)) - r.decimal
	rc.decimal = uint(len(rc.digits))

	// we could do a lot of work to determine a suitable x0 with quadratic
	// linear approximation or the like, but instead we'll just pass thing
	// in and out of a floating point parser, invert that, and stick it
	// back in.
	s := rc.String()
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic("could not parse float")
	}
	f = 1 / f
	x0 := initFrom(r)
	x0.SetFloat64(f)
	//	fmt.Println(x0)

	// f(x) = (1/x) - D
	// xi+1 = xi(2-(D*xi))

	x := x0
	two := initFrom(r)
	two.SetUint64(2)

	for i := 0; i < 100; i++ {
		xn := x.mul(two.sub(rc.mul(x)))
		xn.roundTo(2 * xn.precision)
		if xn.Compare(x) == 0 {
			x = xn
			break
		}
		x = xn
	}

	// restore shift
	x.decimal += dshift
	if x.decimal > uint(len(x.digits)) {
		pad := make([]byte, x.decimal-uint(len(x.digits)))
		x.digits = append(pad, x.digits...)
	}

	x.trim()

	return x
}
