package ch10

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func cubic(x int) int {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	return x * x * x
}

func timeit(inner func(int) int) func(int) int {
	// 装饰器
	return func(x int) int {
		tic := time.Now()
		ret := inner(x)
		fmt.Println("time spent =>", time.Since(tic).Seconds())
		return ret
	}
}

func TestFunc(t *testing.T) {
	// 值传递，函数可以作为变量值、参数、返回值，可以有多个返回值
	t.Log(timeit(cubic)(9))
}

func sum(ops ...int) int {
	// 可变长参数会被转换为一个数组
	if len(ops) == 0 {
		return 0
	}
	// 语法将切片解包传递给递归调用（Python 用 * 表示）
	return sum(ops[1:]...) + ops[0]
}

func sum2(ops ...int) int {
	sum := 0
	for _, op := range ops {
		sum += op
	}
	return sum
}

func TestVariadic(t *testing.T) {
	t.Log(sum(1, 2, 3, 4, 5))
	t.Log(sum2(1, 2, 3, 4, 5))
}

func TestDefer(t *testing.T) {
	// 函数（异常）返回前，类似于 Java 的 finally
	// 通常用于（安全地）释放资源
	defer func() {
		fmt.Println("Clear resources")
	}()
	fmt.Println("Started")
	// defer 仍然会执行
	panic("Fatal error")
}
