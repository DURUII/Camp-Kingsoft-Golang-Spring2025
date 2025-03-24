package ch4

import "testing"

func TestCompareArray(t *testing.T) {
	// Go 语言没有前置的 ++/-- （Python 连后置都没有）

	// 在比较数组时，主流语言比较引用是否相同
	// 1. 相同维数可以比较
	// 2. 每个元素都相同才相等 (与 Python 相同)
	a := [...]int{1, 2, 3}
	b := [...]int{1, 2, 4}
	t.Log(&a == &b)
	t.Log(a == b)
	b[2] = 3
	t.Log(a == b)
}

func TestName(t *testing.T) {
	// 按位清零
	const (
		Readable = 1 << iota
		Writable
		Executable
	)

	a := 7 // 111
	t.Log(
		a&Readable == Readable,
		a&Writable == Writable,
		a&Executable == Executable,
		a&^Readable == Readable,
		a&^Writable == Writable,
		a&^Executable == Executable,
	)
}
