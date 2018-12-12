// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package configurationv4

import (
	"encoding/binary"
	"encoding/gob"
	"errors"

	"github.com/rmacster/gozw/cc"
)

const CommandNameGet cc.CommandID = 0x0A

func init() {
	gob.Register(NameGet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x70),
		Command:      cc.CommandID(0x0A),
		Version:      4,
	}, NewNameGet)
}

func NewNameGet() cc.Command {
	return &NameGet{}
}

// <no value>
type NameGet struct {
	ParameterNumber uint16
}

func (cmd NameGet) CommandClassID() cc.CommandClassID {
	return 0x70
}

func (cmd NameGet) CommandID() cc.CommandID {
	return CommandNameGet
}

func (cmd NameGet) CommandIDString() string {
	return "CONFIGURATION_NAME_GET"
}

func (cmd *NameGet) UnmarshalBinary(data []byte) error {
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

	cmd.ParameterNumber = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	return nil
}

func (cmd *NameGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	{
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, cmd.ParameterNumber)
		payload = append(payload, buf...)
	}

	return
}
