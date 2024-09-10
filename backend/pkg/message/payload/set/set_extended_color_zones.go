package set

import (
	"bytes"
	"encoding/binary"

	"github.com/oyal2/LIFXMaster/pkg/message/payload/enum"
	"github.com/oyal2/LIFXMaster/pkg/message/payload/structure"
)

/*
*

	This message lets you change the HSBK values for all zones on the strip in one message.

	Will return one StateExtendedColorZones (512) message

	This packet requires the device has the Extended Linear Zones capability. You may use GetVersion (32), GetHostFirmware (14) and the Product Registry to determine whether your device has this capability

*
*/
type SetExtendedColorZones struct {
	Duration    uint32                                   // The time it takes to transition to the new values in milliseconds.
	Apply       enum.MultiZoneExtendedApplicationRequest // using MultiZoneExtendedApplicationRequest Enum Whether to make this change now
	ZoneIndex   uint16                                   // The first zone to apply colors from. If the light has more than 82 zones, then send multiple messages with different indices to update the whole device.
	ColorsCount uint16                                   // The number of colors in the colors field to apply to the strip
	Colors      [82]structure.HSBK                       // The HSBK values to change the strip with.
}

// NewSetExtendedColorZones  initializes a new SetExtendedColorZones  with default values.
func NewSetExtendedColorZones(durationMillisecond uint32, multiZoneExtendedApplicationRequest enum.MultiZoneExtendedApplicationRequest, zoneIndex, ColorsCount uint16, colors [82]structure.HSBK) SetExtendedColorZones {
	return SetExtendedColorZones{
		Duration:    durationMillisecond,
		Apply:       multiZoneExtendedApplicationRequest,
		ZoneIndex:   zoneIndex,
		ColorsCount: ColorsCount,
		Colors:      colors,
	}
}

func (secz *SetExtendedColorZones) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := binary.Write(buf, binary.LittleEndian, secz)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (secz *SetExtendedColorZones) Decode(data []byte) error {
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, secz); err != nil {
		return err
	}

	return nil
}
