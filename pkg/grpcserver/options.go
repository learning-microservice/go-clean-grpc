package grpcserver

import (
	"fmt"
	"log/slog"
	"net"
	"time"
)

// Option -.
type Option func(*Server)

// Address -.
func Address(host string, port int) Option {
	return func(s *Server) {
		s.address = net.JoinHostPort(host, fmt.Sprint(port))
	}
}

// ReadHeaderTimeout -.
func ReadHeaderTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.readHeaderTimeout = timeout
	}
}

// ShutdownTimeout -.
func ShutdownTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.shutdownTimeout = timeout
	}
}

// Logger -.
func Logger(logger *slog.Logger) Option {
	return func(s *Server) {
		s.logger = logger
	}
}
