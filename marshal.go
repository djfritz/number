// Copyright 2025 David Fritz. All rights reserved.
// This software may be modified and distributed under the terms of the BSD
// 2-clause license. See the LICENSE file for details.

package number

import (
	"bytes"
	"encoding/gob"
)

func (x *Real) GobEncode() ([]byte, error) {
	w := bytes.Buffer{}
	enc := gob.NewEncoder(&w)
	var err error

	if x.significand == nil {
		err = enc.Encode(false)
		if err != nil {
			return nil, err
		}
	} else {
		err = enc.Encode(true)
		if err != nil {
			return nil, err
		}
		err = enc.Encode(x.significand)
		if err != nil {
			return nil, err
		}
	}
	err = enc.Encode(x.negative)
	if err != nil {
		return nil, err
	}
	err = enc.Encode(x.exponent)
	if err != nil {
		return nil, err
	}
	err = enc.Encode(x.precision)
	if err != nil {
		return nil, err
	}
	err = enc.Encode(x.form)
	if err != nil {
		return nil, err
	}
	err = enc.Encode(x.mode)
	if err != nil {
		return nil, err
	}

	return w.Bytes(), nil

}

func (x *Real) GobDecode(b []byte) error {
	r := bytes.NewReader(b)
	dec := gob.NewDecoder(r)

	var hasS bool
	err := dec.Decode(&hasS)
	if err != nil {
		return err
	}
	if hasS {
		err = dec.Decode(&x.significand)
		if err != nil {
			return err
		}
	}
	err = dec.Decode(&x.negative)
	if err != nil {
		return err
	}
	err = dec.Decode(&x.exponent)
	if err != nil {
		return err
	}
	err = dec.Decode(&x.precision)
	if err != nil {
		return err
	}
	err = dec.Decode(&x.form)
	if err != nil {
		return err
	}
	err = dec.Decode(&x.mode)
	if err != nil {
		return err
	}

	return nil
}
