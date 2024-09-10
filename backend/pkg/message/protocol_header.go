package message

import (
	"github.com/oyal2/LIFXMaster/pkg/message/payload"
)

type ProtocolHeader struct {
	_    uint64              // 64 bits: Bytes 24-31
	Type payload.PayloadType // 16 bits: Bytes 32-33
	_    uint16              // 16 bits: Bytes 34-35
}

func (ph *ProtocolHeader) SetMessageType(messageType payload.PayloadType) {
	ph.Type = messageType
}
