package connection

import (
	"errors"

	"github.com/oyal2/LIFXMaster/pkg/device"
	"github.com/oyal2/LIFXMaster/pkg/message"
	"github.com/oyal2/LIFXMaster/pkg/message/payload/state"
)

type PayloadProcessor func(*message.LXPacket, *device.LightDevice) error

func hostFirmwareProcessor() PayloadProcessor {
	return func(packet *message.LXPacket, device *device.LightDevice) error {
		if packet == nil {
			return errors.New("message packet is empty")
		}

		if device == nil {
			return errors.New("device is empty")
		}

		if firmware, ok := packet.Payload.(*state.HostFirmware); ok {
			device.Firmware = firmware
			return nil
		}
		return errors.New("unexpected payload type")
	}
}

func wifiInfoProcessor() PayloadProcessor {
	return func(packet *message.LXPacket, device *device.LightDevice) error {
		if packet == nil {
			return errors.New("message packet is empty")
		}

		if device == nil {
			return errors.New("device is empty")
		}

		if info, ok := packet.Payload.(*state.WifiInfo); ok {
			device.Wifi.Info = info
			return nil
		}
		return errors.New("unexpected payload type")
	}
}

func wifiFirmwareProcessor() PayloadProcessor {
	return func(packet *message.LXPacket, device *device.LightDevice) error {
		if packet == nil {
			return errors.New("message packet is empty")
		}

		if device == nil {
			return errors.New("device is empty")
		}

		if wifiFirmware, ok := packet.Payload.(*state.WifiFirmware); ok {
			device.Wifi.Firmware = wifiFirmware
			return nil
		}
		return errors.New("unexpected payload type")
	}
}

func powerProcessor() PayloadProcessor {
	return func(packet *message.LXPacket, device *device.LightDevice) error {
		if packet == nil {
			return errors.New("message packet is empty")
		}

		if device == nil {
			return errors.New("device is empty")
		}

		if power, ok := packet.Payload.(*state.Power); ok {
			device.Power = power
			return nil
		}
		return errors.New("unexpected payload type")
	}
}

func labelProcessor() PayloadProcessor {
	return func(packet *message.LXPacket, device *device.LightDevice) error {
		if packet == nil {
			return errors.New("message packet is empty")
		}

		if device == nil {
			return errors.New("device is empty")
		}

		if label, ok := packet.Payload.(*state.Label); ok {
			device.Label = label
			return nil
		}
		return errors.New("unexpected payload type")
	}
}

func versionProcessor() PayloadProcessor {
	return func(packet *message.LXPacket, device *device.LightDevice) error {
		if packet == nil {
			return errors.New("message packet is empty")
		}

		if device == nil {
			return errors.New("device is empty")
		}

		if version, ok := packet.Payload.(*state.Version); ok {
			if version.Product == 0 || version.Vendor == 0 {
				return errors.New("product version is undefined")
			}
			device.Version = version
			return nil
		}
		return errors.New("unexpected payload type")
	}
}

func infoProcessor() PayloadProcessor {
	return func(packet *message.LXPacket, device *device.LightDevice) error {
		if packet == nil {
			return errors.New("message packet is empty")
		}

		if device == nil {
			return errors.New("device is empty")
		}

		if info, ok := packet.Payload.(*state.Info); ok {
			device.Info = info
			return nil
		}
		return errors.New("unexpected payload type")
	}
}

func locationProcessor() PayloadProcessor {
	return func(packet *message.LXPacket, device *device.LightDevice) error {
		if packet == nil {
			return errors.New("message packet is empty")
		}

		if device == nil {
			return errors.New("device is empty")
		}

		if location, ok := packet.Payload.(*state.Location); ok {
			device.Location = location
			return nil
		}
		return errors.New("unexpected payload type")
	}
}

func groupProcessor() PayloadProcessor {
	return func(packet *message.LXPacket, device *device.LightDevice) error {
		if packet == nil {
			return errors.New("message packet is empty")
		}

		if device == nil {
			return errors.New("device is empty")
		}

		if group, ok := packet.Payload.(*state.Group); ok {
			device.Group = group
			return nil
		}
		return errors.New("unexpected payload type")
	}
}

func lightProcessor() PayloadProcessor {
	return func(packet *message.LXPacket, device *device.LightDevice) error {
		if packet == nil {
			return errors.New("message packet is empty")
		}

		if device == nil {
			return errors.New("device is empty")
		}

		if light, ok := packet.Payload.(*state.Light); ok {
			device.Light = light
			return nil
		}
		return errors.New("unexpected payload type")
	}
}

