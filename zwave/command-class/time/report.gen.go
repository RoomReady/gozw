// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package time

import "errors"

// <no value>

type TimeReport struct {
	HourLocalTime struct {
		HourLocalTime byte

		RtcFailure bool
	}

	MinuteLocalTime byte

	SecondLocalTime byte
}

func (cmd *TimeReport) UnmarshalBinary(payload []byte) error {
	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.HourLocalTime.HourLocalTime = (payload[i] & 0x1F)

	if payload[i]&0x80 == 0x80 {
		cmd.HourLocalTime.RtcFailure = true
	} else {
		cmd.HourLocalTime.RtcFailure = false
	}

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.MinuteLocalTime = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.SecondLocalTime = payload[i]
	i++

	return nil
}
