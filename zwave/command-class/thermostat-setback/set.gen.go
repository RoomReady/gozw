// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package thermostatsetback

import "errors"

// <no value>

type ThermostatSetbackSet struct {
	Properties1 struct {
		SetbackType byte
	}

	SetbackState byte
}

func (cmd *ThermostatSetbackSet) UnmarshalBinary(payload []byte) error {
	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.SetbackType = (payload[i] & 0x03)

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.SetbackState = payload[i]
	i++

	return nil
}
