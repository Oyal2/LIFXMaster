package set

import (
	"bytes"
	"encoding/binary"
)

/*
*

	This packet lets you set default values for a HEV cycle on the device

	Will return one StateHevCycleConfiguration (147) message

	This packet requires the device has the hev capability. You may use GetVersion (32), GetHostFirmware (14) and the Product Registry to determine whether your device has this capability

*
*/
type SetHevCycleConfiguration struct {
	Indication byte   // Set this to true to run a short flashing indication at the end of the HEV cycle
	DurationS  uint32 // This is the default duration that is used when SetHevCycle (143) is given 0 for duration_s.
}

// NewSetHevCycleConfiguration  initializes a new SetHevCycleConfiguration  with default values.
func NewSetHevCycleConfiguration(indication bool, seconds uint32) SetHevCycleConfiguration {
	var byteBool byte = 0
	if indication {
		byteBool = 1
	}
	return SetHevCycleConfiguration{
		Indication: byteBool,
		DurationS:  seconds,
	}
}

func (shc *SetHevCycleConfiguration) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := binary.Write(buf, binary.LittleEndian, shc)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (shc *SetHevCycleConfiguration) Decode(data []byte) error {
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, shc); err != nil {
		return err
	}

	return nil
}
