// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package proprietary

import "errors"

// <no value>

type ProprietaryGet struct {
	Data []byte
}

func (cmd *ProprietaryGet) UnmarshalBinary(payload []byte) error {
	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	val.Data = payload[i:]

	return nil
}
