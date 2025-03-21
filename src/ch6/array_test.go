package ch6

import "testing"

func TestArray(t *testing.T) {
	var a [3]int
	t.Log(a)
	// 注意，加...的时候，表示定长数组
	b := [...]int{1, 2, 3, 4, 5, 6}
	t.Log(b)
	// Go 语言有严格约束，声明必须使用，可以由_占位
	//（_ 使用与 Python 类似，但 Python 不强制要求必须使用）
	for _, val := range b {
		t.Log(val)
	}

	// 数组截取（与 Python 类似）
	// 与 Python 不同，不支持步进，也不支持 -1 等倒数元素
	t.Log(b[3:])
}

func TestSlice(t *testing.T) {
	// 没有...表示，一个切片（可变长数组）
	var s0 []int
	t.Log(len(s0), cap(s0))

	// 切片是一个结构体，有len/cap
	// 用 make 指定容量，避免扩容复制
	var s1 = make([]int, 3, 4)
	// 当存储空间不足，自动开辟新的存储空间，并拷贝原有数值
	s1 = append(s1, 10)
	t.Log(s1, len(s1), cap(s1))
	s1 = append(s1, 40)
	// 容量成倍增长
	t.Log(s1, len(s1), cap(s1))
}

func TestSlice2(t *testing.T) {
	year := []string{
		"Jan", "Feb", "Mar", "Apr", "May",
		"Jun", "Jul", "Aug", "Sep", "Oct", "Nov",
		"Dec",
	}
	Q2 := year[3:6]
	// 这里的 cap 是 9，剩余连续存储空间
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5:8]
	summer[0] = "Unknown"
	// 共享存储空间
	t.Log(year, Q2, summer)
}
