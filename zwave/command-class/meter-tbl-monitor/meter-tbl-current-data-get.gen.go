// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package metertblmonitor

import (
	"encoding/binary"
	"errors"
)

// <no value>

type MeterTblCurrentDataGet struct {
	DatasetRequested uint32
}

func (cmd *MeterTblCurrentDataGet) UnmarshalBinary(payload []byte) error {
	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.DatasetRequested = binary.BigEndian.Uint32(payload[i : i+3])
	i += 3

	return nil
}
