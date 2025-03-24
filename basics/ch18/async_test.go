package ch18

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(2 * time.Second)
	return "done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(1 * time.Second)
	fmt.Println("task finished")
}

func TestSync(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

// 类似于 Java 中的 FutureTask
func asyncService() chan string {
	// channel
	//retChan := make(chan string, 1)
	// buffered channel
	retChan := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("returned result")
		// 放数据，只要没有接受消息，会被阻塞
		retChan <- ret
		fmt.Println("service exited")
	}()
	return retChan
}

func TestAsync(t *testing.T) {
	retChan := asyncService()
	otherTask()
	// 取数据
	fmt.Println(<-retChan)
}
