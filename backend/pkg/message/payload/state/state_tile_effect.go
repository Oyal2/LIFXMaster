package state

import (
	"bytes"
	"encoding/binary"

	"github.com/oyal2/LIFXMaster/pkg/message/payload/enum"
	"github.com/oyal2/LIFXMaster/pkg/message/payload/structure"
)

type TileEffect struct {
	_            byte
	InsteanceID  uint32              `json:"instanceid"` // The unique value identifying the request
	Type         enum.TileEffectType `json:"type"`       // Uint8 using TileEffectType Enum
	Speed        uint32              `json:"speed"`      // The time it takes for one cycle in milliseconds.
	Duration     uint64              `json:"duration"`   // The width of each row
	_            [4]byte
	_            [4]byte
	Parameters   [32]byte           `json:"parameters"`    // The parameters as specified in the request.
	PaletteCount byte               `json:"palette_count"` // The number of colors in the palette that are relevant
	Palette      [16]structure.HSBK `json:"palette"`       // The colors specified for the effect.
}

func (te *TileEffect) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := binary.Write(buf, binary.LittleEndian, te)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (te *TileEffect) Decode(data []byte) error {
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, te); err != nil {
		return err
	}

	return nil
}
