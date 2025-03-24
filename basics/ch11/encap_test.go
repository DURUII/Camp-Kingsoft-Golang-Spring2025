package ch11

import (
	"fmt"
	"testing"
	"unsafe"
)

// 对数据的封装（与 C 类似），没有逗号
type Employee struct {
	Id   string
	Name string
	Age  int
}

// 对行为的定义，不写在 struct 里，不是典型的 OOP
func (e *Employee) String() string {
	// 没有对象复制产生
	fmt.Println("Address = ", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("{id:%s name:%s age:%d}", e.Id, e.Name, e.Age)
}

// Go 语言不支持继承
func TestCreateData(t *testing.T) {
	e := Employee{"1", "Jack", 20}
	// 返回指针类型
	e2 := new(Employee)
	e2.Id = "123"
	e2.Name = "Tom"
	e2.Age = 30
	// ch11.Employee, *ch11.Employee
	t.Logf("%T", e)
	t.Logf("%T", e2)
	// 不需要箭头符号（与 C 不同）
	t.Log(e2.String())
	fmt.Println("Address = ", unsafe.Pointer(&e2.Name))
}
