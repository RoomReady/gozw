// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package security

import "errors"

// <no value>

type SecurityMessageEncapsulationNonceGet struct {
	InitializationVectorByte []byte

	Properties1 struct {
		SequenceCounter byte

		Sequenced bool

		SecondFrame bool
	}

	CommandClassIdentifier byte

	CommandIdentifier byte

	CommandByte []byte

	ReceiversNonceIdentifier byte

	MessageAuthenticationCodeByte []byte
}

func (cmd *SecurityMessageEncapsulationNonceGet) UnmarshalBinary(payload []byte) error {
	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.InitializationVectorByte = payload[i : i+8]

	i += 8

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.Properties1.SequenceCounter = (payload[i] & 0x0F)

	if payload[i]&0x10 == 0x10 {
		cmd.Properties1.Sequenced = true
	} else {
		cmd.Properties1.Sequenced = false
	}

	if payload[i]&0x20 == 0x20 {
		cmd.Properties1.SecondFrame = true
	} else {
		cmd.Properties1.SecondFrame = false
	}

	i += 1

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.CommandClassIdentifier = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.CommandIdentifier = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	val.CommandByte = payload[i:]

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ReceiversNonceIdentifier = payload[i]
	i++

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.MessageAuthenticationCodeByte = payload[i : i+8]

	i += 8

	return nil
}