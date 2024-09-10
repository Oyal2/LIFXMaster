package structure

// Tile represents the information for a single device in a chain.
type Tile struct {
	AccelMeasX           int16 `json:"accel_meas_x"` // Accelerometer measurement X-axis
	AccelMeasY           int16 `json:"accel_meas_y"` // Accelerometer measurement Y-axis
	AccelMeasZ           int16 `json:"accel_meas_z"` // Accelerometer measurement Z-axis
	Reserved2            [2]byte
	UserX                float32 `json:"user_x"` // X coordinate in user-defined positioning
	UserY                float32 `json:"user_y"` // Y coordinate in user-defined positioning
	Width                uint8   `json:"width"`  // Number of zones per row
	Height               uint8   `json:"height"` // Number of zones per column
	Reserved7            [1]byte
	DeviceVersionVendor  uint32 `json:"device_version_vendor"`  // Vendor ID of the device
	DeviceVersionProduct uint32 `json:"device_version_product"` // Product ID of the device
	Reserved4            [4]byte
	FirmwareBuild        uint64 `json:"firmware_build"` // Firmware build time (epoch time)
	Reserved8            [8]byte
	FirmwareVersionMinor uint16 `json:"firmware_version_minor"` // Minor firmware version number
	FirmwareVersionMajor uint16 `json:"firmware_version_major"` // Major firmware version number
	Reserved10           [4]byte
}
