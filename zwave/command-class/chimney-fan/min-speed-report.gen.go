// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package chimneyfan

import "errors"

// <no value>

type ChimneyFanMinSpeedReport struct {
	MinSpeed byte
}

func (cmd *ChimneyFanMinSpeedReport) UnmarshalBinary(payload []byte) error {
	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.MinSpeed = payload[i]
	i++

	return nil
}
