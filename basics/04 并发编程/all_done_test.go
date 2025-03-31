package ch04

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("Result from %d", id)
}

func AllResponse() string {
	numOfRunner := 10
	// 协程阻塞泄漏，可能导致资源耗尽
	// ch := make(chan string)
	// 采用 buffered channel 将 发送/接受 解耦
	ch := make(chan string, 1)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}

	finalRet := ""
	for j := 0; j < numOfRunner; j++ {
		// aggregation
		finalRet += <-ch + "\n"
	}
	return finalRet
}

func TestAllResponse(t *testing.T) {
	t.Log("Before: ", runtime.NumGoroutine())
	t.Log(AllResponse())
	time.Sleep(2 * time.Second)
	t.Log("Before: ", runtime.NumGoroutine())
}
