package ch00

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestServerBuilder(t *testing.T) {
	srv, err := new(ServerBuilder).
		New("127.0.0.1", -1).
		WithProtocol("xxx").
		WithMaxConn(1024).
		WithTimeout(30 * time.Second).
		Build()
	assert.Error(t, err)
	assert.Equal(t, "127.0.0.1", srv.Addr)
	assert.Equal(t, 1024, srv.MaxConn)
	assert.Equal(t, 30*time.Second, srv.Timeout)
	assert.Equal(t, true, errors.Is(err, ErrInvalidPort))
	assert.Equal(t, false, errors.Is(err, ErrInvalidAddress))
	assert.Equal(t, true, errors.Is(err, ErrInvalidProtocol))
	assert.Equal(t, false, errors.Is(err, ErrInvalidMaxConn))
	assert.Equal(t, false, errors.Is(err, ErrInvalidTimeout))
}
