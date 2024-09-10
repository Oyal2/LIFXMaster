package state

import (
	"bytes"
	"encoding/binary"

	"github.com/oyal2/LIFXMaster/pkg/message/payload/structure"
)

type Tile64 struct {
	TileIndex byte `json:"tile_index"` // The index of the device in the chain this packet refers to. This is 0 based starting from the device closest to the controller.
	_         byte
	X         byte               `json:"x"`      // The x coordinate the colors start from
	Y         byte               `json:"y"`      // The y coordinate the colors start from
	Width     byte               `json:"width"`  // The width of each row
	Colors    [64]structure.HSBK `json:"colors"` // 64 Color structures
}

func (t *Tile64) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := binary.Write(buf, binary.LittleEndian, t)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (t *Tile64) Decode(data []byte) error {
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, t); err != nil {
		return err
	}

	return nil
}
