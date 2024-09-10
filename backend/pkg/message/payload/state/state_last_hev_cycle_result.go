package state

import (
	"bytes"
	"encoding/binary"

	"github.com/oyal2/LIFXMaster/pkg/message/payload/enum"
)

type LastHEVCycleResult struct {
	Result enum.LightLastHevCycleResult `json:"result"` // An enum saying whether the last cycle completed or interrupted.
}

func (lhcr *LastHEVCycleResult) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := binary.Write(buf, binary.LittleEndian, lhcr)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (lhcr *LastHEVCycleResult) Decode(data []byte) error {
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, lhcr); err != nil {
		return err
	}

	return nil
}
