package ch6

import "testing"

func TestArray(t *testing.T) {
	var a [3]int
	t.Log(a)
	// 注意，加...的时候，表示定长数组
	b := [...]int{1, 2, 3, 4, 5, 6}
	t.Log(b)
	// Go 语言有严格约束，声明必须使用
	// 可以由_占位，作为变量名，用于存储无用的值
	//（_ 使用与 Python 类似，但 Python 不强制要求必须使用）
	for _, val := range b {
		t.Log(val)
	}

	// 数组截取（与 Python 类似）
	// 与 Python 不同，不支持步进，也不支持 -1 等倒数元素
	t.Log(b[3:])
}
