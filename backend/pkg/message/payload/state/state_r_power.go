package state

import (
	"bytes"
	"encoding/binary"
)

type RPower struct {
	RelayIndex byte   `json:"relay_index"`
	Level      uint16 `json:"level"`
}

func (rp *RPower) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := binary.Write(buf, binary.LittleEndian, rp)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (rp *RPower) Decode(data []byte) error {
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, rp); err != nil {
		return err
	}

	return nil
}
