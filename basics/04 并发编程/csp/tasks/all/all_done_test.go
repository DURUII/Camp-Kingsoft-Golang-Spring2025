package all

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
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}

	// 等所有结果返回
	finalRet := ""
	for j := 0; j < numOfRunner; j++ {
		// unordered aggregation
		finalRet += <-ch + "\n"
	}
	return finalRet
}

func TestAllResponse(t *testing.T) {
	t.Log("Before: ", runtime.NumGoroutine())
	t.Log(AllResponse())
	time.Sleep(2 * time.Second)
	t.Log("After: ", runtime.NumGoroutine())
}
