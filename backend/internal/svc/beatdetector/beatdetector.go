package beatdetector

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"math"
	"math/cmplx"

	"github.com/oyal2/LIFXMaster/pkg/connection"
	"github.com/oyal2/LIFXMaster/pkg/message/payload/set"
	"gonum.org/v1/gonum/dsp/fourier"
	"gonum.org/v1/plot/plotter"
)

const (
	sampleRate = 48000
	maxFreq    = 10000
)

func Run(ctx context.Context, c *connection.LIFXClient, targets ...uint64) {
	RunAudio(ctx, c, targets...)
}

func calculateSpectrum(buffer []float64, c *connection.LIFXClient, targets ...uint64) {
	bufferSize := len(buffer)
	hannWindow := make([]float64, bufferSize)
	for i := range hannWindow {
		hannWindow[i] = 0.5 * (1 - math.Cos(2*math.Pi*float64(i)/float64(bufferSize-1)))
	}
	complexData := make([]float64, len(buffer))
	for i, v := range buffer {
		complexData[i] = v * hannWindow[i]
	}
	fft := fourier.NewFFT(bufferSize)

	coeffs := fft.Coefficients(nil, complexData)
	spectrum := make(plotter.XYs, len(coeffs))

	// Find dominant frequency
	maxMagnitude := 0.0
	dominantFreqBin := 0
	for i := range spectrum {
		freq := (float64(i) * sampleRate / float64(bufferSize)) * 2
		if freq > maxFreq {
			spectrum = spectrum[:i]
			break
		}
		magnitude := cmplx.Abs(coeffs[i])

		db := 20 * math.Log10(magnitude)
		spectrum[i].X = freq
		spectrum[i].Y = db
		if magnitude > maxMagnitude {
			maxMagnitude = magnitude
			dominantFreqBin = i
		}
	}

	dominantFreq := float64(dominantFreqBin) * sampleRate / float64(len(spectrum)) * 2

	// Map frequency to hue
	hue := mapRange(dominantFreq, 0, maxFreq, 0, 65535)

	brightness := uint16(mapRange(calculateBrightness(buffer), 0, 1, 0, 65535))

	go c.SetColor(context.Background(), set.NewSetColor(uint16(hue), math.MaxUint16, brightness, 6500, 250), targets...)
}

func mapRange(value, inMin, inMax, outMin, outMax float64) float64 {
	return (value-inMin)*(outMax-outMin)/(inMax-inMin) + outMin
}

func calculateBrightness(buffer []float64) float64 {
	// Calculate RMS (Root Mean Square) amplitude for brightness
	sumSquares := 0.0
	maxAmplitude := 0.0
	for _, v := range buffer {
		sumSquares += v * v
		if math.Abs(v) > maxAmplitude {
			maxAmplitude = math.Abs(v)
		}
	}
	rmsAmplitude := math.Sqrt(sumSquares / float64(len(buffer)))

	// Normalize RMS amplitude
	normalizedAmplitude := rmsAmplitude / maxAmplitude

	// Apply threshold
	threshold := 0.1 // Adjustable
	if normalizedAmplitude < threshold {
		normalizedAmplitude = 0
	}

	// Compress dynamic range
	compressedAmplitude := math.Log1p(normalizedAmplitude) / math.Log1p(1.0)

	steepness := 8.0 // Adjustable
	midpoint := 0.5  // Adjustable
	return 1 / (1 + math.Exp(-steepness*(compressedAmplitude-midpoint)))
}

func processAudioBuffer(buffer []byte) ([]float64, error) {
	// Convert the byte array to float64 values
	floatData, err := byteArrayToFloat64(buffer, 32) // 32-bit audio
	if err != nil {
		return nil, fmt.Errorf("error converting byte array to float64: %v", err)
	}

	return floatData, nil
}

func byteArrayToFloat64(byteArray []byte, bitDepth int) ([]float64, error) {
	bytesPerSample := bitDepth / 8
	if len(byteArray)%bytesPerSample != 0 {
		return nil, fmt.Errorf("byte array length must be a multiple of %d for %d-bit audio", bytesPerSample, bitDepth)
	}

	sampleCount := len(byteArray) / bytesPerSample
	floatArray := make([]float64, sampleCount)

	reader := bytes.NewReader(byteArray)

	for i := 0; i < sampleCount; i++ {
		var sample int64
		switch bitDepth {
		case 8:
			var s uint8
			err := binary.Read(reader, binary.LittleEndian, &s)
			if err != nil {
				return nil, err
			}
			sample = int64(s) - 128 // convert unsigned to signed
		case 16:
			var s int16
			err := binary.Read(reader, binary.LittleEndian, &s)
			if err != nil {
				return nil, err
			}
			sample = int64(s)
		case 24:
			var s [3]byte
			err := binary.Read(reader, binary.LittleEndian, &s)
			if err != nil {
				return nil, err
			}
			sample = int64(int32(s[0]) | int32(s[1])<<8 | int32(s[2])<<16)
			if sample&0x800000 != 0 {
				sample |= ^0xffffff // sign extend
			}
		case 32:
			var s int32
			err := binary.Read(reader, binary.LittleEndian, &s)
			if err != nil {
				return nil, err
			}
			sample = int64(s)
		default:
			return nil, fmt.Errorf("unsupported bit depth: %d", bitDepth)
		}

		// Normalize to [-1.0, 1.0]
		floatArray[i] = float64(sample) / float64(int(1)<<(bitDepth-1))
	}

	return floatArray, nil
}
