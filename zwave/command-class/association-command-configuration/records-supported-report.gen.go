// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package associationcommandconfiguration

import (
	"encoding/binary"
	"errors"
)

// <no value>

type CommandRecordsSupportedReport struct {
	Properties1 struct {
		MaxCommandLength byte

		ConfCmd bool

		Vc bool
	}

	FreeCommandRecords uint16

	MaxCommandRecords uint16
}

func (cmd *CommandRecordsSupportedReport) UnmarshalBinary(payload []byte) error {
	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.MaxCommandLength = (payload[i] & 0xFC) << 2

	if payload[i]&0x01 == 0x01 {
		cmd.Properties1.ConfCmd = true
	} else {
		cmd.Properties1.ConfCmd = false
	}

	if payload[i]&0x02 == 0x02 {
		cmd.Properties1.Vc = true
	} else {
		cmd.Properties1.Vc = false
	}

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.FreeCommandRecords = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.MaxCommandRecords = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	return nil
}
