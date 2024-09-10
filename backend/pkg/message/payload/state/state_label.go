package state

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Label struct {
	Label string `json:"Label"`
}

func (l *Label) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := binary.Write(buf, binary.LittleEndian, l)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (l *Label) Decode(data []byte) error {
	if len(data) < 32 {
		return fmt.Errorf("data too short, expected at least 32 bytes, got %d", len(data))
	}
	l.Label = string(bytes.Trim(data[:32], "\x00"))

	return nil
}
