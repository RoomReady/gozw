// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package thermostatoperatingstatev2

import "errors"

// <no value>

type ThermostatOperatingStateLoggingGet struct {
	BitMask byte
}

func (cmd *ThermostatOperatingStateLoggingGet) UnmarshalBinary(payload []byte) error {
	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.BitMask = payload[i]
	i++

	return nil
}
