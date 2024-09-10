package state

import (
	"bytes"
	"encoding/binary"

	"github.com/oyal2/LIFXMaster/pkg/message/payload/structure"
)

type DeviceChain struct {
	StartIndex       byte               `json:"start_index"`        // The index of the first device in the chain this packet refers to
	TileDevices      [16]structure.Tile `json:"tile_devices"`       // The information for each device in the chain
	TileDevicesCount byte               `json:"tile_devices_count"` // The number of device in tile_devices that map to devices in the chain.
}

func (dc *DeviceChain) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := binary.Write(buf, binary.LittleEndian, dc)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (dc *DeviceChain) Decode(data []byte) error {
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, dc); err != nil {
		return err
	}

	return nil
}
