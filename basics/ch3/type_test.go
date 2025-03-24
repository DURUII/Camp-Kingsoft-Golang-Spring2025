package ch3

import (
	"math"
	"testing"
)

func TestImplicit(t *testing.T) {
	t.Log(math.MaxUint16)

	var (
		// Go 对于类型转换很严苛，只支持显式类型转换
		a int32 = 1
		b       = int64(a)
	)
	t.Log(a, b)
}
