package enum

// Services represents the available services.
type Services uint8

const (
	UDP Services = iota + 1
	_
	_
	_
	_
)

// Direction indicates the direction of an effect or operation.
type Direction uint8

const (
	RIGHT Direction = iota
	LEFT
)

// LightLastHevCycleResult represents possible outcomes of a cycle.
type LightLastHevCycleResult uint8

const (
	SUCCESS LightLastHevCycleResult = iota
	BUSY
	INTERRUPTED_BY_RESET
	INTERRUPTED_BY_HOMEKIT
	INTERRUPTED_BY_LAN
	INTERRUPTED_BY_CLOUD
	NONE = 255
)

// MultiZoneApplicationRequest represents application request types for multi-zone operations.
type MultiZoneApplicationRequest uint8

const (
	NO_APPLY MultiZoneApplicationRequest = iota
	APPLY
	APPLY_ONLY
)

// MultiZoneEffectType defines types of effects applicable to multi-zone setups.
type MultiZoneEffectType uint8

const (
	OFF MultiZoneEffectType = iota
	MOVE
	_
	_
)

// MultiZoneExtendedApplicationRequest represents extended application request types for multi-zone operations.
type MultiZoneExtendedApplicationRequest uint8

const (
	NO_APPLY_EXTENDED MultiZoneExtendedApplicationRequest = iota
	APPLY_EXTENDED
	APPLY_ONLY_EXTENDED
)

// TileEffectType defines types of effects applicable to tile setups.
type TileEffectType uint8

const (
	TILE_OFF TileEffectType = iota
	_
	MORPH
	FLAME
	_
)

// Waveform represents different waveform types for light modulation.
type Waveform uint8

const (
	SAW Waveform = iota
	SINE
	HALF_SINE
	TRIANGLE
	PULSE
)