func infraredProcessor() PayloadProcessor {
	return func(packet *message.LXPacket, device *device.LightDevice) error {
		if packet == nil {
			return errors.New("message packet is empty")
		}

		if device == nil {
			return errors.New("device is empty")
		}

		if infrared, ok := packet.Payload.(*state.Infrared); ok {
			device.Infrared = infrared
			return nil
		}
		return errors.New("unexpected payload type")
	}
}

func hevCycleProcessor() PayloadProcessor {
	return func(packet *message.LXPacket, device *device.LightDevice) error {
		if packet == nil {
			return errors.New("message packet is empty")
		}

		if device == nil {
			return errors.New("device is empty")
		}

		if hevCycle, ok := packet.Payload.(*state.HEVCycle); ok {
			device.HEV.Cycle = hevCycle
			return nil
		}
		return errors.New("unexpected payload type")
	}
}

func hevCycleConfigProcessor() PayloadProcessor {
	return func(packet *message.LXPacket, device *device.LightDevice) error {
		if packet == nil {
			return errors.New("message packet is empty")
		}

		if device == nil {
			return errors.New("device is empty")
		}

		if hevCycleConfig, ok := packet.Payload.(*state.HEVCycleConfig); ok {
			device.HEV.Config = hevCycleConfig
			return nil
		}
		return errors.New("unexpected payload type")
	}
}

func extendedColorZoneProcessor() PayloadProcessor {
	return func(packet *message.LXPacket, device *device.LightDevice) error {
		if packet == nil {
			return errors.New("message packet is empty")
		}

		if device == nil {
			return errors.New("device is empty")
		}

		if extendedColorZone, ok := packet.Payload.(*state.ExtendedColorZone); ok {
			device.ExtendedColorZone = extendedColorZone
			return nil
		}
		return errors.New("unexpected payload type")
	}
}

func lastHEVCycleResultProcessor() PayloadProcessor {
	return func(packet *message.LXPacket, device *device.LightDevice) error {
		if packet == nil {
			return errors.New("message packet is empty")
		}

		if device == nil {
			return errors.New("device is empty")
		}

		if lastHEVCycleResult, ok := packet.Payload.(*state.LastHEVCycleResult); ok {
			device.HEV.LastCycleResult = lastHEVCycleResult
			return nil
		}
		return errors.New("unexpected payload type")
	}
}

func relayProcessor() PayloadProcessor {
	return func(packet *message.LXPacket, device *device.LightDevice) error {
		if packet == nil {
			return errors.New("message packet is empty")
		}

		if device == nil {
			return errors.New("device is empty")
		}

		if rPower, ok := packet.Payload.(*state.RPower); ok {
			device.Relay.RPower = rPower
			return nil
		}
		return errors.New("unexpected payload type")
	}
}

func deviceChainProcessor() PayloadProcessor {
	return func(packet *message.LXPacket, device *device.LightDevice) error {
		if packet == nil {
			return errors.New("message packet is empty")
		}

		if device == nil {
			return errors.New("device is empty")
		}

		if deviceChain, ok := packet.Payload.(*state.DeviceChain); ok {
			device.Tile.DeviceChain = deviceChain
			return nil
		}
		return errors.New("unexpected payload type")
	}
}

func tile64Processor() PayloadProcessor {
	return func(packet *message.LXPacket, device *device.LightDevice) error {
		if packet == nil {
			return errors.New("message packet is empty")
		}

		if device == nil {
			return errors.New("device is empty")
		}

		if tile64, ok := packet.Payload.(*state.Tile64); ok {
			device.Tile.Tile64 = tile64
			return nil
		}
		return errors.New("unexpected payload type")
	}
}

func tileEffectProcessor() PayloadProcessor {
	return func(packet *message.LXPacket, device *device.LightDevice) error {
		if packet == nil {
			return errors.New("message packet is empty")
		}

		if device == nil {
			return errors.New("device is empty")
		}

		if tileEffect, ok := packet.Payload.(*state.TileEffect); ok {
			device.Tile.TileEffect = tileEffect
			return nil
		}
		return errors.New("unexpected payload type")
	}
}

func sensorAmbientLightProcessor() PayloadProcessor {
	return func(packet *message.LXPacket, device *device.LightDevice) error {
		if packet == nil {
			return errors.New("message packet is empty")
		}

		if device == nil {
			return errors.New("device is empty")
		}

		if sensorAmbientLight, ok := packet.Payload.(*state.SensorAmbientLight); ok {
			device.Tile.SensorAmbientLight = sensorAmbientLight
			return nil
		}
		return errors.New("unexpected payload type")
	}
}
