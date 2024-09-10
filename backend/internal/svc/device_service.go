package svc

import (
	"context"
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/google/uuid"
	pb "github.com/oyal2/LIFXMaster/internal/proto"
	"github.com/oyal2/LIFXMaster/internal/svc/beatdetector"
	"github.com/oyal2/LIFXMaster/internal/svc/theater"
	"github.com/oyal2/LIFXMaster/pkg/connection"
	"github.com/oyal2/LIFXMaster/pkg/message/payload/enum"
	"github.com/oyal2/LIFXMaster/pkg/message/payload/set"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DeviceServer struct {
	Client        *connection.LIFXClient
	ActiveEffects *ActiveEffects
	pb.UnimplementedDeviceServiceServer
}

func NewDeviceSvc(client *connection.LIFXClient) pb.DeviceServiceServer {
	return &DeviceServer{
		Client:                           client,
		ActiveEffects:                    NewActiveEffects(),
		UnimplementedDeviceServiceServer: pb.UnimplementedDeviceServiceServer{},
	}
}

func (ds *DeviceServer) GetDevices(ctx context.Context, request *pb.GetDevicesRequest) (resp *pb.GetDevicesResponse, err error) {
	resp = &pb.GetDevicesResponse{
		Locations: make(map[string]*pb.LocationMap),
	}
	if err = ds.Client.GetDevices(ctx); err != nil {
		return resp, err
	}

	devices, err := ds.Client.GetAllInfo(ctx)
	if err != nil {
		return resp, err
	}

	ds.Client.RLock()
	for _, device := range devices {
		pbDevice := deviceToProto(&device)
		if pbDevice != nil {
			locUUID, err := uuid.FromBytes(device.Location.Location[:])
			if err != nil {
				return resp, fmt.Errorf("error converting bytes to uuid: %v", err)
			}
			locationKey := locUUID.String()

			grpUUID, err := uuid.FromBytes(device.Location.Location[:])
			if err != nil {
				return resp, fmt.Errorf("error converting bytes to uuid: %v", err)
			}
			groupKey := grpUUID.String()
			if resp.Locations[locationKey] == nil {
				resp.Locations[locationKey] = &pb.LocationMap{
					Groups:    make(map[string]*pb.GroupMap),
					Label:     device.Location.Label,
					UpdatedAt: device.Location.UpdatedAt.Local().String(),
				}
			}

			locationMap := resp.Locations[locationKey]
			if locationMap.Groups == nil {
				locationMap.Groups = make(map[string]*pb.GroupMap)
			}
			if locationMap.Groups[groupKey] == nil {
				locationMap.Groups[groupKey] = &pb.GroupMap{
					Devices:   make([]*pb.Device, 0),
					Label:     device.Group.Label,
					UpdatedAt: device.Group.UpdatedAt.Local().String(),
				}
			}

			locationMap.Groups[groupKey].Devices = append(locationMap.Groups[groupKey].Devices, pbDevice)
		}
	}
	ds.Client.RUnlock()

	return resp, err
}

func (ds *DeviceServer) SetLocationLabel(ctx context.Context, request *pb.SetLocationLabelRequest) (*pb.SetLocationLabelResponse, error) {
	resp := &pb.SetLocationLabelResponse{}
	uuid, err := uuid.Parse(request.GetLocationID())
	if err != nil {
		return resp, status.Error(codes.InvalidArgument, "unable to parse location id")
	}
	err = ds.Client.SetLocation(ctx, set.NewSetLocation(uuid, request.GetNewLabel(), uint64(time.Now().UnixNano())))
	if err != nil {
		return resp, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (ds *DeviceServer) SetGroupLabel(ctx context.Context, request *pb.SetGroupLabelRequest) (*pb.SetGroupLabelResponse, error) {
	resp := &pb.SetGroupLabelResponse{}
	uuid, err := uuid.Parse(request.GetGroupID())
	if err != nil {
		return resp, status.Error(codes.InvalidArgument, "unable to parse group id")
	}
	err = ds.Client.SetGroup(ctx, set.NewSetGroup(uuid, request.GetNewLabel(), uint64(time.Now().UnixNano())))
	if err != nil {
		return resp, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (ds *DeviceServer) SetDeviceLabel(ctx context.Context, request *pb.SetDeviceLabelRequest) (*pb.SetDeviceLabelResponse, error) {
	resp := &pb.SetDeviceLabelResponse{}
	err := ds.Client.SetLabel(ctx, set.NewSetLabel(request.GetNewLabel()), request.GetDeviceID())
	if err != nil {
		return resp, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (ds *DeviceServer) SetColor(ctx context.Context, request *pb.SetColorRequest) (*pb.SetColorResponse, error) {
	resp := &pb.SetColorResponse{}

	for target, hsbk := range request.GetColors() {
		ds.ActiveEffects.CancelEffect(target)
		go ds.Client.SetColor(ctx, set.NewSetColor(uint16(hsbk.Hue), uint16(hsbk.Saturation), uint16(hsbk.Brightness), uint16(hsbk.Kelvin), 50), target)
	}

	return resp, nil
}

func (ds *DeviceServer) SetPower(ctx context.Context, request *pb.SetPowerRequest) (*pb.SetPowerResponse, error) {
	resp := &pb.SetPowerResponse{}

	for target, power := range request.GetPowers() {
		ds.ActiveEffects.CancelEffect(target)
		var level uint16 = 0
		if power {
			level = 65535
		}
		go ds.Client.SetPower(ctx, set.NewSetLightPower(level, 0), target)
	}

	return resp, nil
}

func (ds *DeviceServer) Strobe(ctx context.Context, request *pb.StrobeRequest) (*pb.StrobeResponse, error) {
	resp := &pb.StrobeResponse{}

	deviceIDs := request.GetDeviceIDs()
	for _, deviceID := range deviceIDs {
		ds.ActiveEffects.CancelEffect(deviceID)
	}

	if request.TurnOn {
		ctxCancel, cancel := context.WithCancel(context.Background())
		for _, deviceID := range deviceIDs {
			ds.ActiveEffects.Add(deviceID, cancel)
		}
		go func(ctx context.Context, interval float32, targets ...uint64) {
			defer cancel()
			d := time.Duration(interval)
			ticker := time.NewTicker((d + 25) * time.Millisecond)
			defer ticker.Stop()
			for {
				newWaveForm := set.NewSetWaveform(true, 0, 0, 0, 0, uint32(interval), 1, 0, enum.PULSE)
				ds.Client.SetWaveform(ctx, newWaveForm, targets...)
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					ds.Client.SetWaveform(ctx, newWaveForm, targets...)
				}
			}
		}(ctxCancel, request.Speed, deviceIDs...)
	}
	return resp, nil
}

func (ds *DeviceServer) ColorCycle(ctx context.Context, request *pb.ColorCycleRequest) (*pb.ColorCycleResponse, error) {
	resp := &pb.ColorCycleResponse{}
	deviceIDs := request.GetDeviceIDs()
	for _, deviceID := range deviceIDs {
		ds.ActiveEffects.CancelEffect(deviceID)
	}

	if request.TurnOn {
		ctxCancel, cancel := context.WithCancel(context.Background())
		for _, deviceID := range deviceIDs {
			ds.ActiveEffects.Add(deviceID, cancel)
		}

		go func(ctx context.Context, speed float32, targets ...uint64) {
			defer cancel()
			msSpeed := speed * 1000
			ticker := time.NewTicker(time.Duration(msSpeed))
			defer ticker.Stop()
			hues := generateHues(16)
			ds.Client.SetColor(ctx, set.NewSetColor(hues[0], math.MaxUint16, math.MaxUint16, math.MaxUint16, uint32(msSpeed)), targets...)
			currIdx := 1
			for {
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					ds.Client.SetColor(ctx, set.NewSetColor(hues[currIdx], math.MaxUint16, math.MaxUint16, math.MaxUint16, uint32(msSpeed)), targets...)
					currIdx = (currIdx + 1) % len(hues)
				}
			}
		}(ctxCancel, request.GetSpeed(), deviceIDs...)
	}

	return resp, nil
}

func (ds *DeviceServer) Twinkle(ctx context.Context, request *pb.TwinkleRequest) (*pb.TwinkleResponse, error) {
	resp := &pb.TwinkleResponse{}

	deviceColors := request.GetDeviceColors()
	for deviceID := range deviceColors {
		ds.ActiveEffects.CancelEffect(deviceID)
	}

	if request.TurnOn {
		ctxCancel, cancel := context.WithCancel(context.Background())
		for deviceID := range deviceColors {
			ds.ActiveEffects.Add(deviceID, cancel)
		}
		go func(ctx context.Context, speed float32, intensity float32, deviceColors map[uint64]*pb.HSBK) {
			defer cancel()
			msSpeed := speed * 1000
			ticker := time.NewTicker(time.Duration(msSpeed) * time.Millisecond)
			defer ticker.Stop()
			for {
				for deviceID, color := range deviceColors {
					n := rand.Intn(10)
					if n < 5 {
						newWaveForm := set.NewSetWaveform(true, uint16(color.Hue), uint16(color.Saturation), uint16(math.Floor(float64(color.Brightness)*float64(1-intensity))), uint16(color.Kelvin), uint32(msSpeed), 1, 0, enum.HALF_SINE)
						ds.Client.SetWaveform(ctx, newWaveForm, deviceID)
					}
				}
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					fmt.Println("run")
					for deviceID, color := range deviceColors {
						n := rand.Intn(10)
						if n < 5 {
							newWaveForm := set.NewSetWaveform(true, uint16(color.Hue), uint16(color.Saturation), uint16(math.Floor(float64(color.Brightness)*float64(1-intensity))), uint16(color.Kelvin), uint32(msSpeed), 1, 0, enum.HALF_SINE)
							ds.Client.SetWaveform(ctx, newWaveForm, deviceID)
						}
					}
				}
			}
		}(ctxCancel, request.GetSpeed(), request.GetIntensity(), deviceColors)
	}
	return resp, nil
}

func (ds *DeviceServer) Theater(ctx context.Context, request *pb.TheaterRequest) (*pb.TheaterResponse, error) {
	resp := &pb.TheaterResponse{}
	deviceIDs := request.GetDeviceIDs()
	for _, deviceID := range deviceIDs {
		ds.ActiveEffects.CancelEffect(deviceID)
	}

	if request.TurnOn {
		ctxCancel, cancel := context.WithCancel(context.Background())
		for _, deviceID := range deviceIDs {
			ds.ActiveEffects.Add(deviceID, cancel)
		}
		go func(ctx context.Context, deviceIDs []uint64, screen int) {
			defer cancel()
			ticker := time.NewTicker(100 * time.Millisecond)
			defer ticker.Stop()
			for {
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					color, err := theater.GenerateDominantColorPalette(1, screen)
					if err != nil {
						fmt.Println(err)
						continue
					}
					go ds.Client.SetColor(ctx, set.NewSetColor(uint16(color[0].Hue), uint16(color[0].Saturation), uint16(color[0].Brightness), uint16(color[0].Kelvin), 200), deviceIDs...)
				}
			}
		}(ctxCancel, deviceIDs, int(request.GetScreen()))
	}

	return resp, nil
}

func (ds *DeviceServer) Visualizer(ctx context.Context, request *pb.VisualizerRequest) (*pb.VisualizerResponse, error) {
	resp := &pb.VisualizerResponse{}
	deviceIDs := request.GetDeviceIDs()
	for _, deviceID := range deviceIDs {
		ds.ActiveEffects.CancelEffect(deviceID)
	}

	if request.TurnOn {
		ctxCancel, cancel := context.WithCancel(context.Background())
		for _, deviceID := range deviceIDs {
			ds.ActiveEffects.Add(deviceID, cancel)
		}
		go func(ctx context.Context, deviceIDs []uint64) {
			defer cancel()
			beatdetector.Run(ctx, ds.Client, deviceIDs...)
		}(ctxCancel, deviceIDs)
	}

	return resp, nil
}
