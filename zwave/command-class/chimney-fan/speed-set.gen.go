// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package chimneyfan

import "errors"

// <no value>

type ChimneyFanSpeedSet struct {
	Speed byte
}

func (cmd *ChimneyFanSpeedSet) UnmarshalBinary(payload []byte) error {
	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Speed = payload[i]
	i++

	return nil
}
