package state

import (
	"bytes"
	"encoding/binary"

	"github.com/oyal2/LIFXMaster/pkg/message/payload/enum"
)

type Service struct {
	Service enum.Services `json:"service"` // Using Services Enum
	Port    uint32        `json:"port"`    // The port of the service. This value is usually 56700 but you should not assume this is always the case.
}

func (s *Service) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := binary.Write(buf, binary.LittleEndian, s)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (s *Service) Decode(data []byte) error {
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, s); err != nil {
		return err
	}

	return nil
}
