package main

var data_global_i int

func f_data_heap_slice() {
	data_heap_slice := make([]int32, 10, 10)
	data_heap_slice[1] = 1
}

func main() {
	var data_stack_int int

	f_data_heap_slice()

	data_global_i = 10
	data_stack_int = 111
	data_global_i = +data_stack_int

	//fmt.Println("data_global_i", &data_global_i)
	//fmt.Println("data_heap_slice", &data_heap_slice)
	//fmt.Println("data_stack_int", &data_stack_int)

	// 		0x0049 00073  MOVQ    $10, main.data_global_i(SB)
	//     	0x0054 00084  MOVQ    $111, main.data_stack_int+16(SP)
	//      0x0032 00050  MOVQ    DX, main.data_heap_slice+64(SP)

	// go tool compile -S vars.go | grep 'vars.go:3'
}
