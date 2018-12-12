// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package switchbinary

import (
	"encoding/gob"
	"errors"

	"github.com/rmacster/gozw/cc"
)

const CommandReport cc.CommandID = 0x03

func init() {
	gob.Register(Report{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x25),
		Command:      cc.CommandID(0x03),
		Version:      1,
	}, NewReport)
}

func NewReport() cc.Command {
	return &Report{}
}

// <no value>
type Report struct {
	Value byte
}

func (cmd Report) CommandClassID() cc.CommandClassID {
	return 0x25
}

func (cmd Report) CommandID() cc.CommandID {
	return CommandReport
}

func (cmd Report) CommandIDString() string {
	return "SWITCH_BINARY_REPORT"
}

func (cmd *Report) UnmarshalBinary(data []byte) error {
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

	cmd.Value = payload[i]
	i++

	return nil
}

func (cmd *Report) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.Value)

	return
}
