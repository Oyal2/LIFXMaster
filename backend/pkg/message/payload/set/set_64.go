package set

import (
	"bytes"
	"encoding/binary"

	"github.com/oyal2/LIFXMaster/pkg/message/payload/structure"
)

/*
*

	This lets you set up to 64 HSBK values on the device.

	This message has no response packet even if you set res_required=1.

	This packet requires the device has the Matrix Zones capability. You may use GetVersion (32), GetHostFirmware (14) and the Product Registry to determine whether your device has this capability

*
*/
type Set64 struct {
	TileIndex byte // The device to change. This is 0 indexed and starts from the device closest to the controller.
	Length    byte // The number of devices in the chain to change starting from tile_index
	reserved6 [1]uint8
	X         byte               // The x co-ordinate to start applying colors from. You likely want this to be 0.
	Y         byte               // The y co-ordinate to start applying colors from. You likely want this to be 0.
	Width     byte               // The width of the square you're applying colors to. This should be 8 for the LIFX Tile and 5 for the LIFX Candle.
	Duration  uint32             // The time it will take to transition to new state in milliseconds.
	Colors    [64]structure.HSBK // The HSBK values to assign to each zone specified by this packet.
}

// NewSet64  initializes a new Set64  with default values.
func NewSet64(titleIndex, length, x, y, width byte, durationMilliseconds uint32, colors [64]structure.HSBK) Set64 {
	return Set64{
		TileIndex: titleIndex,
		Length:    length,
		reserved6: [1]uint8{},
		X:         x,
		Y:         y,
		Width:     width,
		Duration:  durationMilliseconds,
		Colors:    colors,
	}
}

func (scz *Set64) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := binary.Write(buf, binary.LittleEndian, scz)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (scz *Set64) Decode(data []byte) error {
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, scz); err != nil {
		return err
	}

	return nil
}
