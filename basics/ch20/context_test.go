package ch20

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func isCanceled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func TestCancle(t *testing.T) {
	// 父节点
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCanceled(ctx) {
					break
				}
				time.Sleep(5 * time.Millisecond)
			}
			fmt.Println(i, "is cancled")
		}(i, ctx)
	}
	// 广播机制
	cancel()
	time.Sleep(time.Second * 1)
}

// context 用于管理相关任务的上下文，包含了共享值的传递，超时，取消通知
func TestContext(t *testing.T) {

}
