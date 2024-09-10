package state

import (
	"bytes"
	"encoding/binary"

	"github.com/oyal2/LIFXMaster/pkg/message/payload/enum"
)

type MultiZoneEffect struct {
	InstanceID uint32                   `json:"instanceid"` // The total number of zones on the strip.
	Type       enum.MultiZoneEffectType `json:"type"`       // The total number of zones on the strip.
	_          [2]byte
	Speed      uint32 `json:"speed"`    // The time it takes for one cycle of the effect in milliseconds
	Duration   uint64 `json:"duration"` // The amount of time left in the current effect in nanoseconds
	_          [4]byte
	_          [4]byte
	Parameters [32]byte `json:"parameters"` // The parameters that was used in the request.
}

func (mze *MultiZoneEffect) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := binary.Write(buf, binary.LittleEndian, mze)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (mze *MultiZoneEffect) Decode(data []byte) error {
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, mze); err != nil {
		return err
	}

	return nil
}
