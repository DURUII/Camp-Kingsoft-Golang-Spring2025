package unit

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert" // 引入断言
)

func malSquare(i int) int {
	return i*i + rand.Intn(3)
}

func TestSquare(t *testing.T) {
	// 表格测试/表驱动测试
	cases := []struct {
		input    int
		expected int
	}{
		{1, 1},
		{2, 4},
		{3, 9},
	}
	for _, c := range cases {
		// 也可以使用 t.Fail/Error/Fatal
		got := malSquare(c.input)
		want := c.expected
		assert.Equal(t, want, got)
	}
}
