package message

type FrameHeader struct {
	Size    uint16 // 16 bits: Total Size of the message
	Control uint16 // 12 bits: This will contain Protocol (12 bits), Addressable (1 bit), Tagged (1 bit), and Origin (2 bits)
	Source  uint32 // 32 bits: Identifier for the Source
}

// Define constants for bit manipulation masks based on the specification
const (
	protocolMask    = 0xFFF
	addressableMask = 1 << 12
	taggedMask      = 1 << 13
	originMask      = 3 << 14
)

// SetProtocol safely sets the protocol field ensuring only the lower 12 bits are used.
func (fh *FrameHeader) SetProtocol(protocol uint16) {
	fh.Control = (fh.Control &^ protocolMask) | (protocol & protocolMask)
}

// Protocol returns the current protocol value.
func (fh *FrameHeader) Protocol() uint16 {
	return fh.Control & protocolMask
}

// SetAddressable sets the addressable bit.
func (fh *FrameHeader) SetAddressable(value bool) {
	if value {
		fh.Control |= addressableMask
	} else {
		fh.Control &^= addressableMask
	}
}

// IsAddressable checks the state of the addressable bit.
func (fh *FrameHeader) Addressable() bool {
	return fh.Control&addressableMask != 0
}

// SetTagged sets the tagged bit.
func (fh *FrameHeader) SetTagged(value bool) {
	if value {
		fh.Control |= taggedMask
	} else {
		fh.Control &^= taggedMask
	}
}

// IsTagged checks the state of the tagged bit.
func (fh *FrameHeader) Tagged() bool {
	return fh.Control&taggedMask != 0
}

// SetOrigin sets the 2-bit origin field
func (fh *FrameHeader) SetOrigin(origin uint8) {
	// Clear the current origin bits then set the new value
	fh.Control = (fh.Control &^ originMask) | (uint16(origin) << 14)
}

// Origin retrieves the 2-bit origin field from the Control field.
func (fh *FrameHeader) Origin() uint8 {
	return uint8((fh.Control & originMask) >> 14)
}

func (fh *FrameHeader) SetSize(Size uint16) {
	fh.Size = Size
}

// Source returns the current Source value.
func (fh *FrameHeader) SetSource(Source uint32) {
	fh.Source = Source
}
