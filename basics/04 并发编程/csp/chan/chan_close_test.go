package _chan

import (
	"fmt"
	"sync"
	"testing"
)

// 生产者
func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		// 关闭通道，否则两个消费者都会死等数据，避免 dead lock
		// 向关闭的 channel 发送数据，会导致 panic
		close(ch)
		wg.Done()
	}()
}

// 消费者
func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			// ok 用于退出信号的广播机制
			if data, ok := <-ch; ok {
				fmt.Println(data)
			} else {
				break
			}
		}
		wg.Done()
	}()
}

func TestChanClose(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int, 10)

	wg.Add(1)
	dataProducer(ch, &wg)

	wg.Add(1)
	dataReceiver(ch, &wg)

	wg.Add(1)
	dataReceiver(ch, &wg)

	wg.Wait()
}
