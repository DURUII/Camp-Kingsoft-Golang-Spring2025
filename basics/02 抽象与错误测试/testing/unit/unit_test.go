package unit

import (
	"github.com/stretchr/testify/assert" // 引入断言
	"math/rand"
	"testing"
)

func malSquare(i int) int {
	return i*i + rand.Intn(5)
}

func TestSquare(t *testing.T) {
	// 表格测试
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}
	for i := 0; i < len(inputs); i++ {
		// 也可以使用 t.Fail/Error/Fatal
		assert.Equal(t, expected[i], malSquare(inputs[i]))
	}
}
