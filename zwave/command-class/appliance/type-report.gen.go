// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package appliance

// <no value>

type ApplianceTypeReport struct {
	ApplianceType byte

	ApplianceModeSupportedBitmask byte
}

func ParseApplianceTypeReport(payload []byte) ApplianceTypeReport {
	val := ApplianceTypeReport{}

	i := 2

	val.ApplianceType = (payload[i] & 0x3F)

	i += 1

	val.ApplianceModeSupportedBitmask = payload[i]
	i++

	return val
}