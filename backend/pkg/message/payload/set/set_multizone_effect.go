package set

import (
	"bytes"
	"encoding/binary"

	"github.com/oyal2/LIFXMaster/pkg/message/payload/enum"
)

/*
*

	Start a multizone Firmware Effect on the device.

	Will return one StateMultiZoneEffect (509) message

	This packet requires the device has the Linear Zones capability. You may use GetVersion (32), GetHostFirmware (14) and the Product Registry to determine whether your device has this capability

*
*/
type SetMultiZoneEffect struct {
	Instanceid uint32                   // A unique number identifying this effect.
	Type       enum.MultiZoneEffectType // using MultiZoneEffectType Enum
	reserved6  [2]byte
	Speed      uint32 // The time it takes for one cycle of the effect in milliseconds.
	Duration   uint64 // The time the effect will run for in nanoseconds.
	reserved7  [4]byte
	reserved8  [4]byte
	parameter  [8]uint32 // This field is 8 4 byte fields which change meaning based on the effect that is running. When the effect is MOVE only the second field is used and is a Uint32 representing the DIRECTION enum. This field is currently ignored for all other multizone effects.
}

// NewSetMultiZoneEffect  initializes a new SetMultiZoneEffect with default values.
func NewSetMultiZoneEffect(instanceID uint32, multiZoneEffectType enum.MultiZoneEffectType, speedMilliSeconds uint32, durationNanoSeconds uint64, parameter [8]uint32) SetMultiZoneEffect {
	return SetMultiZoneEffect{
		Instanceid: instanceID,
		Type:       multiZoneEffectType,
		reserved6:  [2]byte{},
		Speed:      speedMilliSeconds,
		Duration:   durationNanoSeconds,
		reserved7:  [4]byte{},
		reserved8:  [4]byte{},
		parameter:  parameter,
	}
}

func (smze *SetMultiZoneEffect) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := binary.Write(buf, binary.LittleEndian, smze)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (smze *SetMultiZoneEffect) Decode(data []byte) error {
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, smze); err != nil {
		return err
	}

	return nil
}
