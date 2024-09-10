package svc

import (
	"math"

	pb "github.com/oyal2/LIFXMaster/internal/proto"
	"github.com/oyal2/LIFXMaster/pkg/device"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func deviceToProto(device *device.LightDevice) *pb.Device {
	if device == nil {
		return nil
	}
	pbDevice := &pb.Device{
		Address: device.Addr.String(),
		Port:    device.Port,
		Target:  device.Target,
		Firmware: &pb.Firmware{
			Build:        device.Firmware.Build,
			VersionMinor: uint32(device.Firmware.VersionMinor),
			VersionMajor: uint32(device.Firmware.VersionMajor),
		},
		Wifi: &pb.WiFi{
			Info: &pb.WifiInfo{
				Signal: device.Wifi.Info.Signal,
			},
			Firmware: &pb.Firmware{Build: device.Wifi.Firmware.Build, VersionMinor: uint32(device.Wifi.Firmware.VersionMinor), VersionMajor: uint32(device.Wifi.Firmware.VersionMajor)},
		},
		Power: &pb.Power{
			Level: uint32(device.Power.Level),
		},
		Label: &pb.Label{
			Label: device.Label.Label,
		},
		Version: &pb.Version{
			Vendor:  device.Version.Vendor,
			Product: device.Version.Product,
		},
		Info: &pb.Info{
			Time:     device.Info.Time,
			Uptime:   device.Info.Uptime,
			Downtime: device.Info.Downtime,
		},
		Product: &pb.Product{
			Pid:  uint32(device.Product.PID),
			Name: device.Product.Name,
			Features: &pb.Features{
				Hev:                        device.Product.Features.Hev,
				Color:                      device.Product.Features.Color,
				Chain:                      device.Product.Features.Chain,
				Matrix:                     device.Product.Features.Matrix,
				Relays:                     device.Product.Features.Relays,
				Buttons:                    device.Product.Features.Buttons,
				Infrared:                   device.Product.Features.Infrared,
				Multizone:                  device.Product.Features.Multizone,
				TemperatureRange:           []int32{},
				ExtendedMultizone:          device.Product.Features.ExtendedMultizone,
				MinExtMzFirmware:           &wrapperspb.Int32Value{},
				MinExtMzFirmwareComponents: []int32{},
			},
			Upgrades: &pb.Upgrade{
				Major: 0,
				Minor: 0,
				Features: &pb.Features{
					Hev:                        device.Product.Features.Hev,
					Color:                      device.Product.Features.Color,
					Chain:                      device.Product.Features.Chain,
					Matrix:                     device.Product.Features.Matrix,
					Relays:                     device.Product.Features.Relays,
					Buttons:                    device.Product.Features.Buttons,
					Infrared:                   device.Product.Features.Infrared,
					Multizone:                  device.Product.Features.Multizone,
					TemperatureRange:           []int32{},
					ExtendedMultizone:          device.Product.Features.ExtendedMultizone,
					MinExtMzFirmware:           &wrapperspb.Int32Value{},
					MinExtMzFirmwareComponents: []int32{},
				},
			},
		},
		Light: &pb.Light{
			Hue:        uint32(device.Light.Hue),
			Saturation: uint32(device.Light.Saturation),
			Brightness: uint32(device.Light.Brightness),
			Kelvin:     uint32(device.Light.Kelvin),
			Power:      uint32(device.Light.Power),
			Label:      device.Light.Label,
		},
		Infrared: &pb.Infrared{
			Brightness: uint32(device.Infrared.Brightness),
		},
		Hev: &pb.HEV{
			Cycle: &pb.Cycle{
				DurationS:  device.HEV.Cycle.DurationS,
				RemainingS: device.HEV.Cycle.RemainingS,
				LastPower:  device.HEV.Cycle.LastPower,
			},
			Config: &pb.Config{
				Indication: device.HEV.Config.Indication,
				DurationS:  device.HEV.Config.DurationS,
			},
			LastCycleResult: &pb.LastCycleResult{
				Result: uint32(device.HEV.LastCycleResult.Result),
			},
		},
		ExtendedColorZone: &pb.ExtendedColorZone{
			ZonesCount:  uint32(device.ExtendedColorZone.ZoneCount),
			ZoneIndex:   uint32(device.ExtendedColorZone.ZoneIndex),
			ColorsCount: uint32(device.ExtendedColorZone.ColorsCount),
			Colors:      []*pb.HSBK{},
		},
		Relay: &pb.Relay{
			RPower: &pb.RPower{
				RelayIndex: uint32(device.Relay.RPower.RelayIndex),
				Level:      uint32(device.Relay.RPower.Level),
			},
		},
		Tile: &pb.Tile{
			DeviceChain: &pb.DeviceChain{
				StartIndex:       uint32(device.Tile.DeviceChain.StartIndex),
				TileDevices:      []*pb.TileDevices{},
				TileDevicesCount: uint32(device.Tile.DeviceChain.TileDevicesCount),
			},
			Tile_64: &pb.Tile64{
				TileIndex: uint32(device.Tile.Tile64.TileIndex),
				X:         uint32(device.Tile.Tile64.X),
				Y:         uint32(device.Tile.Tile64.Y),
				Width:     uint32(device.Tile.Tile64.Width),
				Colors:    []*pb.HSBK{},
			},
			TileEffect: &pb.TileEffect{
				Instanceid:   device.Tile.TileEffect.InsteanceID,
				Type:         uint32(device.Tile.TileEffect.Type),
				Speed:        device.Tile.TileEffect.Speed,
				Duration:     device.Tile.TileEffect.Duration,
				Parameters:   []uint32{},
				PaletteCount: uint32(device.Tile.TileEffect.PaletteCount),
				Palette:      []*pb.HSBK{},
			},
			SensorAmbientLight: &pb.SensorAmbientLight{
				Lux: []uint32{},
			},
		},
		Group: &pb.Group{
			Label:     device.Group.Label,
			UpdatedAt: device.Group.UpdatedAt.String(),
		},
		Location: &pb.Location{
			Label:     device.Location.Label,
			UpdatedAt: device.Location.UpdatedAt.String(),
		},
	}

	return pbDevice
}

func generateHues(steps int) []uint16 {
	colors := make([]uint16, steps)
	hueIncrement := math.MaxUint16 / steps

	for i := 0; i < steps; i++ {
		colors[i] = uint16((i * hueIncrement) % math.MaxUint16)
	}

	return colors
}
