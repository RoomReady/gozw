// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package doorlock

import "errors"

// <no value>

type DoorLockOperationReport struct {
	DoorLockMode byte

	Properties1 struct {
		InsideDoorHandlesMode byte

		OutsideDoorHandlesMode byte
	}

	DoorCondition byte

	LockTimeoutMinutes byte

	LockTimeoutSeconds byte
}

func (cmd *DoorLockOperationReport) UnmarshalBinary(payload []byte) error {
	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.DoorLockMode = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.InsideDoorHandlesMode = (payload[i] & 0x0F)

	cmd.Properties1.OutsideDoorHandlesMode = (payload[i] & 0xF0) << 4

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.DoorCondition = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.LockTimeoutMinutes = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.LockTimeoutSeconds = payload[i]
	i++

	return nil
}
