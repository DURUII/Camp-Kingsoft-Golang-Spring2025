package ch7

import "testing"

func TestMap(t *testing.T) {
	// 声明 + 初始化
	m1 := map[int]int{1: 1, 2: 4, 3: 9, 4: 16}
	t.Log(m1[2])
	t.Logf("len m1=%d", len(m1))

	// 声明
	m2 := map[int]int{}
	m2[3] = 27
	t.Logf("len m2=%d", len(m2))

	// make 方法，可以设定 cap，但不能用 cap(m3)
	m3 := make(map[string]int, 10)
	t.Log(m3, len(m3))
}

func TestNonExist(t *testing.T) {
	// Go 语言没有提供 set，但可以通过 map 实现类似的功能
	m1 := map[int]int{2: 8}

	/*```python
	map = {2:8}

	In [9]: map[3]
	KeyError: 3

	In [12]: type(map.get(3))
	Out[12]: NoneType
	```*/

	/*```java
	import java.util.*;

	public class MyClass {
	  public static void main(String args[]) {
		Map<Integer, Integer> map = new HashMap<>();
		// null
		System.out.println(map.get(4));
	  }
	}
	```*/

	// 当 key 不存在时，默认返回零值
	// 这类似 Python 中 map.get(key, default)，而 default 系统自动返回
	if val, ok := m1[3]; ok {
		t.Log("retrieved val =", val)
	} else {
		t.Log("key not found")
	}
}

func TestMapIteration(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9, 4: 16}
	// 数组第一个值是 index，不是 key
	for k, v := range m1 {
		t.Log(k, "=>", v)
	}
	t.Log(len(m1))

	// 删除元素
	delete(m1, 4)
	t.Log(len(m1))
}
