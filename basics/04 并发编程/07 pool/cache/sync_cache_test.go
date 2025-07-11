package cache

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object.")
			return 100
		},
	}

	pool.Put(1)
	pool.Put(2)
	pool.Put(3)
	runtime.GC()

	// 第一个 Get取出的是私有池中的，私有池仅可放一个对象
	// 后面的Get取出的则是共享池中的，而共享池的是后进先出的。
	fmt.Println(pool.Get().(int))
	fmt.Println(pool.Get().(int))
	fmt.Println(pool.Get().(int))
	fmt.Println(pool.Get().(int))
}
