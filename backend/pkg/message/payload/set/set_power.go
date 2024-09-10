package set

import (
	"bytes"
	"encoding/binary"
)

type SetPower struct {
	Level uint16
}

// NewSetPower initializes a new SetPower with default values.
func NewSetPower(level uint16) SetPower {
	return SetPower{
		Level: level,
	}
}

func (sp *SetPower) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := binary.Write(buf, binary.LittleEndian, sp)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (sp *SetPower) Decode(data []byte) error {
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, sp); err != nil {
		return err
	}

	return nil
}
