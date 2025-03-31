package benchmark

import (
	"strings"
	"testing"
)

// 基准测试：go test -bench <to-test-func> -benchmem
func BenchmarkConcatStringByAdd(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ret := ""
		for _, elem := range elems {
			ret += elem
		}
		//fmt.Println(ret)
	}
	b.StopTimer()
}

func BenchmarkConcatStringByBuilder(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 也可以写 bytes.Buffer 类型
		var s strings.Builder
		for _, elem := range elems {
			s.WriteString(elem)
		}
		//fmt.Println(s.String())
		_ = s.String()
	}
	b.StopTimer()

}
