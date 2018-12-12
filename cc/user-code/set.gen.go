// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package usercode

import (
	"encoding/gob"
	"errors"

	"github.com/rmacster/gozw/cc"
)

const CommandSet cc.CommandID = 0x01

func init() {
	gob.Register(Set{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x63),
		Command:      cc.CommandID(0x01),
		Version:      1,
	}, NewSet)
}

func NewSet() cc.Command {
	return &Set{}
}

// <no value>
type Set struct {
	UserIdentifier byte

	UserIdStatus byte

	UserCode string
}

func (cmd Set) CommandClassID() cc.CommandClassID {
	return 0x63
}

func (cmd Set) CommandID() cc.CommandID {
	return CommandSet
}

func (cmd Set) CommandIDString() string {
	return "USER_CODE_SET"
}

func (cmd *Set) UnmarshalBinary(data []byte) error {
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

	cmd.UserIdentifier = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.UserIdStatus = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.UserCode = string(payload[i : i+10])

	i += 10

	return nil
}

func (cmd *Set) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.UserIdentifier)

	payload = append(payload, cmd.UserIdStatus)

	if paramLen := len(cmd.UserCode); paramLen > 10 {
		return nil, errors.New("Length overflow in array parameter UserCode")
	}

	payload = append(payload, []byte(cmd.UserCode)...)

	return
}
