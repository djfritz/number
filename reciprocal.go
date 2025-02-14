package real

import (
	"strconv"
)

func (x *Real) Reciprocal() *Real {
	x2 := x.Copy()
	x2.SetPrecision(2 * x.precision)
	z := x2.reciprocal()
	z.SetPrecision(x.precision)
	return z
}

func (x *Real) reciprocal() *Real {
	xscaled := x.Copy()
	xscaled.exponent = 0

	s := xscaled.String()
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic("could not parse float")
	}
	f = 1 / f

	z0 := initFrom(x)
	z0.SetFloat64(f)

	z := z0
	two := initFrom(x)
	two.SetInt64(2)

	for i := 0; i < 10; i++ {
		zn := z.mul(two.Sub(xscaled.mul(z)))
		z = zn
	}

	z.exponent += x.exponent * -1
	z.round()
	return z
}
