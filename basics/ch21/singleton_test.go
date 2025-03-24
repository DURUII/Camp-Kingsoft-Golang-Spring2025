package ch21

import (
	"fmt"
	"sync"
	"testing"
)

// 单例模式（懒汉式，线程安全）
type Object struct {
	id int
}

var singleInstance *Object
var once sync.Once

func GetSingletonObject() *Object {
	//fmt.Println("Create Object")
	//singleInstance = new(Object)
	once.Do(func() {
		fmt.Println("Create Object")
		singleInstance = new(Object)
	})
	return singleInstance
}

func TestSingleton(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObject()
			fmt.Printf("Get SingletonObject: %p\n", obj)
			wg.Done()
		}()
	}
	wg.Wait()
}
