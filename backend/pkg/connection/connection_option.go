package connection

import (
	"net"
	"time"
)

type ConnectionOptions struct {
	Deadline time.Duration
	IP       net.IP
	Port     int
}

type ConnectionOption func(*ConnectionOptions)

func WithDeadline(deadline time.Duration) ConnectionOption {
	return func(s *ConnectionOptions) {
		s.Deadline = deadline
	}
}

func WithIP(ip net.IP) ConnectionOption {
	return func(s *ConnectionOptions) {
		s.IP = ip
	}
}

func WithPort(port int) ConnectionOption {
	return func(s *ConnectionOptions) {
		s.Port = port
	}
}
