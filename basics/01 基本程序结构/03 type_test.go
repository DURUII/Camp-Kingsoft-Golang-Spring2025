package main

import (
	"math"
	"testing"
)

type MyInt int64

// Type Alias
type MyFloat float64

func TestImplicitCasting(t *testing.T) {
	/*
		Go 很简洁得克制，无括号、无三元运算符、无继承、无指针运算
		它的设计哲学之一是：
		“显式”，要求程序员明确表达意图，即使多写几字，也不能“稀里糊涂”
	*/
	t.Log(math.MaxUint16)

	var (
		a int32   = 1
		b         = int64(a) // Go 对于类型转换很严苛，只支持显式类型转换
		c         = MyInt(a)
		d MyFloat = 1.0
	)
	t.Log(a, b, c, d)
}
