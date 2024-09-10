package set

import (
	"bytes"
	"encoding/binary"

	"github.com/oyal2/LIFXMaster/pkg/message/payload/enum"
)

/*
*

	Set a segment of your strip to a HSBK value. If your devices supports extended multizone messages it is recommended you use those messages instead.

	Will return one StateMultiZone (506) message

	This packet requires the device has the Linear Zones capability. You may use GetVersion (32), GetHostFirmware (14) and the Product Registry to determine whether your device has this capability

*
*/
type SetColorZones struct {
	StartIndex uint8 // The first zone in the segment we are changing.
	EndIndex   uint8 // The last zone in the segment we are changing
	Hue        uint16
	Saturation uint16
	Brightness uint16
	Kelvin     uint16
	Duration   uint32                           // The amount of time it takes to transition to the new values in milliseconds.
	Apply      enum.MultiZoneApplicationRequest // using MultiZoneApplicationRequest Enum
}

// NewSetColorZones  initializes a new SetColorZones  with default values.
func NewSetColorZones(startIndex, endIndex uint8, hue, saturation, brightness, kelvin uint16, milliseconds uint32, multiZoneApplicationRequest enum.MultiZoneApplicationRequest) SetColorZones {
	return SetColorZones{
		StartIndex: startIndex,
		EndIndex:   endIndex,
		Hue:        hue,
		Saturation: saturation,
		Brightness: brightness,
		Kelvin:     kelvin,
		Duration:   milliseconds,
		Apply:      multiZoneApplicationRequest,
	}
}

func (scz *SetColorZones) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := binary.Write(buf, binary.LittleEndian, scz)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (scz *SetColorZones) Decode(data []byte) error {
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, scz); err != nil {
		return err
	}

	return nil
}
