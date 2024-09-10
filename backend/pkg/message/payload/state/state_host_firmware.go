package state

import (
	"bytes"
	"encoding/binary"
)

type HostFirmware struct {
	Build        uint64 `json:"build"`
	_            uint64
	VersionMinor uint16 `json:"version_minor"`
	VersionMajor uint16 `json:"version_major"`
}

func (hf *HostFirmware) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := binary.Write(buf, binary.LittleEndian, hf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (hf *HostFirmware) Decode(data []byte) error {
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, hf); err != nil {
		return err
	}

	return nil
}
