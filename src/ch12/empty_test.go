package ch12

import (
	"fmt"
	"testing"
)

// 类似于 C 中的 void* 和 Java 中的 Object
func DoSomething(p interface{}) {
	// 类型 switch 专用语法
	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("String", v)
	default:
		fmt.Println("Unknown Type")
	}
}

func TestEmpty(t *testing.T) {
	DoSomething(1)
	DoSomething("1")
	DoSomething(fmt.Println)
}
