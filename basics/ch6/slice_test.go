package ch6

import "testing"

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

func TestSlice3(t *testing.T) {
	
}
