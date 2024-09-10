package state

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"time"
)

type Group struct {
	Group     [16]byte  `json:"group"`      //The unique identifier of this group as a uuid.
	Label     string    `json:"label"`      //The name assigned to this group
	UpdatedAt time.Time `json:"updated_at"` //An epoch in nanoseconds of when this group was set on the device
}

func (g *Group) Encode() ([]byte, error) {
	buf := make([]byte, 56)

	copy(buf[:16], g.Group[:])

	labelBytes := make([]byte, 32)
	copy(labelBytes, g.Label)
	copy(buf[16:48], labelBytes)

	binary.LittleEndian.PutUint64(buf[48:56], uint64(g.UpdatedAt.UnixNano()))

	return buf, nil
}

func (g *Group) Decode(data []byte) error {
	if len(data) < 56 {
		return fmt.Errorf("data too short, expected at least 56 bytes, got %d", len(data))
	}
	copy(g.Group[:], data[:16])

	g.Label = string(bytes.TrimRight(data[16:48], "\x00"))

	g.UpdatedAt = time.Unix(0, int64(binary.LittleEndian.Uint64(data[48:56])))

	return nil
}
