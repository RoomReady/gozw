// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package switchmultilevel

import "errors"

// <no value>

type SwitchMultilevelStartLevelChange struct {
	Level struct {
		IgnoreStartLevel bool

		UpDown bool
	}

	StartLevel byte
}

func (cmd *SwitchMultilevelStartLevelChange) UnmarshalBinary(payload []byte) error {
	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	if payload[i]&0x20 == 0x20 {
		cmd.Level.IgnoreStartLevel = true
	} else {
		cmd.Level.IgnoreStartLevel = false
	}

	if payload[i]&0x40 == 0x40 {
		cmd.Level.UpDown = true
	} else {
		cmd.Level.UpDown = false
	}

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.StartLevel = payload[i]
	i++

	return nil
}
