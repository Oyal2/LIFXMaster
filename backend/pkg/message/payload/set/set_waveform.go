package set

import (
	"bytes"
	"encoding/binary"

	"github.com/oyal2/LIFXMaster/pkg/message/payload/enum"
)

/*
*

	This packet lets you set the HSBK and Waveforms value for the light.
	For devices that have multiple zones, this will treat all the zones on the device as one.
	LightState represents the state of the light, including its current settings and effects.

*
*/
type SetWaveform struct {
	Reserved   [1]byte       // Reserved byte, not used but must be read
	Transient  bool          // Indicates whether the light state is temporary
	Hue        uint16        // Color hue
	Saturation uint16        // Color saturation
	Brightness uint16        // Light brightness
	Kelvin     uint16        // Color temperature in Kelvin
	Period     uint32        // Duration of one cycle of the waveform in milliseconds
	Cycles     float32       // Number of cycles the waveform is performed
	SkewRatio  int16         // Skew ratio of the waveform, typically between -32768 and 32767
	Waveform   enum.Waveform // Waveform to be used for light modulation
}

// NewSetWaveform  initializes a new SetWaveform  with default values.
func NewSetWaveform(transient bool, hue uint16, saturation uint16, brightness uint16, kelvin uint16, milliseconds uint32, cycles float32, skewRatio int16, waveform enum.Waveform) SetWaveform {
	return SetWaveform{
		Reserved:   [1]byte{},
		Transient:  transient,
		Hue:        hue,
		Saturation: saturation,
		Brightness: brightness,
		Kelvin:     kelvin,
		Period:     milliseconds,
		Cycles:     cycles,
		SkewRatio:  skewRatio,
		Waveform:   waveform,
	}
}

func (swf *SetWaveform) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := binary.Write(buf, binary.LittleEndian, swf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (swf *SetWaveform) Decode(data []byte) error {
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, swf); err != nil {
		return err
	}

	return nil
}
