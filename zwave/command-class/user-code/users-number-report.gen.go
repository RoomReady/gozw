// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package usercode

import "errors"

// <no value>

type UsersNumberReport struct {
	SupportedUsers byte
}

func (cmd *UsersNumberReport) UnmarshalBinary(payload []byte) error {
	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.SupportedUsers = payload[i]
	i++

	return nil
}
