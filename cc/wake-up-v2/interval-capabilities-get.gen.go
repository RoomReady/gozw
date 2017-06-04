// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package wakeupv2

import (
	"encoding/gob"

	"github.com/gozwave/gozw/cc"
)

const CommandIntervalCapabilitiesGet cc.CommandID = 0x09

func init() {
	gob.Register(IntervalCapabilitiesGet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x84),
		Command:      cc.CommandID(0x09),
		Version:      2,
	}, NewIntervalCapabilitiesGet)
}

func NewIntervalCapabilitiesGet() cc.Command {
	return &IntervalCapabilitiesGet{}
}

// <no value>
type IntervalCapabilitiesGet struct {
}

func (cmd IntervalCapabilitiesGet) CommandClassID() cc.CommandClassID {
	return 0x84
}

func (cmd IntervalCapabilitiesGet) CommandID() cc.CommandID {
	return CommandIntervalCapabilitiesGet
}

func (cmd IntervalCapabilitiesGet) CommandIDString() string {
	return "WAKE_UP_INTERVAL_CAPABILITIES_GET"
}

func (cmd *IntervalCapabilitiesGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	return nil
}

func (cmd *IntervalCapabilitiesGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	return
}