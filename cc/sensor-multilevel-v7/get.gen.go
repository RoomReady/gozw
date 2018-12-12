// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package sensormultilevelv7

import (
	"encoding/gob"
	"errors"

	"github.com/rmacster/gozw/cc"
)

const CommandGet cc.CommandID = 0x04

func init() {
	gob.Register(Get{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x31),
		Command:      cc.CommandID(0x04),
		Version:      7,
	}, NewGet)
}

func NewGet() cc.Command {
	return &Get{}
}

// <no value>
type Get struct {
	SensorType byte

	Properties1 struct {
		Scale byte
	}
}

func (cmd Get) CommandClassID() cc.CommandClassID {
	return 0x31
}

func (cmd Get) CommandID() cc.CommandID {
	return CommandGet
}

func (cmd Get) CommandIDString() string {
	return "SENSOR_MULTILEVEL_GET"
}

func (cmd *Get) UnmarshalBinary(data []byte) error {
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

	cmd.SensorType = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.Scale = (payload[i] & 0x18) >> 3

	i += 1

	return nil
}

func (cmd *Get) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.SensorType)

	{
		var val byte

		val |= (cmd.Properties1.Scale << byte(3)) & byte(0x18)

		payload = append(payload, val)
	}

	return
}
