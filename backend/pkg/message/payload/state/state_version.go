package state

import (
	"bytes"
	"encoding/binary"
)

type Version struct {
	Vendor  uint32 `json:"vendor"`
	Product uint32 `json:"product"`
	_       uint32
}

func (v *Version) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := binary.Write(buf, binary.LittleEndian, v)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (v *Version) Decode(data []byte) error {
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, v); err != nil {
		return err
	}

	return nil
}
