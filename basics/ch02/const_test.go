package ch2

import "testing"

func TestConst(t *testing.T) {
	const (
		Mon = iota + 1
		Tue
		Wed
		Thu
		Fri
		Sat
		Sun
	)

	t.Log(Sun, Mon, Tue, Wed, Thu, Fri, Sat)
}

func TestConst2(t *testing.T) {
	const (
		Readable = 1 << iota
		Writable
		Executable
	)

	a := 6 // 110
	t.Log(
		a&Readable == Readable,
		a&Writable == Writable,
		a&Executable == Executable,
	)
}
