package device

import (
	"encoding/json"
	"net"

	"github.com/oyal2/LIFXMaster/pkg/message/payload/state"
)

type LightDevice struct {
	Addr              net.IP                   `json:"address,omitempty"`
	Port              uint32                   `json:"port,omitempty"`
	Target            uint64                   `json:"target,omitempty"`
	Firmware          *state.HostFirmware      `json:"firmware,omitempty"`
	Wifi              *Wifi                    `json:"wifi,omitempty"`
	Power             *state.Power             `json:"power,omitempty"`
	Label             *state.Label             `json:"label,omitempty"`
	Version           *state.Version           `json:"version,omitempty"`
	Info              *state.Info              `json:"info,omitempty"`
	Location          *state.Location          `json:"location,omitempty"`
	Group             *state.Group             `json:"group,omitempty"`
	Product           *Product                 `json:"product,omitempty"`
	Light             *state.Light             `json:"light,omitempty"`
	Infrared          *state.Infrared          `json:"infrared,omitempty"`
	HEV               *HEV                     `json:"hev,omitempty"`
	ExtendedColorZone *state.ExtendedColorZone `json:"extended_color_zone,omitempty"`
	Relay             *Relay                   `json:"relay,omitempty"`
	Tile              *Tile                    `json:"tile,omitempty"`
}

type Tile struct {
	DeviceChain        *state.DeviceChain        `json:"deviceChain,omitempty"`
	Tile64             *state.Tile64             `json:"tile_64,omitempty"`
	TileEffect         *state.TileEffect         `json:"tile_effect,omitempty"`
	SensorAmbientLight *state.SensorAmbientLight `json:"sensor_ambient_light,omitempty"`
}

type Relay struct {
	RPower *state.RPower `json:"r_power,omitempty"`
}

type Wifi struct {
	Info     *state.WifiInfo
	Firmware *state.WifiFirmware
}

type HEV struct {
	Cycle           *state.HEVCycle
	Config          *state.HEVCycleConfig
	LastCycleResult *state.LastHEVCycleResult
}

func NewLightBulb(addr net.IP, port uint32, target uint64) *LightDevice {
	return &LightDevice{
		Addr:              addr,
		Port:              port,
		Target:            target,
		Firmware:          &state.HostFirmware{},
		Wifi:              &Wifi{Info: &state.WifiInfo{}, Firmware: &state.WifiFirmware{}},
		Power:             &state.Power{},
		Label:             &state.Label{},
		Version:           &state.Version{},
		Info:              &state.Info{},
		Location:          &state.Location{},
		Group:             &state.Group{},
		Product:           &Product{},
		Light:             &state.Light{},
		Infrared:          &state.Infrared{},
		HEV:               &HEV{Cycle: &state.HEVCycle{}, Config: &state.HEVCycleConfig{}, LastCycleResult: &state.LastHEVCycleResult{}},
		ExtendedColorZone: &state.ExtendedColorZone{},
		Relay:             &Relay{RPower: &state.RPower{}},
		Tile: &Tile{
			DeviceChain:        &state.DeviceChain{},
			Tile64:             &state.Tile64{},
			TileEffect:         &state.TileEffect{},
			SensorAmbientLight: &state.SensorAmbientLight{},
		},
	}
}

func (ld *LightDevice) JSON() (string, error) {
	lightDeviceBytes, err := json.MarshalIndent(ld, "", "\t")
	return string(lightDeviceBytes), err
}
