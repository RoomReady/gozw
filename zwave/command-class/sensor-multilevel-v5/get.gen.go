// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package sensormultilevelv5

import "errors"

// <no value>

type SensorMultilevelGet struct {
	SensorType byte

	Properties1 struct {
		Scale byte
	}
}

func (cmd *SensorMultilevelGet) UnmarshalBinary(payload []byte) error {
	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.SensorType = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.Scale = (payload[i] & 0x18) << 3

	i += 1

	return nil
}
