package set

import (
	"bytes"
	"encoding/binary"

	"github.com/oyal2/LIFXMaster/pkg/message/payload/enum"
)

/*
*

		This behaves like SetWaveform (103) but allows you to keep certain parts of the original HSBK values during the transition.

	 	Will return one LightState (107) message

*
*/
type SetWaveformOptional struct {
	Reserved      [1]byte       // Reserved byte, not used but must be read
	Transient     bool          // Indicates whether the light state is temporary
	Hue           uint16        // Color hue
	Saturation    uint16        // Color saturation
	Brightness    uint16        // Light brightness
	Kelvin        uint16        // Color temperature in Kelvin
	Period        uint32        // Duration of one cycle of the waveform in milliseconds
	Cycles        float32       // Number of cycles the waveform is performed
	SkewRatio     int16         // Skew ratio of the waveform, typically between -32768 and 32767
	Waveform      enum.Waveform // Waveform to be used for light modulation
	SetHue        bool          // Flag to apply the hue value
	SetSaturation bool          // Flag to apply the saturation value
	SetBrightness bool          // Flag to apply the brightness value
	SetKelvin     bool          // Flag to apply the kelvin value
}

// NewSetWaveformOptional initializes a new SetWaveformOptional with default values.
func NewSetWaveformOptional(transient bool, hue uint16, saturation uint16, brightness uint16, kelvin uint16, milliseconds uint32, cycles float32, skewRatio int16, waveform enum.Waveform, setHue bool, setSaturation bool, setBrightness bool, setKelvin bool) SetWaveformOptional {
	return SetWaveformOptional{
		Reserved:      [1]byte{},
		Transient:     transient,
		Hue:           hue,
		Saturation:    saturation,
		Brightness:    brightness,
		Kelvin:        kelvin,
		Period:        milliseconds,
		Cycles:        cycles,
		SkewRatio:     skewRatio,
		Waveform:      waveform,
		SetHue:        setHue,
		SetSaturation: setSaturation,
		SetBrightness: setBrightness,
		SetKelvin:     setKelvin,
	}
}

func (swfo *SetWaveformOptional) Encode() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := binary.Write(buf, binary.LittleEndian, swfo)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (swfo *SetWaveformOptional) Decode(data []byte) error {
	if err := binary.Read(bytes.NewReader(data), binary.LittleEndian, swfo); err != nil {
		return err
	}

	return nil
}
