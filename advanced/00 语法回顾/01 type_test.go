package ch00

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

// Go 语言是
// 静态【编译期间而非运行时进行类型检查】
// 强类型【不允许不合理的隐式类型混合转换】系统

type myInt int
type myMap map[int]string

func TestTypeConversion(t *testing.T) {
	var i myInt
	var j = 1
	// 无法将 'i' (类型 myInt) 用作类型 int
	i = myInt(j)
	assert.Equal(t, 1, i)
	require.Equal(t, true, i == 1)

	//var p myMap
	//var q = map[int]string{1: "one"}
	//p = q
	//assert.Equal(t, p, q)
}
