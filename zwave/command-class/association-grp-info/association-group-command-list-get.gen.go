// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package associationgrpinfo

import "errors"

// <no value>

type AssociationGroupCommandListGet struct {
	Properties1 struct {
		AllowCache bool
	}

	GroupingIdentifier byte
}

func (cmd *AssociationGroupCommandListGet) UnmarshalBinary(payload []byte) error {
	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	if payload[i]&0x80 == 0x80 {
		cmd.Properties1.AllowCache = true
	} else {
		cmd.Properties1.AllowCache = false
	}

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.GroupingIdentifier = payload[i]
	i++

	return nil
}
