// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package metertblmonitorv2

import "errors"

// <no value>

type MeterTblTableIdReport struct {
	Properties1 struct {
		NumberOfCharacters byte
	}

	MeterIdCharacter []byte
}

func (cmd *MeterTblTableIdReport) UnmarshalBinary(payload []byte) error {
	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.NumberOfCharacters = (payload[i] & 0x1F)

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.MeterIdCharacter = payload[i : i+0]
	i += 0

	return nil
}
