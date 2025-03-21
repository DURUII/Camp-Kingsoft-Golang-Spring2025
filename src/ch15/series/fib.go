package series

import "fmt"

func init() {
	fmt.Print("A ")
}

func init() {
	fmt.Print("B ")
}

func GetFibonacciList(n int) []int {
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList[:]
}

func init() {
	fmt.Print("C ")
}
