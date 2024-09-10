package state

import (
	"bytes"
	"encoding/binary"

	"github.com/oyal2/LIFXMaster/pkg/message/payload/structure"
)

type Zone struct {
	ZoneCount      byte `json:"zones_count"` // The total number of zones on the strip.
	ZoneIndex      byte `json:"zone_index"`  // The total number of zones on the strip.
	structure.HSBK `json:"colors"`
}

func (z *Zone) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := binary.Write(buf, binary.LittleEndian, z)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (z *Zone) Decode(data []byte) error {
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, z); err != nil {
		return err
	}

	return nil
}
