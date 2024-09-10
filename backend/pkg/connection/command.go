package connection

import (
	"context"

	"github.com/oyal2/LIFXMaster/pkg/message"
	"github.com/oyal2/LIFXMaster/pkg/message/payload"
	"github.com/oyal2/LIFXMaster/pkg/message/payload/set"
	"github.com/oyal2/LIFXMaster/pkg/message/payload/state"
)

func (c *LIFXClient) SetColor(ctx context.Context, newColor set.SetColor, targets ...uint64) error {
	if len(targets) > 0 && len(targets) != len(c.Device) {
		for _, target := range targets {
			device := c.Device[target]
			m := message.NewLXPacket(message.WithPayloadType(payload.SetColor), message.WithTarget(target))
			m.Payload = &newColor
			msg, err := m.Encode()
			if err != nil {
				return err
			}

			networkPacketArr, err := SendAndReceive(ctx, msg, WithIP(device.Addr), WithPort(int(device.Port)))
			if err != nil {
				return err
			}

			for _, networkPacket := range networkPacketArr {
				newMessage := message.NewLXPacket()
				err = newMessage.Decode(networkPacket.Data)
				if err != nil {
					return err
				}
				if light, ok := newMessage.Payload.(*state.Light); ok {
					c.Lock()
					device.Light = light
					c.Unlock()
				}
			}
		}
	} else {
		m := message.NewLXPacket(message.WithPayloadType(payload.SetColor))
		m.Payload = &newColor
		msg, err := m.Encode()
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
			if light, ok := newMessage.Payload.(*state.Light); ok {
				c.Lock()
				c.Device[newMessage.Header.FrameAddress.Target].Light = light
				c.Unlock()
			}
		}
	}

	return nil
}

func (c *LIFXClient) SetPower(ctx context.Context, setLightPower set.SetLightPower, targets ...uint64) error {
	if len(targets) > 0 && len(targets) != len(c.Device) {
		for _, target := range targets {
			device := c.Device[target]
			m := message.NewLXPacket(message.WithPayloadType(payload.SetLightPower), message.WithTarget(target), message.WithAckRequired(true))
			m.Payload = &setLightPower
			msg, err := m.Encode()
			if err != nil {
				return err
			}

			networkPacketArr, err := SendAndReceive(ctx, msg, WithIP(device.Addr), WithPort(int(device.Port)))
			if err != nil {
				return err
			}

			for _, networkPacket := range networkPacketArr {
				newMessage := message.NewLXPacket()
				err = newMessage.Decode(networkPacket.Data)
				if err != nil {
					return err
				}
				if light, ok := newMessage.Payload.(*state.Light); ok {
					c.Lock()
					device.Light = light
					c.Unlock()
				}
			}
		}
	} else {
		m := message.NewLXPacket(message.WithPayloadType(payload.SetLightPower), message.WithAckRequired(true))
		m.Payload = &setLightPower
		msg, err := m.Encode()
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
			if light, ok := newMessage.Payload.(*state.Light); ok {
				c.Lock()
				c.Device[newMessage.Header.FrameAddress.Target].Light = light
				c.Unlock()
			}
		}
	}

	return nil
}

func (c *LIFXClient) SetLocation(ctx context.Context, setLocation set.SetLocation) error {
	m := message.NewLXPacket(message.WithPayloadType(payload.SetLocation), message.WithAckRequired(true))
	m.Payload = &setLocation
	msg, err := m.Encode()
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
		if location, ok := newMessage.Payload.(*state.Location); ok {
			c.Lock()
			c.Device[newMessage.Header.FrameAddress.Target].Location = location
			c.Unlock()
		}
	}
	return nil
}

func (c *LIFXClient) SetGroup(ctx context.Context, setGroup set.SetGroup) error {
	m := message.NewLXPacket(message.WithPayloadType(payload.SetGroup), message.WithAckRequired(true))
	m.Payload = &setGroup
	msg, err := m.Encode()
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
		if group, ok := newMessage.Payload.(*state.Group); ok {
			c.Lock()
			c.Device[newMessage.Header.FrameAddress.Target].Group = group
			c.Unlock()
		}
	}
	return nil
}

func (c *LIFXClient) SetLabel(ctx context.Context, setLabel set.SetLabel, target uint64) error {
	device := c.Device[target]
	m := message.NewLXPacket(message.WithPayloadType(payload.SetLabel), message.WithTarget(target), message.WithAckRequired(true))
	m.Payload = &setLabel
	msg, err := m.Encode()
	if err != nil {
		return err
	}

	networkPacketArr, err := SendAndReceive(ctx, msg, WithIP(device.Addr), WithPort(int(device.Port)))
	if err != nil {
		return err
	}

	for _, networkPacket := range networkPacketArr {
		newMessage := message.NewLXPacket()
		err = newMessage.Decode(networkPacket.Data)
		if err != nil {
			return err
		}
		if label, ok := newMessage.Payload.(*state.Label); ok {
			c.Lock()
			device.Label = label
			c.Unlock()
		}
	}
	return nil
}

func (c *LIFXClient) SetWaveform(ctx context.Context, setWaveform set.SetWaveform, targets ...uint64) error {
	for _, target := range targets {
		device := c.Device[target]
		m := message.NewLXPacket(message.WithPayloadType(payload.SetWaveform), message.WithTarget(target))
		m.Payload = &setWaveform
		msg, err := m.Encode()
		if err != nil {
			return err
		}

		err = SendAndForget(ctx, msg, WithIP(device.Addr), WithPort(int(device.Port)))
		if err != nil {
			return err
		}
	}
	return nil
}
