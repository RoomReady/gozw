// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package protectionv2

import "errors"

// <no value>

type ProtectionTimeoutSet struct {
	Timeout byte
}

func (cmd *ProtectionTimeoutSet) UnmarshalBinary(payload []byte) error {
	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Timeout = payload[i]
	i++

	return nil
}
