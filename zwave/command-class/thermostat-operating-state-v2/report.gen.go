// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package thermostatoperatingstatev2

import "errors"

// <no value>

type ThermostatOperatingStateReport struct {
	Properties1 struct {
		OperatingState byte
	}
}

func (cmd *ThermostatOperatingStateReport) UnmarshalBinary(payload []byte) error {
	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.OperatingState = (payload[i] & 0x0F)

	i += 1

	return nil
}
