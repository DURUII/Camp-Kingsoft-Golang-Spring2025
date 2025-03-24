package ch5

import (
	"fmt"
	"testing"
)

func TestForLoop(t *testing.T) {
	// Go 语言关键字极少，对于循环只支持 for
	// 与 C++/Java 不同，不需要括号

	/*```python
	for i in range(5):
		print(i)
	```*/

	/*```java
	class Main {
		public static void main(String[] args) {
			for (int i = 0; i < 5; i++){
				System.out.println(i);
			}
		}
	}
	```*/

	// 为保持简洁，Go 语言将规定自增只能是语句，而不是表达式
	// 不支持 ++i 或者 --i 操作
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// 也不支持链式赋值（与 C 不同）
	//i := 10
	//x := i++
}
