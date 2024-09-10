package set

import (
	"bytes"
	"encoding/binary"
)

type SetLabel struct {
	Label string
}

// NewSetLabel initializes a new SetLabel with default values.
func NewSetLabel(label string) SetLabel {
	return SetLabel{
		Label: label,
	}
}

func (sl *SetLabel) Encode() ([]byte, error) {
	buf := make([]byte, 32)

	labelBytes := []byte(sl.Label)
	if len(labelBytes) > 32 {
		labelBytes = labelBytes[:32]
	} else if len(labelBytes) < 32 {
		labelBytes = append(labelBytes, make([]byte, 32-len(labelBytes))...)
	}
	copy(buf[:32], labelBytes)

	return buf, nil
}

func (sl *SetLabel) Decode(data []byte) error {
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, sl); err != nil {
		return err
	}

	return nil
}
