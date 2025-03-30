package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func isCanceled(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

func cancel_1(cancelChan chan struct{}) {
	cancelChan <- struct{}{}

}

func cancel_2(cancelChan chan struct{}) {
	close(cancelChan)
}

func TestCancle(t *testing.T) {
	cancelChan := make(chan struct{}, 1)
	for i := 0; i < 10; i++ {
		go func(i int, cancelCh chan struct{}) {
			for {
				if isCanceled(cancelChan) {
					break
				}
				time.Sleep(5 * time.Millisecond)
			}
			fmt.Println(i, "is cancled")
		}(i, cancelChan)
	}
	// 广播机制
	cancel_2(cancelChan)
	time.Sleep(time.Second * 1)
}
