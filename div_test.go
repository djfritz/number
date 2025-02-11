package real

import "testing"

func TestDiv1(t *testing.T) {
	x := NewInt64(10)
	y := NewInt64(5)

	q := x.Div(y)
	q.SetPrecision(8)

	if q.Compare(NewInt64(2)) != 0 {
		t.Fatal("invalid div", q)
	}
}

func TestDiv2(t *testing.T) {
	x := NewInt64(2)
	y := NewInt64(50)

	q := x.Div(y)
	q.SetPrecision(8)

	expected := NewInt64(4)
	expected.exponent = -2
	if q.Compare(expected) != 0 {
		t.Fatal("invalid div", q)
	}
}

func TestDiv3(t *testing.T) {
	x := NewInt64(23)
	y := NewInt64(5011513)

	q := x.Div(y)
	q.SetPrecision(8)

	if q.String() != "4.5894324e-6" {
		t.Fatal("invalid div", q)
	}
}
