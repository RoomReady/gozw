// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package climatecontrolschedule

import (
	"encoding/binary"
	"errors"
)

// <no value>

type ScheduleReport struct {
	Properties1 struct {
		Weekday byte
	}

	Switchpoint0 uint32

	Switchpoint1 uint32

	Switchpoint2 uint32

	Switchpoint3 uint32

	Switchpoint4 uint32

	Switchpoint5 uint32

	Switchpoint6 uint32

	Switchpoint7 uint32

	Switchpoint8 uint32
}

func (cmd *ScheduleReport) UnmarshalBinary(payload []byte) error {
	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.Weekday = (payload[i] & 0x07)

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Switchpoint0 = binary.BigEndian.Uint32(payload[i : i+3])
	i += 3

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Switchpoint1 = binary.BigEndian.Uint32(payload[i : i+3])
	i += 3

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Switchpoint2 = binary.BigEndian.Uint32(payload[i : i+3])
	i += 3

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Switchpoint3 = binary.BigEndian.Uint32(payload[i : i+3])
	i += 3

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Switchpoint4 = binary.BigEndian.Uint32(payload[i : i+3])
	i += 3

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Switchpoint5 = binary.BigEndian.Uint32(payload[i : i+3])
	i += 3

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Switchpoint6 = binary.BigEndian.Uint32(payload[i : i+3])
	i += 3

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Switchpoint7 = binary.BigEndian.Uint32(payload[i : i+3])
	i += 3

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Switchpoint8 = binary.BigEndian.Uint32(payload[i : i+3])
	i += 3

	return nil
}
