// THIS FILE IS AUTO-GENERATED BY CCGEN
// DO NOT MODIFY

package manufacturerspecificv2

import (
	"encoding/binary"
	"errors"
)

// <no value>

type ManufacturerSpecificReport struct {
	ManufacturerId uint16

	ProductTypeId uint16

	ProductId uint16
}

func (cmd *ManufacturerSpecificReport) UnmarshalBinary(payload []byte) error {
	i := 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ManufacturerId = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ProductTypeId = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	if len(payload) <= i {
		return errors.New("slice index out of bounds")
	}

	cmd.ProductId = binary.BigEndian.Uint16(payload[i : i+2])
	i += 2

	return nil
}
