package real

import (
	"math"
	"strconv"
)

const MaxLnIterations = 1000

// OEIS A002392 ln(10)
var ln10digits = []byte{2, 3, 0, 2, 5, 8, 5, 0, 9, 2, 9, 9, 4, 0, 4, 5, 6, 8, 4, 0, 1, 7, 9, 9, 1, 4, 5, 4, 6, 8, 4, 3, 6, 4, 2, 0, 7, 6, 0, 1, 1, 0, 1, 4, 8, 8, 6, 2, 8, 7, 7, 2, 9, 7, 6, 0, 3, 3, 3, 2, 7, 9, 0, 0, 9, 6, 7, 5, 7, 2, 6, 0, 9, 6, 7, 7, 3, 5, 2, 4, 8, 0, 2, 3, 5, 9, 9}

func (x *Real) Ln() *Real {
	// z1 = z0 * 2*((x-exp(z0))/(x+exp(z0)))

	xscaled := x.Copy()
	xscaled.exponent = 0

	s := xscaled.String()
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic("could not parse float")
	}
	f = math.Log(f)

	z0 := initFrom(x)
	z0.SetFloat64(f)

	z := z0
	two := initFrom(x)
	two.SetInt64(2)

	for i := 0; i < MaxLnIterations; i++ {
		ez := z.Exp()
		n := xscaled.Sub(ez)
		d := xscaled.Add(ez)
		q := n.Div(d)
		q2 := two.Mul(q)
		znext := z.Add(q2)
		if znext.Compare(z) == 0 {
			z = znext
			break
		}
		z = znext
	}

	// exponent part
	if x.exponent != 0 {
		ln10 := initFrom(x)
		ln10.significand = ln10digits
		ln10.round()
		e := initFrom(x)
		e.SetInt64(int64(x.exponent))
		eln10 := e.Mul(ln10)
		z = z.Add(eln10)
	}
	return z
}
