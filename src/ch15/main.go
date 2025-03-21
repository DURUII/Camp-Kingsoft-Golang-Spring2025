package main

import (
	utils "ch15/series" // 设置别名
	"fmt"
	// 此外，还可以写 github 链接，注意 go-get 仓库不要 src 目录
)

func main() {
	fmt.Println(utils.GetFibonacciList(5))
}
