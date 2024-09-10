package state

import (
	"bytes"
	"encoding/binary"

	"github.com/oyal2/LIFXMaster/pkg/message/payload/structure"
)

type ExtendedColorZone struct {
	ZoneCount   byte               `json:"zones_count"`  // The number of zones on your strip
	ZoneIndex   byte               `json:"zone_index"`   // The first zone represented in the packet. If the light has more than 82 zones, then this property indicates the relative positioning of the colors contained in a given message.
	ColorsCount byte               `json:"colors_count"` // The number of HSBK values in the colors array that map to zones.
	Colors      [82]structure.HSBK `json:"colors"`       // The HSBK values currently set on each zone.
}

func (ecz *ExtendedColorZone) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := binary.Write(buf, binary.LittleEndian, ecz)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (ecz *ExtendedColorZone) Decode(data []byte) error {
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, ecz); err != nil {
		return err
	}

	return nil
}
