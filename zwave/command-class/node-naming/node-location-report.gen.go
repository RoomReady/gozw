// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package nodenaming

import "errors"

// <no value>

type NodeNamingNodeLocationReport struct {
	Level struct {
		CharPresentation byte
	}

	NodeLocationChar string
}

func (cmd *NodeNamingNodeLocationReport) UnmarshalBinary(payload []byte) error {
	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Level.CharPresentation = (payload[i] & 0x07)

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.NodeLocationChar = string(payload[i : i+16])

	i += 16

	return nil
}
