package device

type DeviceProfiler interface {
	JSON() (string, error)
}
