// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package multichannelv2

import "errors"

// <no value>

type MultiChannelEndPointFindReport struct {
	ReportsToFollow byte

	GenericDeviceClass byte

	SpecificDeviceClass byte
}

func (cmd *MultiChannelEndPointFindReport) UnmarshalBinary(payload []byte) error {
	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ReportsToFollow = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.GenericDeviceClass = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.SpecificDeviceClass = payload[i]
	i++

	return nil
}
