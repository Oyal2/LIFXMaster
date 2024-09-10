package set

import (
	"bytes"
	"encoding/binary"
)

/*
*

	Allows you to specify the position of this device in the chain relative to other device in the chain.

	You can find more information about this data by looking at Tile Positions.

	This message has no response packet even if you set res_required=1.

	This packet requires the device has the Matrix Zones capability. You may use GetVersion (32), GetHostFirmware (14) and the Product Registry to determine whether your device has this capability

*
*/
type SetUserPosition struct {
	TileIndex byte // The device to change. This is 0 indexed and starts from the device closest to the controller.
	reserved6 [2]uint8
	UserX     float32
	UserY     float32
}

// NewSetUserPosition  initializes a new SetUserPosition  with default values.
func NewSetUserPosition(titleIndex byte, userX, userY float32) SetUserPosition {
	return SetUserPosition{
		TileIndex: titleIndex,
		reserved6: [2]uint8{},
		UserX:     userX,
		UserY:     userY,
	}
}

func (scz *SetUserPosition) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := binary.Write(buf, binary.LittleEndian, scz)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (scz *SetUserPosition) Decode(data []byte) error {
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, scz); err != nil {
		return err
	}

	return nil
}
