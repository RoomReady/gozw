// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package lock

// <no value>

type LockReport struct {
	LockState byte
}

func ParseLockReport(payload []byte) LockReport {
	val := LockReport{}

	i := 2

	val.LockState = payload[i]
	i++

	return val
}