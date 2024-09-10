package state

import (
	"bytes"
	"encoding/binary"

	"github.com/oyal2/LIFXMaster/pkg/message/payload/structure"
)

type MultiZone struct {
	ZoneCount byte              `json:"zones_count"` // The total number of zones on the strip.
	ZoneIndex byte              `json:"zone_index"`  // The total number of zones on the strip.
	Colors    [8]structure.HSBK `json:"colors"`
}

func (mz *MultiZone) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := binary.Write(buf, binary.LittleEndian, mz)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (mz *MultiZone) Decode(data []byte) error {
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, mz); err != nil {
		return err
	}

	return nil
}
