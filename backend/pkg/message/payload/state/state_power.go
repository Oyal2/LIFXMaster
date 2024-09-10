package state

import (
	"bytes"
	"encoding/binary"
)

type Power struct {
	Level uint16 `json:"level"`
}

func (p *Power) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := binary.Write(buf, binary.LittleEndian, p)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (p *Power) Decode(data []byte) error {
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, p); err != nil {
		return err
	}

	return nil
}
