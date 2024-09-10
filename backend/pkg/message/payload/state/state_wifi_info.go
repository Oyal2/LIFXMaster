package state

import (
	"bytes"
	"encoding/binary"
)

type WifiInfo struct {
	Signal uint32 `json:"signal"`
	_      uint32
	_      uint32
	_      uint16
}

func (wi *WifiInfo) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := binary.Write(buf, binary.LittleEndian, wi)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (wi *WifiInfo) Decode(data []byte) error {
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, wi); err != nil {
		return err
	}

	return nil
}
