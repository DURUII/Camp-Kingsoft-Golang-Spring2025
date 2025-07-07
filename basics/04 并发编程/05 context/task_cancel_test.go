package _context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

// 关联任务的取消（级联退出）
// context 用于管理相关任务的上下文，包含了共享值的传递，超时，取消通知
func isCanceledByContext(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func TestCancelByContext(t *testing.T) {
	// 父节点
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(ctx context.Context, i int) {
			for {
				if isCanceledByContext(ctx) {
					break
				}
				time.Sleep(5 * time.Millisecond)
			}
			fmt.Println(i, "is canceled")
		}(ctx, i)
	}
	// 广播机制
	cancel()
	time.Sleep(time.Second * 1)
	fmt.Println("now n = ", runtime.NumGoroutine())
}
