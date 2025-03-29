package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestArray(t *testing.T) {
	// Go 的数组类型包含两个重要属性：元素的类型和数组长度。如果有一个属性不同，它们就是两个不同的数组类型
	var a [3]int
	t.Log(a)
	// 注意，加...的时候，表示（定长）数组
	b := [...]int{1, 2, 3, 4, 5, 6}
	t.Log(b, reflect.TypeOf(b) == reflect.TypeOf(a))

	// 稀疏数组初始化
	c := [10]int{
		2: 99,
	}
	t.Log(c)

	// Go 语言有严格约束，声明必须使用
	// 可以由_占位，作为变量名，用于存储无用的值
	//（_ 使用与 Python 类似，但 Python 不强制要求必须使用）
	for _, val := range b {
		t.Log(val)
	}

	// 数组截取（与 Python 类似）
	// 与 Python 不同，不支持步进，不支持 -1 等倒数元素
	t.Log(b[3:], reflect.TypeOf(b[3:]), reflect.TypeOf(b))
}

func TestRefArray(t *testing.T) {
	a := [3]int{1, 2, 3}
	fn := func(x [3]int) { x[1] = 4 }
	fn(a)
	t.Log(a)
}

func TestArrayForRangeBug(t *testing.T) {
	a := []int{1, 2, 3}
	b := []*int{}
	// 注意，遍历过程没有返回集合中的实际元素
	// 而是复制在了一个“固定”的变量（值传递），例如 i
	// *复制操作可能会存在性能问题
	for _, i := range a {
		b = append(b, &i)
	}
	// 相当于
	//var i int
	//for k := 0; k < len(a); k++ {
	//	i = a[k]
	//	b = append(b, &i)
	//}

	for _, j := range b {
		fmt.Print(*j, " ")
	}

}
