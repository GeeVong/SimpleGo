package main

import "fmt"

/*

. 验证
	int（stack）
	make slice（heap）
	global param
在内存中地址高低分布情况
*/

var data_global_i int

func main() {
	var data_stack_int int

	// &data_heap_slice 0xc000113ed8	  切片地址
	// &data_heap_slice[0] 0xc000136000   底层数组首元素地址
	data_heap_slice := make([]int32, 1, 10)
	data_heap_slice[0] = 1

	// 0x118be68
	data_global_i = 10

	// 0xc000124018
	data_stack_int = 111
	data_global_i = +data_stack_int

	/*
				data_global_i 0x118be68
				data_heap_slice 0xc000113ed8
				data_stack_int 	0xc000124018

				0x118be68 < 0xc000113ed8 < 0xc000124018

		high_addr	0xc000124018|-------|
								|		|
								|stack	|
								--------
								|		|
								|heap	|
					0xc000113ed8---------
								|		|
								|global |
		low_addr:	0x118be68	|-------|

	*/
	fmt.Println("data_global_i", &data_global_i)
	fmt.Println("data_heap_slice", &data_heap_slice[0])
	fmt.Println("data_stack_int", &data_stack_int)
	fmt.Println("data_stack_int", &data_stack_int)
}
