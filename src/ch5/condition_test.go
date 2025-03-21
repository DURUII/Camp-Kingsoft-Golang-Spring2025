package ch5

import (
	"fmt"
	"runtime"
	"testing"
)

func TestSwitch(t *testing.T) {
	// switch 不限制常量整数/布尔值，可以是字符串
	// switch 句末不需要额外加 break （与 C 不同）
	switch os := runtime.GOOS; os {
	// 支持多结果选项
	case "darwin", "linux":
		fmt.Println("Linux-Based System")
	default:
		fmt.Printf("%s.\n", os)
	}
}
