package state

import (
	"bytes"
	"encoding/binary"
)

type WifiFirmware struct {
	Build        uint64 `json:"build"`
	_            uint64
	VersionMinor uint16 `json:"version_minor"`
	VersionMajor uint16 `json:"version_major"`
}

func (wf *WifiFirmware) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := binary.Write(buf, binary.LittleEndian, wf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (wf *WifiFirmware) Decode(data []byte) error {
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, wf); err != nil {
		return err
	}

	return nil
}
