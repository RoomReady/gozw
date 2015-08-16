// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package thermostatfanmodev2

import "errors"

// <no value>

type ThermostatFanModeSet struct {
	Level struct {
		Off bool

		FanMode byte
	}
}

func (cmd *ThermostatFanModeSet) UnmarshalBinary(payload []byte) error {
	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Level.FanMode = (payload[i] & 0x0F)

	if payload[i]&0x80 == 0x80 {
		cmd.Level.Off = true
	} else {
		cmd.Level.Off = false
	}

	i += 1

	return nil
}
