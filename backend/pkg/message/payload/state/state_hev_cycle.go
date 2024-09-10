package state

import (
	"bytes"
	"encoding/binary"
)

type HEVCycle struct {
	DurationS  uint32 `json:"duration_s"`  // The duration, in seconds, this cycle was set to.
	RemainingS uint32 `json:"remaining_s"` // The duration, in seconds, remaining in this cycle
	LastPower  uint32 `json:"last_power"`  // The power state before the HEV cycle started, which will be the power state once the cycle completes. This is only relevant if remaining_s is larger than 0.
}

func (hc *HEVCycle) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := binary.Write(buf, binary.LittleEndian, hc)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (hc *HEVCycle) Decode(data []byte) error {
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, hc); err != nil {
		return err
	}

	return nil
}
