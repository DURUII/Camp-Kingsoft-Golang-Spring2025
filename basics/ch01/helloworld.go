package main

// 和 Java 类似，模块，但是 package 名不需要与目录名一致
// 但需要约定同一目录的 Go 代码 package 要保持一致
import (
	"fmt"
	"os"
) // 引入依赖

// （非测试）程序入口必须是 main 包 + main 函数，文件名不一定是 main.go
// main 函数并不一定是第一个被执行的函数，Go 语言会先初始化常量变量，并自动调用包初始化的 init 函数
func main() {
	// main 函数不支持入参，极简的 Go 语言 不需要括号、分号
	if len(os.Args) > 1 {
		// 大写代表包外可以访问
		fmt.Println("hello world!", os.Args[1])
	}
	// 与 C 不同的是，main 函数不支持返回值
	os.Exit(-1)
}
