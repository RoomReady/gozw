// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package securitypanelzonesensor

import "errors"

// <no value>

type SecurityPanelZoneSensorInstalledGet struct {
	ZoneNumber byte
}

func (cmd *SecurityPanelZoneSensorInstalledGet) UnmarshalBinary(payload []byte) error {
	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ZoneNumber = payload[i]
	i++

	return nil
}
