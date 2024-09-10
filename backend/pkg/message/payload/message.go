package payload

type PayloadType uint16
type ResponseType uint16

// Queries
const (
	//Discovery
	GetService PayloadType = 2
	//Device
	GetHostFirmware PayloadType = 14
	GetWifiInfo     PayloadType = 16
	GetWifiFirmware PayloadType = 18
	GetPower        PayloadType = 20
	GetLabel        PayloadType = 23
	GetVersion      PayloadType = 32
	GetInfo         PayloadType = 34
	GetLocation     PayloadType = 48
	GetGroup        PayloadType = 51
	EchoRequest     PayloadType = 58
	//Light
	GetColor                 PayloadType = 101
	GetLightPower            PayloadType = 116
	GetInfrared              PayloadType = 120
	GetHevCycle              PayloadType = 142
	GetHevCycleConfiguration PayloadType = 145
	GetLastHevCycleResult    PayloadType = 148
	//MultiZone
	GetColorZones         PayloadType = 502
	GetMultiZoneEffect    PayloadType = 507
	GetExtendedColorZones PayloadType = 511
	//Relay
	GetRPower PayloadType = 816
	//Tile
	GetDeviceChain        PayloadType = 701
	Get64                 PayloadType = 707
	GetTileEffect         PayloadType = 718
	GetSensorAmbientLight PayloadType = 401
)

// Changing a device
const (
	//Device
	SetPower    PayloadType = 21
	SetLabel    PayloadType = 24
	SetReboot   PayloadType = 38
	SetLocation PayloadType = 49
	SetGroup    PayloadType = 52
	//Light
	SetColor                 PayloadType = 102
	SetWaveform              PayloadType = 103
	SetLightPower            PayloadType = 117
	SetWaveformOptional      PayloadType = 119
	SetInfrared              PayloadType = 122
	SetHevCycle              PayloadType = 143
	SetHevCycleConfiguration PayloadType = 146
	//Multizone
	SetColorZones      PayloadType = 501
	SetMultiZoneEffect PayloadType = 508
	//Relay
	SetExtendedColorZones PayloadType = 510
	SetRPower             PayloadType = 817
	//Tile
	SetUserPosition PayloadType = 703
	Set64           PayloadType = 715
	SetTileEffect   PayloadType = 719
)

// Response Messages
const (
	//Core
	Acknowledgement ResponseType = 45
	//Discovery
	StateService ResponseType = 3
	//Device
	StateHostFirmware ResponseType = 15
	StateWifiInfo     ResponseType = 17
	StateWifiFirmware ResponseType = 19
	StatePower        ResponseType = 22
	StateLabel        ResponseType = 25
	StateVersion      ResponseType = 33
	StateInfo         ResponseType = 35
	StateLocation     ResponseType = 50
	StateGroup        ResponseType = 53
	EchoResponse      ResponseType = 59
	StateUnhandled    ResponseType = 223
	//Light
	LightState                 ResponseType = 107
	StateLightPower            ResponseType = 118
	StateInfrared              ResponseType = 121
	StateHevCycle              ResponseType = 144
	StateHevCycleConfiguration ResponseType = 147
	StateLastHevCycleResult    ResponseType = 149
	//MultiZone
	StateZone               ResponseType = 503
	StateMultiZone          ResponseType = 506
	StateMultiZoneEffect    ResponseType = 509
	StateExtendedColorZones ResponseType = 512
	//Relay
	StateRPower ResponseType = 818
	//Tile
	StateDeviceChain        ResponseType = 702
	State64                 ResponseType = 711
	StateTileEffect         ResponseType = 720
	SensorStateAmbientLight ResponseType = 402
)
