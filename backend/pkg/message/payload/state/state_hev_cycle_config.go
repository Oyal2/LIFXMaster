package state

import (
	"bytes"
	"encoding/binary"
)

type HEVCycleConfig struct {
	Indication bool   `json:"indication"` // Whether a short flashing indication is run at the end of an HEV cycle.
	DurationS  uint32 `json:"duration_s"` // This is the default duration that is used when SetHevCycle (143) is given 0 for duration_s.
}

func (hcc *HEVCycleConfig) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := binary.Write(buf, binary.LittleEndian, hcc)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (hcc *HEVCycleConfig) Decode(data []byte) error {
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, hcc); err != nil {
		return err
	}

	return nil
}
