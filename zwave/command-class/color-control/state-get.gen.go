// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package colorcontrol

// <no value>

type StateGet struct {
	CapabilityId byte
}

func ParseStateGet(payload []byte) StateGet {
	val := StateGet{}

	i := 2

	val.CapabilityId = payload[i]
	i++

	return val
}