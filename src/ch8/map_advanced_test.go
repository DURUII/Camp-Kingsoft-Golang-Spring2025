package ch8

import (
	"testing"
)

func TestMapWithFuncValue(t *testing.T) {
	// 与 Python/JavaScript 类似，支持部分函数式编程特性
	intFunc := map[string]func(op int) int{}
	intFunc["cubic"] = func(op int) int { return op * op * op }
	t.Log(intFunc["cubic"](9))
}
