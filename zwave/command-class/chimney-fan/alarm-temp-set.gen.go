// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package chimneyfan

import "errors"

// <no value>

type ChimneyFanAlarmTempSet struct {
	Properties1 struct {
		Size byte

		Scale byte

		Precision byte
	}

	Value []byte
}

func (cmd *ChimneyFanAlarmTempSet) UnmarshalBinary(payload []byte) error {
	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.Size = (payload[i] & 0x07)

	cmd.Properties1.Scale = (payload[i] & 0x18) << 3

	cmd.Properties1.Precision = (payload[i] & 0xE0) << 5

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Value = payload[i : i+0]
	i += 0

	return nil
}
