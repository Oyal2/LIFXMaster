package set

import (
	"bytes"
	"encoding/binary"
)

/**
This packet lets you change the current infrared value on the device

Will return one StateInfrared (121) message

This packet requires the device has the infrared capability. You may use GetVersion (32), GetHostFirmware (14) and the Product Registry to determine whether your device has this capability
*/

type SetInfrared struct {
	Brightness uint16 // The amount of infrared emitted by the device. 0 is no infrared and 65535 is the most infrared.
}

// NewSetInfrared  initializes a new SetInfrared  with default values.
func NewSetInfrared(brightness uint16) SetInfrared {
	return SetInfrared{
		Brightness: brightness,
	}
}

func (si *SetInfrared) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := binary.Write(buf, binary.LittleEndian, si)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (si *SetInfrared) Decode(data []byte) error {
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, si); err != nil {
		return err
	}

	return nil
}
