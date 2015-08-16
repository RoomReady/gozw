// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package multichannelassociationv2

import "errors"

// <no value>

type MultiChannelAssociationRemove struct {
	GroupingIdentifier byte

	NodeId []byte
}

func (cmd *MultiChannelAssociationRemove) UnmarshalBinary(payload []byte) error {
	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.GroupingIdentifier = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	{
		markerIndex := i
		for ; markerIndex < len(payload) && payload[markerIndex] != 0x00; markerIndex++ {
		}
		val.NodeId = payload[i:markerIndex]
	}

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	i += 1 // skipping MARKER

	return nil
}
