package ch00

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestServerBuilder(t *testing.T) {
	_, err := new(ServerBuilder).
		New("127.0.0.1", -1).
		WithProtocol("xxx").
		WithMaxConn(1024).
		WithTimeout(30 * time.Second).
		Build()
	assert.Equal(t, true, errors.Is(err, ErrInvalidPort))
	assert.Equal(t, false, errors.Is(err, ErrInvalidAddress))
	assert.Equal(t, true, errors.Is(err, ErrInvalidProtocol))
	assert.Equal(t, false, errors.Is(err, ErrInvalidMaxConn))
	assert.Equal(t, false, errors.Is(err, ErrInvalidTimeout))
}
