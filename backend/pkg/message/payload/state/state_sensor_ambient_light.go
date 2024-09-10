package state

import (
	"bytes"
	"encoding/binary"
)

type SensorAmbientLight struct {
	Lux [4]byte `json"float"`
}

func (sal *SensorAmbientLight) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := binary.Write(buf, binary.LittleEndian, sal)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (sal *SensorAmbientLight) Decode(data []byte) error {
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, sal); err != nil {
		return err
	}

	return nil
}
