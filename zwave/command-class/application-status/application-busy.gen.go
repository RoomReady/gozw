// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package applicationstatus

import "errors"

// <no value>

type ApplicationBusy struct {
	Status byte

	WaitTime byte
}

func (cmd *ApplicationBusy) UnmarshalBinary(payload []byte) error {
	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Status = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.WaitTime = payload[i]
	i++

	return nil
}
