package set

import (
	"bytes"
	"encoding/binary"

	"github.com/oyal2/LIFXMaster/pkg/message/payload/structure"
)

type SetColor struct {
	_ byte // 1 reserved byte
	structure.HSBK
	Duration uint32 // Duration for the color transition in milliseconds
}

// NewSetColor initializes a new SetColor with default values.
func NewSetColor(hue, saturation, brightness, kelvin uint16, milliseconds uint32) SetColor {
	return SetColor{
		HSBK: structure.HSBK{
			Hue:        hue,
			Saturation: saturation,
			Brightness: brightness,
			Kelvin:     kelvin,
		},
		Duration: milliseconds,
	}
}

func (scp *SetColor) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := binary.Write(buf, binary.LittleEndian, scp)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (scp *SetColor) Decode(data []byte) error {
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, scp); err != nil {
		return err
	}

	return nil
}
