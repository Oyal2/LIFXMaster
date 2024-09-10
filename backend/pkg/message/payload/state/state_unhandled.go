package state

import (
	"bytes"
	"encoding/binary"
)

type Unhandled struct {
	UnhandledType uint16 `json:"unhandled_type"` // The type of the packet that was ignored.
}

func (u *Unhandled) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := binary.Write(buf, binary.LittleEndian, u)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (u *Unhandled) Decode(data []byte) error {
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, u); err != nil {
		return err
	}

	return nil
}
