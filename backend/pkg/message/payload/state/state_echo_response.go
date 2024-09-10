package state

import (
	"bytes"
	"encoding/binary"
)

type EchoResponse struct {
	Echoing [64]byte `json:"echoing"`
}

func (er *EchoResponse) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := binary.Write(buf, binary.LittleEndian, er)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (er *EchoResponse) Decode(data []byte) error {
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, er); err != nil {
		return err
	}

	return nil
}
