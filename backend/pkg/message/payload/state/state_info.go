package state

import (
	"bytes"
	"encoding/binary"
)

type Info struct {
	Time     uint64 `json:"time"`
	Uptime   uint64 `json:"uptime"`
	Downtime uint64 `json:"downtime"`
}

func (i *Info) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := binary.Write(buf, binary.LittleEndian, i)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (i *Info) Decode(data []byte) error {
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, i); err != nil {
		return err
	}

	return nil
}
