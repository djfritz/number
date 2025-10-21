// Copyright 2025 David Fritz. All rights reserved.
// This software may be modified and distributed under the terms of the BSD
// 2-clause license. See the LICENSE file for details.

package number

import "testing"

func TestArctan1(t *testing.T) {
	x, _ := ParseReal(".5", DefaultPrecision)
	z := x.Arctan()

	if z.String() != "4.636476090008061162142562314612144e-1" {
		t.Fatal("invalid arctan", z.String())
	}
}

func TestArctan2(t *testing.T) {
	x, _ := ParseReal("-0.5", DefaultPrecision)
	z := x.Arctan()

	if z.String() != "-4.636476090008061162142562314612144e-1" {
		t.Fatal("invalid arctan", z.String())
	}
}

func TestArctan3(t *testing.T) {
	x, _ := ParseReal("1", DefaultPrecision)
	z := x.Arctan()

	if z.String() != "7.853981633974483096156608458198757e-1" {
		t.Fatal("invalid arctan", z.String())
	}
}

func TestArctan4(t *testing.T) {
	x, _ := ParseReal("1000000000", DefaultPrecision)
	z := x.Arctan()

	if z.String() != "1.570796325794896619231321691973085e0" {
		t.Fatal("invalid arctan", z.String())
	}
}

func TestAtan2_1(t *testing.T) {
	y, _ := ParseReal("0", DefaultPrecision)
	x, _ := ParseReal("1", DefaultPrecision)
	z := Atan2(y, x)

	if z.String() != "0" {
		t.Fatal("invalid arctan", z.String())
	}
}

func TestAtan2_2(t *testing.T) {
	y, _ := ParseReal("0.179", DefaultPrecision)
	x, _ := ParseReal("0.063", DefaultPrecision)
	z := Atan2(y, x)

	if z.String() != "1.23238064868431367733228638812685e0" {
		t.Fatal("invalid arctan", z.String())
	}
}

func TestAtan2_3(t *testing.T) {
	y, _ := ParseReal("0.179", DefaultPrecision)
	x, _ := ParseReal("-0.063", DefaultPrecision)
	z := Atan2(y, x)

	if z.String() != "1.909212004905479561130356995152653e0" {
		t.Fatal("invalid arctan", z.String())
	}
}
