package ch04

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(500 * time.Millisecond)
	return "done"
}

// 类似于 Java 中的 FutureTask
func AsyncService() chan string {
	retChan := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("returned result")
		retChan <- ret
		fmt.Println("service exited")
	}()
	return retChan
}

func TestAsync(t *testing.T) {
	// 超时机制：slow response 是比 quick failure 还可怕的错误
	select {
	case ret := <-AsyncService():
		t.Log(ret)
	case <-time.After(500 * time.Millisecond):
		t.Log("timeout")
	}

}
