package connection

import (
	"context"
	"net"
	"sync"
	"time"

	"github.com/oyal2/LIFXMaster/pkg/device"
	"github.com/oyal2/LIFXMaster/pkg/message"
	"github.com/oyal2/LIFXMaster/pkg/message/payload"
	"github.com/oyal2/LIFXMaster/pkg/message/payload/state"
)

type LIFXClient struct {
	sync.RWMutex
	BcastAddr   net.IP                         // Host's IP
	Device      map[uint64]*device.LightDevice // All devices that are connected to the network
	ProductInfo *device.ProductInfo            // Information about LIFX Products and its capabilities
}

type NetworkPacket struct {
	Data []byte   // The actual data of the packet
	Addr net.Addr // Network address of the sender or receiver
}

const (
	LXPort   = 56700
	DeadLine = 500 * time.Millisecond
	IP       = "192.168.1.255"
)

func NewLXClient(bcastIP string) (*LIFXClient, error) {
	productInfo, err := device.NewProductInfo()
	if err != nil {
		return nil, err
	}
	return &LIFXClient{
		BcastAddr:   net.ParseIP(bcastIP),
		ProductInfo: productInfo,
		Device:      make(map[uint64]*device.LightDevice),
	}, nil
}

func (c *LIFXClient) GetDevices(ctx context.Context) error {
	c.Lock()
	defer c.Unlock()
	msg, err := message.NewLXPacket(message.WithPayloadType(payload.GetService)).Encode()
	if err != nil {
		return err
	}

	networkPacketArr, err := SendAndReceive(ctx, msg)
	if err != nil {
		return err
	}

	for _, networkPacket := range networkPacketArr {
		newMessage := message.NewLXPacket()
		err = newMessage.Decode(networkPacket.Data)
		if err != nil {
			return err
		}
		if server, ok := newMessage.Payload.(*state.Service); ok {
			c.Device[newMessage.Header.FrameAddress.Target] = device.NewLightBulb(networkPacket.Addr.(*net.UDPAddr).IP, server.Port, newMessage.Header.FrameAddress.Target)
		}
	}

	return nil
}

func (c *LIFXClient) GetInfo(ctx context.Context, target uint64) error {
	c.Lock()
	defer c.Unlock()
	device := c.Device[target]

	if err := SendQuery(ctx, device, payload.GetHostFirmware, hostFirmwareProcessor()); err != nil {
		return err
	}

	if err := SendQuery(ctx, device, payload.GetWifiInfo, wifiInfoProcessor()); err != nil {
		return err
	}

	if err := SendQuery(ctx, device, payload.GetWifiFirmware, wifiFirmwareProcessor()); err != nil {
		return err
	}

	if err := SendQuery(ctx, device, payload.GetPower, powerProcessor()); err != nil {
		return err
	}

	if err := SendQuery(ctx, device, payload.GetLabel, labelProcessor()); err != nil {
		return err
	}

	if err := SendQuery(ctx, device, payload.GetVersion, versionProcessor()); err != nil {
		return err
	}

	if err := SendQuery(ctx, device, payload.GetInfo, infoProcessor()); err != nil {
		return err
	}

	if err := SendQuery(ctx, device, payload.GetLocation, locationProcessor()); err != nil {
		return err
	}

	if err := SendQuery(ctx, device, payload.GetGroup, groupProcessor()); err != nil {
		return err
	}

	if err := SendQuery(ctx, device, payload.GetColor, lightProcessor()); err != nil {
		return err
	}

	device.Product = c.ProductInfo.GetProduct(int(device.Version.Product))

	if device.Product.Features.Infrared {
		if err := SendQuery(ctx, device, payload.GetInfrared, infraredProcessor()); err != nil {
			return err
		}
	}

	if device.Product.Features.Hev {
		if err := SendQuery(ctx, device, payload.GetHevCycle, hevCycleProcessor()); err != nil {
			return err
		}

		if err := SendQuery(ctx, device, payload.GetHevCycleConfiguration, hevCycleConfigProcessor()); err != nil {
			return err
		}

		if err := SendQuery(ctx, device, payload.GetLastHevCycleResult, hevCycleConfigProcessor()); err != nil {
			return err
		}
	}

	//TODO
	if device.Product.Features.Multizone {
		// msg, err = message.NewLXPacket(message.WithPayloadType(payload.GetColorZones), message.WithResponseRequired(true), message.WithTarget(target)).Encode()
		// if err != nil {
		// 	return err
		// }

		// networkPacketArr, err = SendAndReceive(ctx, msg, WithIP(device.Addr), WithPort(int(device.Port)))
		// if err != nil {
		// 	return err
		// }

		// for _, networkPacket := range networkPacketArr {
		// 	newMessage := message.NewLXPacket()
		// 	err = newMessage.Decode(networkPacket.Data)
		// 	if err != nil {
		// 		return err
		// 	}
		// 	if hevCycle, ok := newMessage.Payload.(*state.HEVCycle); ok {
		// 		device.HEV.Cycle = hevCycle
		// 	}
		// }

		if device.Product.Features.ExtendedMultizone {
			if err := SendQuery(ctx, device, payload.GetExtendedColorZones, extendedColorZoneProcessor()); err != nil {
				return err
			}
		}
	}

	if device.Product.Features.Relays {
		if err := SendQuery(ctx, device, payload.GetRPower, relayProcessor()); err != nil {
			return err
		}
	}

	if device.Product.Features.Matrix {
		if err := SendQuery(ctx, device, payload.GetDeviceChain, deviceChainProcessor()); err != nil {
			return err
		}

		if err := SendQuery(ctx, device, payload.Get64, tile64Processor()); err != nil {
			return err
		}

		if err := SendQuery(ctx, device, payload.GetTileEffect, tileEffectProcessor()); err != nil {
			return err
		}

		if err := SendQuery(ctx, device, payload.GetSensorAmbientLight, sensorAmbientLightProcessor()); err != nil {
			return err
		}
	}

	return nil
}

