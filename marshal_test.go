package real

import (
	"bytes"
	"encoding/gob"
	"testing"
)

func TestGobEncodeDecode(t *testing.T) {
	x := NewInt64(1234)

	b := bytes.Buffer{}

	enc := gob.NewEncoder(&b)
	err := enc.Encode(x)
	if err != nil {
		t.Fatal(err)
	}

	bb := bytes.NewReader(b.Bytes())

	dec := gob.NewDecoder(bb)

	z := new(Real)
	err = dec.Decode(z)
	if err != nil {
		t.Fatal(err)
	}

	if z.String() != "1.234e3" {
		t.Fatal("invalid decode")
	}
}

func TestGobEncodeDecodeNilSignificand(t *testing.T) {
	x := new(Real)

	b := bytes.Buffer{}

	enc := gob.NewEncoder(&b)
	err := enc.Encode(x)
	if err != nil {
		t.Fatal(err)
	}

	bb := bytes.NewReader(b.Bytes())

	dec := gob.NewDecoder(bb)

	z := new(Real)
	err = dec.Decode(z)
	if err != nil {
		t.Fatal(err)
	}

	// If we didn't panic already we're good.
}
