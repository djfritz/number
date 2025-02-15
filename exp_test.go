package real

import "testing"

func TestExp1(t *testing.T) {
	x := NewUint64(5)
	z := x.Exp()

	if z.String() != "1.484131591025766034211155800405523e2" {
		t.Fatal("invalid exp", z)
	}
}

func TestExp2(t *testing.T) {
	x := NewUint64(1)
	x.exponent = -2 // .01
	z := x.Exp()

	if z.String() != "1.01005016708416805754216545690286e0" {
		t.Fatal("invalid exp", z)
	}
}

func TestExp3(t *testing.T) {
	x := NewUint64(100)
	z := x.Exp()

	if z.String() != "2.688117141816135448412625551580014e43" {
		t.Fatal("invalid exp", z)
	}
}

func TestExp4(t *testing.T) {
	x := NewUint64(200)
	z := x.Exp()

	if z.String() != "7.225973768125749258177477042189306e86" {
		t.Fatal("invalid exp", z)
	}
}

func TestExp5(t *testing.T) {
	x := NewUint64(500)
	z := x.Exp()

	if z.String() != "1.403592217852837410739770332840912e217" {
		t.Fatal("invalid exp", z)
	}
}
