// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package basic

import "errors"

// <no value>

type BasicSet struct {
	Value byte
}

func (cmd *BasicSet) UnmarshalBinary(payload []byte) error {
	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Value = payload[i]
	i++

	return nil
}
