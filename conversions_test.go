// Copyright 2024 David Fritz. All rights reserved.
// This software may be modified and distributed under the terms of the BSD
// 2-clause license. See the LICENSE file for details.

package number

import (
	"fmt"
	"testing"
)

func TestFormatterDecimal1(t *testing.T) {
	x := NewInt64(1234)
	x.exponent = 1 // 12.34

	if fmt.Sprintf("%d", x) != "12" {
		t.Fatal("invalid format", fmt.Sprintf("%d", x))
	}
}

func TestFormatterDecimal2(t *testing.T) {
	x := NewInt64(1234)
	x.exponent = 6 // 1234000

	if fmt.Sprintf("%d", x) != "1234000" {
		t.Fatal("invalid format", fmt.Sprintf("%d", x))
	}
}

func TestFormatterFloat1(t *testing.T) {
	x := NewInt64(0)

	if fmt.Sprintf("%f", x) != "0.0" {
		t.Fatal("invalid format", fmt.Sprintf("%f", x))
	}
}

func TestFormatterFloat2(t *testing.T) {
	x := NewInt64(1234)
	x.exponent = -2 // .01234

	if fmt.Sprintf("%f", x) != "0.01234" {
		t.Fatal("invalid format", fmt.Sprintf("%f", x))
	}
}

func TestFormatterFloat3(t *testing.T) {
	x := NewInt64(1234)
	x.exponent = 0 // 1.234

	if fmt.Sprintf("%f", x) != "1.234" {
		t.Fatal("invalid format", fmt.Sprintf("%f", x))
	}
}

func TestFormatterFloat4(t *testing.T) {
	x := NewInt64(1234)
	x.exponent = 6 // 1234000

	if fmt.Sprintf("%f", x) != "1234000.0" {
		t.Fatal("invalid format", fmt.Sprintf("%f", x))
	}
}

func TestFormatterV1(t *testing.T) {
	x := NewInt64(1234)
	x.exponent = 6 // 1234000

	if fmt.Sprintf("%v", x) != "1234000" {
		t.Fatal("invalid format", fmt.Sprintf("%v", x))
	}
}

func TestFormatterV2(t *testing.T) {
	x := NewInt64(1234)
	x.exponent = 100

	if fmt.Sprintf("%v", x) != "1.234e100" {
		t.Fatal("invalid format", fmt.Sprintf("%v", x))
	}
}

func TestFormatterV3(t *testing.T) {
	x := NewInt64(1234)
	x.exponent = 0

	if fmt.Sprintf("%v", x) != "1.234" {
		t.Fatal("invalid format", fmt.Sprintf("%v", x))
	}
}

func TestFormatterV4(t *testing.T) {
	x := NewInt64(1234)
	x.exponent = -100

	if fmt.Sprintf("%v", x) != "1.234e-100" {
		t.Fatal("invalid format", fmt.Sprintf("%v", x))
	}
}

func TestUint641(t *testing.T) {
	x := NewInt64(1234)
	x.exponent = 10

	u, ok := x.Uint64()
	if !ok || u != 12340000000 {
		t.Fatal("invalid cast", u, ok)
	}
}

func TestUint642(t *testing.T) {
	x := NewInt64(1234)
	x.exponent = -10

	u, ok := x.Uint64()
	if !ok || u != 0 {
		t.Fatal("invalid cast", u, ok)
	}
}

func TestInt641(t *testing.T) {
	x := NewInt64(-1234)
	x.exponent = 10

	u, ok := x.Int64()
	if !ok || u != -12340000000 {
		t.Fatal("invalid cast", u, ok)
	}
}

func TestInt642(t *testing.T) {
	x := NewInt64(-1234)
	x.exponent = -10

	u, ok := x.Int64()
	if !ok || u != 0 {
		t.Fatal("invalid cast", u, ok)
	}
}

func TestFloat641(t *testing.T) {
	x := NewInt64(-1234)
	x.exponent = 10

	u, ok := x.Float64()
	if !ok || u != -1.234e+10 {
		t.Fatal("invalid cast", u, ok)
	}
}

func TestFloat642(t *testing.T) {
	x := NewInt64(-1234)
	x.exponent = -10

	u, ok := x.Float64()
	if !ok || u != -1.234e-10 {
		t.Fatal("invalid cast", u, ok)
	}
}
