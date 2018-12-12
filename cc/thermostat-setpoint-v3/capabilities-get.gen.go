// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package thermostatsetpointv3

import (
	"encoding/gob"
	"errors"

	"github.com/rmacster/gozw/cc"
)

const CommandCapabilitiesGet cc.CommandID = 0x09

func init() {
	gob.Register(CapabilitiesGet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x43),
		Command:      cc.CommandID(0x09),
		Version:      3,
	}, NewCapabilitiesGet)
}

func NewCapabilitiesGet() cc.Command {
	return &CapabilitiesGet{}
}

// <no value>
type CapabilitiesGet struct {
	Properties1 struct {
		SetpointType byte
	}
}

func (cmd CapabilitiesGet) CommandClassID() cc.CommandClassID {
	return 0x43
}

func (cmd CapabilitiesGet) CommandID() cc.CommandID {
	return CommandCapabilitiesGet
}

func (cmd CapabilitiesGet) CommandIDString() string {
	return "THERMOSTAT_SETPOINT_CAPABILITIES_GET"
}

func (cmd *CapabilitiesGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	payload := make([]byte, len(data))
	copy(payload, data)

	if len(payload) < 2 {
		return errors.New("Payload length underflow")
	}

	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.SetpointType = (payload[i] & 0x0F)

	i += 1

	return nil
}

func (cmd *CapabilitiesGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	{
		var val byte

		val |= (cmd.Properties1.SetpointType) & byte(0x0F)

		payload = append(payload, val)
	}

	return
}
