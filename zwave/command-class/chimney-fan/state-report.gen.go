// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package chimneyfan

import "errors"

// <no value>

type ChimneyFanStateReport struct {
	State byte
}

func (cmd *ChimneyFanStateReport) UnmarshalBinary(payload []byte) error {
	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.State = payload[i]
	i++

	return nil
}
