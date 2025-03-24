package ch9

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringPkg(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	t.Log(strings.Join(parts, "->"))
	// 拼写历史来由：Integer to ASCII
	s = strconv.Itoa(10)
	t.Log(string('*') + s + "*") // 注意强制类型转换
	if i, err := strconv.Atoi("100"); err == nil {
		t.Log(100 + i)
	}
}
