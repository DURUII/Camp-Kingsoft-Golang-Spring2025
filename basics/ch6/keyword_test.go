package ch6

import (
	"fmt"
	"reflect"
	"testing"
)

func TestKeyword(t *testing.T) {
	var a [3]int
	var b []int
	// 数组切片的初始值为元素零值
	fmt.Println(reflect.TypeOf(a), "a=", a)
	// 切片的初始值为 nil
	fmt.Println(reflect.TypeOf(b), "b=", b, b == nil)

	// nil 不是 25 个关键字之一，因为可以被覆盖作为变量名 (不推荐)
	// shadows declaration at builtin.go
	nil := []int{1, 2, 3}
	// append 要赋值给新的变量
	b = append(b, 1, 2, 3)
	fmt.Println(reflect.TypeOf(b), "b=", b, reflect.DeepEqual(b, nil))
}
