package message

import (
	"bytes"
	"encoding/binary"
	"encoding/json"

	"github.com/oyal2/LIFXMaster/pkg/message/payload"
)

type LXPacket struct {
	Header  *LXHeader
	Payload payload.PayloadCoder
}

func NewLXPacket(opts ...PacketOptions) *LXPacket {
	packetOptions := &PacketOption{}
	for _, opt := range opts {
		opt(packetOptions)
	}

	header := packetOptions.Header
	if header == nil {
		header = NewLXHeader()
	}

	lxPacket := &LXPacket{
		Header: header,
	}
	lxPacket.Header.FrameHeader.SetTagged(true)
	lxPacket.Header.FrameHeader.SetAddressable(true)

	if packetOptions.AckRequired {
		lxPacket.Header.FrameAddress.SetAckRequired(true)
	}

	if packetOptions.ResponseRequired {
		lxPacket.Header.FrameAddress.SetResRequired(true)
	}

	if packetOptions.PayloadType != 0 {
		lxPacket.Header.ProtocolHeader.SetMessageType(packetOptions.PayloadType)
	}

	if packetOptions.Payload != nil {
		lxPacket.Payload = packetOptions.Payload
	}

	return lxPacket
}

func (lxp *LXPacket) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	var payloadBytes []byte
	var err error
	if lxp.Payload != nil {
		payloadBytes, err = lxp.Payload.Encode()
		if err != nil {
			return nil, err
		}
	}

	lxp.Header.FrameHeader.SetSize(HeaderSize + uint16(len(payloadBytes)))
	headerBytes, err := lxp.Header.Encode()
	if err != nil {
		return nil, err
	}

	err = binary.Write(buf, binary.LittleEndian, headerBytes)
	if err != nil {
		return nil, err
	}

	err = binary.Write(buf, binary.LittleEndian, payloadBytes)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (lxp *LXPacket) Decode(data []byte) error {
	if len(data) < int(HeaderSize) {
		return ErrHeaderIncomplete
	}

	if err := lxp.Header.Decode(data[:36]); err != nil {
		return err
	}

	if len(data) > int(HeaderSize) {
		payload, err := payload.HandleResponsePayload(payload.ResponseType(lxp.Header.ProtocolHeader.Type), data[36:])
		if err != nil {
			return err
		}
		lxp.Payload = payload
	}

	return nil
}
func (lxp *LXPacket) JSON() ([]byte, error) {
	return json.Marshal(lxp)
}
