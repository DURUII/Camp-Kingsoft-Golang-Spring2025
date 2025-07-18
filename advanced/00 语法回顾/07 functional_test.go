package ch00

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type ServerOption func(*Server)

func NewServer(addr string, port int, options ...ServerOption) *Server {
	srv := &Server{
		Addr: addr,
		Port: port,
	}
	for _, option := range options {
		option(srv)
	}
	return srv
}

// WithProtocol is an option to set the protocol of the server.
func WithProtocol(protocol string) ServerOption {
	return func(s *Server) {
		if protocol != "tcp" && protocol != "udp" {
			s.Err = errors.Join(s.Err, ErrInvalidProtocol)
		}
		s.Protocol = protocol
	}
}

// WithMaxConn is an option to set the max connections of the server.
func WithMaxConn(maxConn int) ServerOption {
	return func(s *Server) {
		if maxConn <= 0 {
			s.Err = errors.Join(s.Err, ErrInvalidMaxConn)
		}
		s.MaxConn = maxConn
	}
}

// WithTimeout is an option to set the timeout of the server.
func WithTimeout(timeout time.Duration) ServerOption {
	return func(s *Server) {
		if timeout <= 0 {
			s.Err = errors.Join(s.Err, ErrInvalidTimeout)
		}
		s.Timeout = timeout
	}
}

func TestServerOptions(t *testing.T) {
	srv := NewServer("127.0.0.1", 8080, WithProtocol("tcp"), WithMaxConn(1000), WithTimeout(100*time.Millisecond))
	assert.Equal(t, "tcp", srv.Protocol)
	assert.Equal(t, 1000, srv.MaxConn)
	assert.Equal(t, 100*time.Millisecond, srv.Timeout)
}
