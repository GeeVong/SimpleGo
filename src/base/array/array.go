package main

import "fmt"

// array

func tArr() {
	// 变量定义,需要制定大小
	var arr [5]int32
	arr[1] = 10
	fmt.Println(arr)

	// 类型推断
	arr_1 := []int{1, 23, 34, 11, 1, 1123, 113}
	fmt.Printf("%d", arr_1)
}

// slice

func testSlice() {
	foo := make([]int, 5)

	foo[3] = 42
	foo[4] = 100
	bar := foo[1:4]
	bar[1] = 99

	fmt.Println(foo)
	fmt.Println(bar)
}

func main() {
	testSlice()

}
