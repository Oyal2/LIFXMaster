package set

import (
	"bytes"
	"encoding/binary"
)

/*
*

	This is the same as SetPower (21) but allows you to specify how long it will take to transition to the new power state.

	Will return one StateLightPower (118) message

*
*/
type SetLightPower struct {
	Level    uint16 // If you specify 0 the light will turn off and if you specify 65535 the device will turn on.
	Duration uint32 // The time it will take to transition to the new state in milliseconds.
}

// NewSetLightPower initializes a new SetLightPower with default values.
func NewSetLightPower(Level uint16, milliseconds uint32) SetLightPower {
	return SetLightPower{
		Level:    Level,
		Duration: milliseconds,
	}
}

func (slp *SetLightPower) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := binary.Write(buf, binary.LittleEndian, slp)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (slp *SetLightPower) Decode(data []byte) error {
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, slp); err != nil {
		return err
	}

	return nil
}
