// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package lock

import "errors"

// <no value>

type LockReport struct {
	LockState byte
}

func (cmd *LockReport) UnmarshalBinary(payload []byte) error {
	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.LockState = payload[i]
	i++

	return nil
}
