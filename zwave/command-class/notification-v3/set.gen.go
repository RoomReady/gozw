// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package notificationv3

import "errors"

// <no value>

type NotificationSet struct {
	NotificationType byte

	NotificationStatus byte
}

func (cmd *NotificationSet) UnmarshalBinary(payload []byte) error {
	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.NotificationType = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.NotificationStatus = payload[i]
	i++

	return nil
}
