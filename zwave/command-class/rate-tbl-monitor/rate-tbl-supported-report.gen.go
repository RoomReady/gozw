// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package ratetblmonitor

import (
	"encoding/binary"
	"errors"
)

// <no value>

type RateTblSupportedReport struct {
	RatesSupported byte

	ParameterSetSupportedBitMask uint16
}

func (cmd *RateTblSupportedReport) UnmarshalBinary(payload []byte) error {
	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.RatesSupported = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ParameterSetSupportedBitMask = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	return nil
}
