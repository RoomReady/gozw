// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package hrvcontrol

import "errors"

// <no value>

type HrvControlModeSupportedReport struct {
	Properties1 struct {
		ManualControlSupported byte
	}

	BitMask byte
}

func (cmd *HrvControlModeSupportedReport) UnmarshalBinary(payload []byte) error {
	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.ManualControlSupported = (payload[i] & 0x0F)

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.BitMask = payload[i]
	i++

	return nil
}
