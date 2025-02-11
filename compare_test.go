package real

import "testing"

func TestCompare1(t *testing.T) {
	x := NewInt64(5)
	y := NewInt64(-5)

	if x.Compare(y) != 1 {
		t.Fatal("invalid compare")
	}
}

func TestCompare2(t *testing.T) {
	x := NewInt64(-5)
	y := NewInt64(100)

	if x.Compare(y) != -1 {
		t.Fatal("invalid compare")
	}
}

func TestCompare3(t *testing.T) {
	x := NewInt64(-5)
	y := NewInt64(-5)

	if x.Compare(y) != 0 {
		t.Fatal("invalid compare")
	}
}

func TestCompare4(t *testing.T) {
	x := NewInt64(-5)
	y := NewFloat64(-5)

	if x.Compare(y) != 0 {
		t.Fatal("invalid compare")
	}
}

func TestCompareSignificand1(t *testing.T) {
	x := []byte{1, 2, 3, 4}
	y := []byte{0, 0, 1, 1}

	if compareSignificand(x, y) != 1 {
		t.Fatal("invalid compare")
	}
}

func TestCompareSignificand2(t *testing.T) {
	x := []byte{1, 2, 3, 4}
	y := []byte{3, 0, 1, 1}

	if compareSignificand(x, y) != -1 {
		t.Fatal("invalid compare")
	}
}

func TestCompareSignificand3(t *testing.T) {
	x := []byte{1, 2}
	y := []byte{1, 2, 1, 1}

	if compareSignificand(x, y) != -1 {
		t.Fatal("invalid compare")
	}
}

func TestCompareSignificand4(t *testing.T) {
	x := []byte{1, 2, 1, 1}
	y := []byte{1, 2}

	if compareSignificand(x, y) != 1 {
		t.Fatal("invalid compare")
	}
}
