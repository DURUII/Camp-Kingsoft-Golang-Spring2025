package ch11

import "fmt"

// 父类
type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	// Go 是不⽀持继承的
	fmt.Println(" ", host)
}

// 子类
type Dog struct {
	Pet // 内嵌组合复用，无继承；不支持方法覆盖
}
