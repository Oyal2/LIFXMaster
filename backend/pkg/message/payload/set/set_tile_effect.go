package set

import (
	"bytes"
	"encoding/binary"

	"github.com/oyal2/LIFXMaster/pkg/message/payload/enum"
	"github.com/oyal2/LIFXMaster/pkg/message/payload/structure"
)

/*
*

	This lets you set up to 64 HSBK values on the device.

	This message has no response packet even if you set res_required=1.

	This packet requires the device has the Matrix Zones capability. You may use GetVersion (32), GetHostFirmware (14) and the Product Registry to determine whether your device has this capability

*
*/
type SetTileEffect struct {
	reserved8    [1]uint8
	reserved9    [1]uint8
	InstanceID   uint32              // A unique number identifying this effect.
	Type         enum.TileEffectType // using TileEffectType Enum
	Speed        uint32              // The time it takes for one cycle of the effect in milliseconds.
	Duration     uint64              // The time the effect will run for in nanoseconds.
	reserved6    [4]uint8
	reserved7    [4]uint8
	Parameters   [8]uint32          // This field is 8 4 byte fields and is currently ignored by all firmware effects.
	PaletteCount byte               // The number of values in palette that you want to use.
	Palette      [16]structure.HSBK // The HSBK values to be used by the effect. Currently only the MORPH effect uses these values.
}

// NewSetTileEffect  initializes a new SetTileEffect  with default values.
func NewSetTileEffect(instanceID uint32, tileEffect enum.TileEffectType, speedMilliseconds uint32, durationNanoseconds uint64, parameters [8]uint32, paletteCount byte, palette [16]structure.HSBK) SetTileEffect {
	return SetTileEffect{
		reserved8:    [1]uint8{},
		reserved9:    [1]uint8{},
		InstanceID:   instanceID,
		Type:         0,
		Speed:        speedMilliseconds,
		Duration:     durationNanoseconds,
		reserved6:    [4]uint8{},
		reserved7:    [4]uint8{},
		Parameters:   parameters,
		PaletteCount: paletteCount,
		Palette:      palette,
	}
}

func (ste *SetTileEffect) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := binary.Write(buf, binary.LittleEndian, ste)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (ste *SetTileEffect) Decode(data []byte) error {
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, ste); err != nil {
		return err
	}

	return nil
}
