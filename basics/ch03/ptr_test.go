package ch3

import "testing"

func TestPointer(t *testing.T) {
	// 不支持指针运算
	var a = 1
	var aPtr = &a
	t.Log(a, aPtr)
	t.Logf("%T %T", a, &a)
}

func TestString(t *testing.T) {
	// 字符串是数值类型，的默认值是""，而不是 None/nil
	var str string
	t.Log("*"+str+"*", len(str))
}
