package state

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/oyal2/LIFXMaster/pkg/message/payload/structure"
)

type Light struct {
	structure.HSBK `json:"hsbk"`
	reserved       [2]byte
	Power          uint16 `json:"power"`
	Label          string `json:"label"`
	reserved2      [8]byte
}

func (l *Light) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := binary.Write(buf, binary.LittleEndian, l)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (l *Light) Decode(data []byte) error {
	reader := bytes.NewReader(data)

	// Decode HSBK directly
	if err := binary.Read(reader, binary.LittleEndian, &l.HSBK); err != nil {
		return fmt.Errorf("error reading HSBK: %w", err)
	}

	// Read and discard the reserved bytes
	if err := binary.Read(reader, binary.LittleEndian, &l.reserved); err != nil {
		return fmt.Errorf("error reading reserved bytes: %w", err)
	}

	// Decode Power
	if err := binary.Read(reader, binary.LittleEndian, &l.Power); err != nil {
		return fmt.Errorf("error reading Power: %w", err)
	}

	// Read Label (assuming fixed size for simplicity, adjust as necessary)
	label := make([]byte, 32) // Adjust size according to your actual needs
	if _, err := reader.Read(label); err != nil {
		return fmt.Errorf("error reading Label: %w", err)
	}
	l.Label = string(bytes.Trim(label, "\x00"))

	// Read and discard the reserved2 bytes
	if err := binary.Read(reader, binary.LittleEndian, &l.reserved2); err != nil {
		return fmt.Errorf("error reading reserved2 bytes: %w", err)
	}

	return nil
}
