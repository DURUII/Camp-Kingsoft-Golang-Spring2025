package ch26

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func square(i int) int {
	return i*i + rand.Intn(5)
}

func TestSquare(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}
	for i := 0; i < len(inputs); i++ {
		assert.Equal(t, expected[i], square(inputs[i]))
	}
}
