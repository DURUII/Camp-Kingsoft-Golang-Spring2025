package ch16

import (
	"fmt"
	"testing"
	"time"
)

func TestGoRoutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		// 有竞争关系
		//go func() {
		//	fmt.Println(i)
		//}()

		// 无竞争关系
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(1 * time.Second)
}
