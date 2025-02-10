package real

import "testing"

func TestExp1(t *testing.T) {
	x := new(Real)
	x.SetUint64(13)
	x.round()

	z := x.Exp()
	z.precision = 8
	z.round()

	if z.String() != "442413.39200892" {
		t.Fatal("invalid exp", z.String())
	}
}

func TestFactorial(t *testing.T) {
	z := factorial(0)
	if z.String() != "1" {
		t.Fatal("invalid factorial")
	}
	z = factorial(1)
	if z.String() != "1" {
		t.Fatal("invalid factorial")
	}
	z = factorial(24)
	if z.String() != "620448401733239439360000" {
		t.Fatal("invalid factorial", z.String())
	}
}

func TestIPow(t *testing.T) {
	z := new(Real)
	z.SetUint64(5)

	z = z.ipow(0)
	if z.String() != "1" {
		t.Fatal("invalid power")
	}
	z.SetUint64(5)
	z = z.ipow(1)
	if z.String() != "5" {
		t.Fatal("invalid power")
	}
	z.SetUint64(5)
	z = z.ipow(50)
	if z.String() != "88817841970012523233890533447265625" {
		t.Fatal("invalid factorial", z.String())
	}
}

func TestLn(t *testing.T) {
	x := new(Real)
	x.SetUint64(13)

	z := x.Ln()
	z.precision = 20
	z.round()

	if z.String() != "442413.39200892" {
		t.Fatal("invalid exp", z.String())
	}
}

func TestLn2(t *testing.T) {
	x := new(Real)
	x.SetFloat64(.014)

	z := x.Ln()
	z.precision = 20
	z.round()

	if z.String() != "442413.39200892" {
		t.Fatal("invalid exp", z.String())
	}
}
