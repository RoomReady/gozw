// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package switchall

import "errors"

// <no value>

type SwitchAllReport struct {
	Mode byte
}

func (cmd *SwitchAllReport) UnmarshalBinary(payload []byte) error {
	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Mode = payload[i]
	i++

	return nil
}
