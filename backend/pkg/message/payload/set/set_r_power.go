package set

import (
	"bytes"
	"encoding/binary"
)

/*
*

	Set the power state of a relay on a switch device. Current models of the LIFX switch do not have dimming capability, so the two valid values are 0 for off and 65535 for on.

	# Will return one StateRPower (818) message

	This packet requires the device has the Relays capability. You may use GetVersion (32), GetHostFirmware (14) and the Product Registry to determine whether your device has this capability

*
*/
type SetRPower struct {
	RelayIndex uint8  // The relay on the switch starting from 0.
	Level      uint16 // The new value of the relay
}

// NewSetRPower  initializes a new SetRPower  with default values.
func NewSetRPower(relayIndex uint8, level uint16) SetRPower {
	return SetRPower{
		RelayIndex: relayIndex,
		Level:      level,
	}
}

func (srp *SetRPower) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := binary.Write(buf, binary.LittleEndian, srp)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (srp *SetRPower) Decode(data []byte) error {
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, srp); err != nil {
		return err
	}

	return nil
}
