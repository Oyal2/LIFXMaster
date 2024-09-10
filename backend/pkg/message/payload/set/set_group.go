package set

import (
	"bytes"
	"encoding/binary"
)

// This packet lets you set the group information on the device.
type SetGroup struct {
	Group     [16]byte // 16 bytes representing a UUID of the group. You should have the same UUID value for each device in this group
	Label     string   // The name of the location.
	UpdatedAt uint64   // The time you updated the location of this device as an epoch in nanoseconds.
}

// NewSetGroup  initializes a new SetGroup  with default values.
func NewSetGroup(group [16]byte, label string, UpdatedAt uint64) SetGroup {
	return SetGroup{
		Group:     group,
		Label:     label,
		UpdatedAt: UpdatedAt,
	}
}

func (sg *SetGroup) Encode() ([]byte, error) {
	buf := make([]byte, 56)

	copy(buf[:16], sg.Group[:])

	labelBytes := []byte(sg.Label)
	if len(labelBytes) > 32 {
		labelBytes = labelBytes[:32]
	} else if len(labelBytes) < 32 {
		labelBytes = append(labelBytes, make([]byte, 32-len(labelBytes))...)
	}
	copy(buf[16:48], labelBytes)

	binary.LittleEndian.PutUint64(buf[48:], sg.UpdatedAt)

	return buf, nil
}

func (sg *SetGroup) Decode(data []byte) error {
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, sg); err != nil {
		return err
	}

	return nil
}
