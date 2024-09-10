package message

type FrameAddress struct {
	Target   uint64  // 64 bits: Bytes 8-15
	_        [6]byte // 48 bits: Bytes 16-21, treated as reserved
	Control  byte    // 8 bits: Byte 22, containing flags and reserved bits
	Sequence uint8   // 8 bits: Byte 23, simple uint8 value
}

const (
	resRequiredMask = 1 << 0 // Bit 0
	ackRequiredMask = 1 << 1 // Bit 1
)

// SetResRequired sets the res_required flag
func (fa *FrameAddress) SetResRequired(flag bool) {
	if flag {
		fa.Control |= resRequiredMask
	} else {
		fa.Control &^= resRequiredMask
	}
}

// IsResRequired returns true if the res_required flag is set
func (fa *FrameAddress) IsResRequired() bool {
	return fa.Control&resRequiredMask != 0
}

// SetAckRequired sets the ack_required flag
func (fa *FrameAddress) SetAckRequired(flag bool) {
	if flag {
		fa.Control |= ackRequiredMask
	} else {
		fa.Control &^= ackRequiredMask
	}
}

// IsAckRequired returns true if the ack_required flag is set
func (fa *FrameAddress) IsAckRequired() bool {
	return fa.Control&ackRequiredMask != 0
}

// SetTarget sets the Target
func (fa *FrameAddress) SetTarget(Target uint64) {
	fa.Target = Target
}

// SetSequence sets the Sequence
func (fa *FrameAddress) SetSequence(Sequence uint8) {
	fa.Sequence = Sequence
}
