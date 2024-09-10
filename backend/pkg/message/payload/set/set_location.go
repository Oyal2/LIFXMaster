package set

import (
	"bytes"
	"encoding/binary"
)

// This packet lets you set the location information on the device.
type SetLocation struct {
	Location  [16]byte // 16 bytes representing a UUID of the location. You should have the same UUID value for each device in this location
	Label     string   // The name of the location.
	UpdatedAt uint64   // The time you updated the location of this device as an epoch in nanoseconds.
}

// NewSetLocation  initializes a new SetLocation  with default values.
func NewSetLocation(location [16]byte, label string, UpdatedAt uint64) SetLocation {
	return SetLocation{
		Location:  location,
		Label:     label,
		UpdatedAt: UpdatedAt,
	}
}

func (sl *SetLocation) Encode() ([]byte, error) {
	buf := make([]byte, 56)

	copy(buf[:16], sl.Location[:])

	labelBytes := []byte(sl.Label)
	if len(labelBytes) > 32 {
		labelBytes = labelBytes[:32]
	} else if len(labelBytes) < 32 {
		labelBytes = append(labelBytes, make([]byte, 32-len(labelBytes))...)
	}
	copy(buf[16:48], labelBytes)

	binary.LittleEndian.PutUint64(buf[48:], sl.UpdatedAt)

	return buf, nil
}

func (sl *SetLocation) Decode(data []byte) error {
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, sl); err != nil {
		return err
	}

	return nil
}
