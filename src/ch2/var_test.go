package ch2

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {
	// 类型关键字在后面（一般用于变量声明，主要是全局/外部）
	// 静态类型（与 Python 相同）
	//var a int = 1
	//var b int = 1

	// 简写形式，具有自动类型推断* （与 Java 不同）
	//var (
	//	a = 1
	//	b = 1
	//)

	// 最简单写法，支持多个变量赋值* （与 Python 相同）
	a, b := 1, 1

	fmt.Print(a, ",")
	for i := 0; i < 10; i++ {
		//tmp := a
		//a = b
		//b = tmp + a
		// 支持多个变量赋值* （与 Python 相同）
		a, b = b, a+b
		fmt.Print(a, ",")
	}
	fmt.Println(b)
}

func TestExchange(t *testing.T) {
	a, b := 1, 3
	t.Log("a=", a, "b=", b)
	a, b = b, a
	t.Log("a=", a, "b=", b)
}
