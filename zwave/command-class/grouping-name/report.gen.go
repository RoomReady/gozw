// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package groupingname

import "errors"

// <no value>

type GroupingNameReport struct {
	GroupingIdentifier byte

	Properties1 struct {
		CharPresentation byte
	}

	GroupingName string
}

func (cmd *GroupingNameReport) UnmarshalBinary(payload []byte) error {
	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.GroupingIdentifier = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.CharPresentation = (payload[i] & 0x07)

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.GroupingName = string(payload[i : i+16])

	i += 16

	return nil
}
