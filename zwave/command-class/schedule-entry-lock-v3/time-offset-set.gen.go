// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package scheduleentrylockv3

import "errors"

// <no value>

type ScheduleEntryLockTimeOffsetSet struct {
	Level struct {
		HourTzo byte

		SignTzo bool
	}

	MinuteTzo byte

	Level2 struct {
		MinuteOffsetDst byte

		SignOffsetDst bool
	}
}

func (cmd *ScheduleEntryLockTimeOffsetSet) UnmarshalBinary(payload []byte) error {
	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Level.HourTzo = (payload[i] & 0x7F)

	if payload[i]&0x80 == 0x80 {
		cmd.Level.SignTzo = true
	} else {
		cmd.Level.SignTzo = false
	}

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.MinuteTzo = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Level2.MinuteOffsetDst = (payload[i] & 0x7F)

	if payload[i]&0x80 == 0x80 {
		cmd.Level2.SignOffsetDst = true
	} else {
		cmd.Level2.SignOffsetDst = false
	}

	i += 1

	return nil
}
