// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package scenecontrollerconf

import "errors"

// <no value>

type SceneControllerConfSet struct {
	GroupId byte

	SceneId byte

	DimmingDuration byte
}

func (cmd *SceneControllerConfSet) UnmarshalBinary(payload []byte) error {
	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.GroupId = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.SceneId = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.DimmingDuration = payload[i]
	i++

	return nil
}
