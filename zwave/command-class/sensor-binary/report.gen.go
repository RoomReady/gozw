// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package sensorbinary

import "errors"

// <no value>

type SensorBinaryReport struct {
	SensorValue byte
}

func (cmd *SensorBinaryReport) UnmarshalBinary(payload []byte) error {
	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.SensorValue = payload[i]
	i++

	return nil
}
