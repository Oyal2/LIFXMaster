package connection

import (
	"context"

	"github.com/oyal2/LIFXMaster/pkg/device"
	"github.com/oyal2/LIFXMaster/pkg/message/payload/set"
)

type LIFXController interface {
	GetDevices(ctx context.Context) error
	GetInfo(ctx context.Context, target uint64) error
	GetAllInfo(ctx context.Context) ([]device.DeviceProfiler, error)
	SetColor(ctx context.Context, newColor set.SetColor, targets ...uint64) error
	SetPower(ctx context.Context, setLightPower set.SetLightPower, targets ...uint64) error
	SetWaveform(ctx context.Context, setWaveform set.SetWaveform, targets ...uint64) error
}
