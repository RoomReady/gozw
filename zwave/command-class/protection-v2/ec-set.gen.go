// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package protectionv2

import "errors"

// <no value>

type ProtectionEcSet struct {
	NodeId byte
}

func (cmd *ProtectionEcSet) UnmarshalBinary(payload []byte) error {
	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.NodeId = payload[i]
	i++

	return nil
}
