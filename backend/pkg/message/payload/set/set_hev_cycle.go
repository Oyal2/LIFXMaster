package set

import (
	"bytes"
	"encoding/binary"
)

/*
*

	This packet lets you start or stop a HEV cycle on the device.

	Will return one StateHevCycle (144) message

	This packet requires the device has the hev capability. You may use GetVersion (32), GetHostFirmware (14) and the Product Registry to determine whether your device has this capability

*
*/
type SetHevCycle struct {
	Enable    byte   // Set this to false to turn off the cycle and true to start the cycle
	DurationS uint32 // The duration, in seconds that the cycle should last for. A value of 0 will use the default duration set by SetHevCycleConfiguration (146).
}

// NewSetHevCycle  initializes a new SetHevCycle  with default values.
func NewSetHevCycle(enable bool, seconds uint32) SetHevCycle {
	var byteBool byte = 0
	if enable {
		byteBool = 1
	}
	return SetHevCycle{
		Enable:    byteBool,
		DurationS: seconds,
	}
}

func (shc *SetHevCycle) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := binary.Write(buf, binary.LittleEndian, shc)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (shc *SetHevCycle) Decode(data []byte) error {
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, shc); err != nil {
		return err
	}

	return nil
}
