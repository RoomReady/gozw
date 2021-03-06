// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package configurationv4

import (
	"encoding/binary"
	"encoding/gob"
	"errors"

	"github.com/rmacster/gozw/cc"
)

const CommandNameReport cc.CommandID = 0x0B

func init() {
	gob.Register(NameReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x70),
		Command:      cc.CommandID(0x0B),
		Version:      4,
	}, NewNameReport)
}

func NewNameReport() cc.Command {
	return &NameReport{}
}

// <no value>
type NameReport struct {
	ParameterNumber uint16

	ReportsToFollow byte

	Name []byte
}

func (cmd NameReport) CommandClassID() cc.CommandClassID {
	return 0x70
}

func (cmd NameReport) CommandID() cc.CommandID {
	return CommandNameReport
}

func (cmd NameReport) CommandIDString() string {
	return "CONFIGURATION_NAME_REPORT"
}

func (cmd *NameReport) UnmarshalBinary(data []byte) error {
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

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ReportsToFollow = payload[i]
	i++

	if len(payload) <= i {
		return nil
	}

	cmd.Name = payload[i:]

	return nil
}

func (cmd *NameReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	{
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, cmd.ParameterNumber)
		payload = append(payload, buf...)
	}

	payload = append(payload, cmd.ReportsToFollow)

	payload = append(payload, cmd.Name...)

	return
}
