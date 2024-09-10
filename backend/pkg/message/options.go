package message

import "github.com/oyal2/LIFXMaster/pkg/message/payload"

// PacketOption holds configuration options for a packet
type PacketOption struct {
	Header           *LXHeader
	Payload          payload.PayloadCoder
	AckRequired      bool
	ResponseRequired bool
	Target           uint64
	PayloadType      payload.PayloadType
}

// PacketOptions is a type for function that modifies PacketOption
type PacketOptions func(*PacketOption)

// WithAckRequired sets the AckRequired field of PacketOption
func WithAckRequired(ack bool) PacketOptions {
	return func(po *PacketOption) {
		po.AckRequired = ack
	}
}

// WithResponseRequired sets the ResponseRequired field of PacketOption
func WithResponseRequired(resp bool) PacketOptions {
	return func(po *PacketOption) {
		po.ResponseRequired = resp
	}
}

// WithTarget sets the Target field of PacketOption
func WithTarget(target uint64) PacketOptions {
	return func(po *PacketOption) {
		po.Target = target
	}
}

// WithPayloadType sets the PayloadType field of PacketOption
func WithPayloadType(pt payload.PayloadType) PacketOptions {
	return func(po *PacketOption) {
		po.PayloadType = pt
	}
}

// WithPayload sets the Payload field of PacketOption
func WithPayload(p payload.PayloadCoder) PacketOptions {
	return func(po *PacketOption) {
		po.Payload = p
	}
}

// WithPayload sets the Payload field of PacketOption
func WithHeader(h *LXHeader) PacketOptions {
	return func(po *PacketOption) {
		po.Header = h
	}
}
