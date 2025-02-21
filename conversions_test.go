// Copyright 2025 David Fritz. All rights reserved.
// This software may be modified and distributed under the terms of the BSD
// 2-clause license. See the LICENSE file for details.

package number

import (
	"fmt"
	"testing"
)

func TestFormatterDecimalLargePrecision(t *testing.T) {
	x := NewUint64(100)
	x.SetPrecision(100)
	x = x.Factorial()
	x.exponent = 0

	if fmt.Sprintf("%.100f", x) != "9.332621544394415268169923885626670049071596826438162146859296389521759999322991560894146397615651822" {
		t.Fatal("invalid format", fmt.Sprintf("%.100f", x))
	}
}

func TestFormatterDecimalLargePrecision2(t *testing.T) {
	x := NewUint64(100)
	x.SetPrecision(100)
	x = x.Factorial()
	x.exponent = 0

	if fmt.Sprintf("%.100v", x) != "9.332621544394415268169923885626670049071596826438162146859296389521759999322991560894146397615651822" {
		t.Fatal("invalid format", fmt.Sprintf("%.100v", x))
	}
}

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

func TestFormatterV5(t *testing.T) {
	x := NewInt64(0)

	if fmt.Sprintf("%v", x) != "0" {
		t.Fatal("invalid format", fmt.Sprintf("%v", x))
	}
}

func TestFormatterV6(t *testing.T) {
	x := NewInt64(-1)

	if fmt.Sprintf("%v", x) != "-1" {
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

func TestParseReal1(t *testing.T) {
	s := "1.234"
	x, err := ParseReal(s, DefaultPrecision)
	if err != nil {
		t.Fatal(err)
	}

	if x.String() != "1.234e0" {
		t.Fatal("invalid parse", x)
	}
}

func TestParseReal2(t *testing.T) {
	s := "1.234e-5"
	x, err := ParseReal(s, DefaultPrecision)
	if err != nil {
		t.Fatal(err)
	}

	if x.String() != "1.234e-5" {
		t.Fatal("invalid parse", x)
	}
}

func TestParseReal3(t *testing.T) {
	s := "-1.234e50"
	x, err := ParseReal(s, DefaultPrecision)
	if err != nil {
		t.Fatal(err)
	}

	if x.String() != "-1.234e50" {
		t.Fatal("invalid parse", x)
	}
}

func TestParseReal4(t *testing.T) {
	s := "9.2342234234234252345232734672364723472342342342523423432456"
	x, err := ParseReal(s, DefaultPrecision)
	if err != nil {
		t.Fatal(err)
	}

	if x.String() != "9.234223423423425234523273467236472e0" {
		t.Fatal("invalid parse", x)
	}
}

func TestParseReal5(t *testing.T) {
	s := "9.2345534234234252345232734672364723472342342342523423432456"
	x, err := ParseReal(s, 5)
	if err != nil {
		t.Fatal(err)
	}

	if x.String() != "9.2346e0" {
		t.Fatal("invalid parse", x)
	}
}

func TestParseReal6(t *testing.T) {
	s := "-"
	_, err := ParseReal(s, 5)
	if err == nil {
		t.Fatal("should have generated error")
	}
}
