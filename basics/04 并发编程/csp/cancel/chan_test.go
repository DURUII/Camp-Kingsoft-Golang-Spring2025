package cancel

import (
	"fmt"
	"testing"
	"time"
)

func isCanceledByChannel(cancelChan chan struct{}) bool {
	// select 监听多个通道实现多路复用
	select {
	case <-cancelChan:
		return true // 收到取消消息

	default:
		return false // 未收到取消消息
	}
}

func cancel1(cancelChan chan struct{}) {
	// 只有一个 cancel 信号
	cancelChan <- struct{}{}

}

func cancel2(cancelChan chan struct{}) {
	// 不需要知道有多少个 runner
	close(cancelChan)
}

func TestCancelByChannel(t *testing.T) {
	cancelChan := make(chan struct{}, 1)
	for i := 0; i < 10; i++ {
		go func(i int, cancelCh chan struct{}) {
			for {
				if isCanceledByChannel(cancelChan) {
					break
				}
				time.Sleep(5 * time.Millisecond)
			}
			fmt.Println(i, "is canceled")
		}(i, cancelChan)
	}
	// 广播机制
	cancel2(cancelChan)
	time.Sleep(time.Second * 1)
}
