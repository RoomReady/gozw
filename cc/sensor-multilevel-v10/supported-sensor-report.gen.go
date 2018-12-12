// THIS FILE IS AUTO-GENERATED BY ZWGEN
// DO NOT MODIFY

package sensormultilevelv10

import (
	"encoding/gob"
	"errors"

	"github.com/rmacster/gozw/cc"
)

const CommandSupportedSensorReport cc.CommandID = 0x02

func init() {
	gob.Register(SupportedSensorReport{})
	cc.Register(cc.CommandIdentifier{
		CommandClass: cc.CommandClassID(0x31),
		Command:      cc.CommandID(0x02),
		Version:      10,
	}, NewSupportedSensorReport)
}

func NewSupportedSensorReport() cc.Command {
	return &SupportedSensorReport{}
}

// <no value>
type SupportedSensorReport struct {
	BitMask []byte
}

func (cmd SupportedSensorReport) CommandClassID() cc.CommandClassID {
	return 0x31
}

func (cmd SupportedSensorReport) CommandID() cc.CommandID {
	return CommandSupportedSensorReport
}

func (cmd SupportedSensorReport) CommandIDString() string {
	return "SENSOR_MULTILEVEL_SUPPORTED_SENSOR_REPORT"
}

func (cmd *SupportedSensorReport) UnmarshalBinary(data []byte) error {
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

	cmd.BitMask = payload[i:]

	return nil
}

func (cmd *SupportedSensorReport) MarshalBinary() (payload []byte, err error) {
	payload = make([]byte, 2)
	payload[0] = byte(cmd.CommandClassID())
	payload[1] = byte(cmd.CommandID())

	payload = append(payload, cmd.BitMask...)

	return
}
