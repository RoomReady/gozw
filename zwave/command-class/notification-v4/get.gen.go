// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package notificationv4

import "errors"

// <no value>

type NotificationGet struct {
	V1AlarmType byte

	NotificationType byte

	Event byte
}

func (cmd *NotificationGet) UnmarshalBinary(payload []byte) error {
	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.V1AlarmType = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.NotificationType = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Event = payload[i]
	i++

	return nil
}
