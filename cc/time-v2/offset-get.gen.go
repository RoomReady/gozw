// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package timev2

import (
	"encoding/gob"

	"github.com/rmacster/gozw/cc"
)

const CommandOffsetGet cc.CommandID = 0x06

func init() {
	gob.Register(OffsetGet{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x8A),
		Command:      cc.CommandID(0x06),
		Version:      2,
	}, NewOffsetGet)
}

func NewOffsetGet() cc.Command {
	return &OffsetGet{}
}

// <no value>
type OffsetGet struct {
}

func (cmd OffsetGet) CommandClassID() cc.CommandClassID {
	return 0x8A
}

func (cmd OffsetGet) CommandID() cc.CommandID {
	return CommandOffsetGet
}

func (cmd OffsetGet) CommandIDString() string {
	return "TIME_OFFSET_GET"
}

func (cmd *OffsetGet) UnmarshalBinary(data []byte) error {
	// According to the docs, we must copy data if we wish to retain it after returning

	return nil
}

func (cmd *OffsetGet) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	return
}
