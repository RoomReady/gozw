// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package metertblmonitorv2

import (
	"encoding/binary"
	"errors"
)

// <no value>

type MeterTblStatusDateGet struct {
	MaximumReports byte

	StartYear uint16

	StartMonth byte

	StartDay byte

	StartHourLocalTime byte

	StartMinuteLocalTime byte

	StartSecondLocalTime byte

	StopYear uint16

	StopMonth byte

	StopDay byte

	StopHourLocalTime byte

	StopMinuteLocalTime byte

	StopSecondLocalTime byte
}

func (cmd *MeterTblStatusDateGet) UnmarshalBinary(payload []byte) error {
	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.MaximumReports = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StartYear = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StartMonth = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StartDay = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StartHourLocalTime = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StartMinuteLocalTime = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StartSecondLocalTime = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StopYear = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StopMonth = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StopDay = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StopHourLocalTime = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StopMinuteLocalTime = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StopSecondLocalTime = payload[i]
	i++

	return nil
}
