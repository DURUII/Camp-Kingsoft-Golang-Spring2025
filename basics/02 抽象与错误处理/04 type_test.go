package ch02

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

// 自定义类型，让程序简单可读
type IntConv func(int) int

func timeit(inner IntConv) IntConv {
	// 装饰器
	return func(x int) int {
		tic := time.Now()
		ret := inner(x)
		fmt.Println("time spent =>", time.Since(tic).Seconds())
		return ret
	}
}

func TestCustomizedType(t *testing.T) {
	t.Log(timeit(cubic)(9))
}
