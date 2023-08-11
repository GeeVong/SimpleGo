package main

import (
	"fmt"
)

// array
func arrayOperate() {
	// 变量定义,需要制定大小
	var arr [5]int32
	fmt.Println("emp:", arr)

	// set the value of arr in position 1
	arr[1] = 10
	fmt.Println("set:", arr)
	fmt.Println("get:", arr[4])

	// 类型推断
	arr_1 := [7]int{1, 23, 34, 11, 1, 1123, 113}
	fmt.Printf("%d\n", arr_1)

	// 2d arr
	var _2dArray [2][3]int32
	fmt.Println("_2dArray:", _2dArray)
	for i := 0; i < len(_2dArray); i++ {
		for j := 0; j < len(_2dArray[i]); j++ {
			_2dArray[i][j] = int32(i + j)
		}
	}
	fmt.Println(_2dArray)

}

func main() {
	arrayOperate()

}
