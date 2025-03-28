# Go 语言基础语法

> Go 语言的设计哲学是“做减法”、“显式”、“组合”、“并发”。

## Hello, Go

下划线在 Go 语言中特殊作用，Go 的命名规则是用全小写字母形式的短小单词命名，无需分隔符，例如，你可以创建`helloworld.go`。

```go
package main

// 和 Java 类似，模块，但是 package 名不需要与目录名一致
// 但需要约定同一目录的 Go 代码 package 要保持一致
import (
	"fmt"
	"os"
) // 引入依赖

// （非测试）程序入口必须是 main 包 + main 函数，文件名不一定是 main.go
func main() {
	// main 函数不支持入参，极简的 Go 语言 不需要括号、分号
	if len(os.Args) > 1 {
		// 大写代表包外可以访问
		fmt.Println("hello world!", os.Args[1])
	}
	// 与 C 不同的是，main 函数不支持返回值
	os.Exit(-1)
}

```

1. `import` 的 `fmt` 代表包的路径（默认在标准库目录下），而 `fmt.Println(...)` 代表包名。这里`Println`之所以首字母要大写，是因为首字母大写的标识符在 Go 语言中代表导出的，即对包外可见的。 
2. 与 C/C++/Java 不同，Go 语言的结尾标识符`;`是可选的。
3. Go 是一种编译型语言，编译后即时没有安装 Go 的环境也可以运行可执行文件。你可以用`go run .`测试运行结果，真正交付的成果是用`go build .`命令构建的。
4. 一个`module`是一个包的集合，你可以用`go mod init/tidy`生成该文件。

## 

