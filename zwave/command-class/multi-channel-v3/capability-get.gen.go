// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package multichannelv3

// <no value>

type MultiChannelCapabilityGet struct {
	EndPoint byte
}

func ParseMultiChannelCapabilityGet(payload []byte) MultiChannelCapabilityGet {
	val := MultiChannelCapabilityGet{}

	i := 2

	val.EndPoint = (payload[i] & 0x7F)

	i += 1

	return val
}