package state

import (
	"bytes"
	"encoding/binary"
)

type Infrared struct {
	Brightness uint16 `json:"Brightness"`
}

func (i *Infrared) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := binary.Write(buf, binary.LittleEndian, i)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (i *Infrared) Decode(data []byte) error {
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, i); err != nil {
		return err
	}

	return nil
}
