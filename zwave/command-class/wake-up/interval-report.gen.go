// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package wakeup

import (
	"encoding/binary"
	"errors"
)

// <no value>

type WakeUpIntervalReport struct {
	Seconds uint32

	Nodeid byte
}

func (cmd *WakeUpIntervalReport) UnmarshalBinary(payload []byte) error {
	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Seconds = binary.BigEndian.Uint32(payload[i : i+3])
	i += 3

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Nodeid = payload[i]
	i++

	return nil
}
