// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package language

import (
	"encoding/binary"
	"errors"
)

// <no value>

type LanguageReport struct {
	Language uint32

	Country uint16
}

func (cmd *LanguageReport) UnmarshalBinary(payload []byte) error {
	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Language = binary.BigEndian.Uint32(payload[i : i+3])
	i += 3

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Country = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	return nil
}
