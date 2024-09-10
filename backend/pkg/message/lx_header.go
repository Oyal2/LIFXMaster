package message

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"math/rand"
	"time"
)

type LXHeader struct {
	FrameHeader    FrameHeader
	FrameAddress   FrameAddress
	ProtocolHeader ProtocolHeader
}

const (
	HeaderSize uint16 = 36
	Protocol   uint16 = 1024
)

type HeaderError string

func (e HeaderError) Error() string {
	return string(e)
}

const (
	ErrHeaderIncomplete HeaderError = "Insufficient data to unpack a Header"
)

func NewLXHeader() *LXHeader {
	lxph := &LXHeader{
		FrameHeader:    FrameHeader{},
		FrameAddress:   FrameAddress{},
		ProtocolHeader: ProtocolHeader{},
	}
	lxph.FrameHeader.SetSize(HeaderSize)
	lxph.FrameHeader.SetProtocol(Protocol)
	lxph.FrameHeader.SetSource(generateRandomSourceID())
	return lxph
}

func (lxph *LXHeader) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, lxph); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (lxph *LXHeader) Decode(data []byte) error {
	if len(data) < int(HeaderSize) {
		return ErrHeaderIncomplete
	}

	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, lxph); err != nil {
		return err
	}

	return nil
}

func (lxph *LXHeader) JSON() ([]byte, error) {
	return json.Marshal(lxph)
}

func generateRandomSourceID() uint32 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Uint32()
}