func (c *LIFXClient) GetAllInfo(ctx context.Context) ([]device.LightDevice, error) {
	c.Lock()
	defer c.Unlock()

	if err := SendAllQuery(ctx, c.Device, payload.GetHostFirmware, hostFirmwareProcessor()); err != nil {
		return nil, err
	}

	if err := SendAllQuery(ctx, c.Device, payload.GetWifiInfo, wifiInfoProcessor()); err != nil {
		return nil, err
	}

	if err := SendAllQuery(ctx, c.Device, payload.GetWifiFirmware, wifiFirmwareProcessor()); err != nil {
		return nil, err
	}

	if err := SendAllQuery(ctx, c.Device, payload.GetPower, powerProcessor()); err != nil {
		return nil, err
	}

	if err := SendAllQuery(ctx, c.Device, payload.GetLabel, labelProcessor()); err != nil {
		return nil, err
	}

	if err := SendAllQuery(ctx, c.Device, payload.GetVersion, versionProcessor()); err != nil {
		return nil, err
	}

	if err := SendAllQuery(ctx, c.Device, payload.GetInfo, infoProcessor()); err != nil {
		return nil, err
	}

	if err := SendAllQuery(ctx, c.Device, payload.GetLocation, locationProcessor()); err != nil {
		return nil, err
	}

	if err := SendAllQuery(ctx, c.Device, payload.GetGroup, groupProcessor()); err != nil {
		return nil, err
	}

	if err := SendAllQuery(ctx, c.Device, payload.GetColor, lightProcessor()); err != nil {
		return nil, err
	}

	var cpyDevices []device.LightDevice
	for _, device := range c.Device {
		device.Product = c.ProductInfo.GetProduct(int(device.Version.Product))
		cpyDevices = append(cpyDevices, *device)
	}

	return cpyDevices, nil
}

func SendAndReceive(ctx context.Context, msg []byte, opts ...ConnectionOption) ([]NetworkPacket, error) {
	connOpts := &ConnectionOptions{
		Deadline: DeadLine,
		IP:       net.ParseIP(IP),
		Port:     LXPort,
	}
	for _, opt := range opts {
		opt(connOpts)
	}

	conn, err := net.ListenUDP("udp", &net.UDPAddr{})
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	conn.SetWriteDeadline(time.Now().Add(connOpts.Deadline))

	dst := &net.UDPAddr{
		IP:   connOpts.IP,
		Port: connOpts.Port,
	}

	if _, err = conn.WriteToUDP(msg, dst); err != nil {
		return nil, err
	}

	if err = conn.SetReadDeadline(time.Now().Add(connOpts.Deadline)); err != nil {
		return nil, err
	}

	var networkPackets []NetworkPacket
	for {
		buffer := make([]byte, 512)
		n, addr, err := conn.ReadFrom(buffer)
		if err != nil {
			break
		}
		networkPackets = append(networkPackets, NetworkPacket{
			Data: buffer[:n],
			Addr: addr,
		})
	}
	return networkPackets, nil
}

func SendAndForget(ctx context.Context, msg []byte, opts ...ConnectionOption) error {
	connOpts := &ConnectionOptions{
		Deadline: DeadLine,
		IP:       net.ParseIP(IP),
		Port:     LXPort,
	}
	for _, opt := range opts {
		opt(connOpts)
	}

	conn, err := net.ListenUDP("udp", &net.UDPAddr{})
	if err != nil {
		return err
	}
	defer conn.Close()

	conn.SetWriteDeadline(time.Now().Add(connOpts.Deadline))

	dst := &net.UDPAddr{
		IP:   connOpts.IP,
		Port: connOpts.Port,
	}

	if _, err = conn.WriteToUDP(msg, dst); err != nil {
		return err
	}

	return nil
}

// SendQuery sends a query to a device and processes the response.
func SendQuery(ctx context.Context, device *device.LightDevice, payloadType payload.PayloadType, payloadProcessor PayloadProcessor) error {
	msg, err := message.NewLXPacket(
		message.WithPayloadType(payloadType),
		message.WithResponseRequired(true),
		message.WithTarget(device.Target),
	).Encode()
	if err != nil {
		return err
	}

	networkPacketArr, err := SendAndReceive(ctx, msg, WithIP(device.Addr), WithPort(int(device.Port)))
	if err != nil {
		return err
	}

	for _, networkPacket := range networkPacketArr {
		newMessage := message.NewLXPacket()
		if err := newMessage.Decode(networkPacket.Data); err != nil {
			return err
		}
		if err := payloadProcessor(newMessage, device); err != nil {
			return err
		}
	}
	return nil
}

// SendAllQuery sends a query to all devices and processes the response.
func SendAllQuery(ctx context.Context, devices map[uint64]*device.LightDevice, payloadType payload.PayloadType, payloadProcessor PayloadProcessor) error {
	msg, err := message.NewLXPacket(
		message.WithPayloadType(payloadType),
		message.WithResponseRequired(true),
	).Encode()
	if err != nil {
		return err
	}

	networkPacketArr, err := SendAndReceive(ctx, msg)
	if err != nil {
		return err
	}

	for _, networkPacket := range networkPacketArr {
		newMessage := message.NewLXPacket()
		if err := newMessage.Decode(networkPacket.Data); err != nil {
			return err
		}
		if err := payloadProcessor(newMessage, devices[newMessage.Header.FrameAddress.Target]); err != nil {
			return err
		}
	}
	return nil
}
