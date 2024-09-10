package state

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"time"
)

type Location struct {
	Location  [16]byte  `json:"location"`   //The unique identifier of this group as a uuid.
	Label     string    `json:"label"`      //The name assigned to this location
	UpdatedAt time.Time `json:"updated_at"` //An epoch in nanoseconds of when this location was set on the device
}

func (l *Location) Encode() ([]byte, error) {
	buf := make([]byte, 56)

	copy(buf[:16], l.Location[:])

	labelBytes := make([]byte, 32)
	copy(labelBytes, l.Label)
	copy(buf[16:48], labelBytes)

	binary.LittleEndian.PutUint64(buf[48:56], uint64(l.UpdatedAt.UnixNano()))

	return buf, nil
}

func (l *Location) Decode(data []byte) error {
	if len(data) < 56 {
		return fmt.Errorf("data too short, expected at least 56 bytes, got %d", len(data))
	}
	copy(l.Location[:], data[:16])

	l.Label = string(bytes.TrimRight(data[16:48], "\x00"))

	l.UpdatedAt = time.Unix(0, int64(binary.LittleEndian.Uint64(data[48:56])))

	return nil
}
