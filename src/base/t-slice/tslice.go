package main

import (
	"fmt"
	"unsafe"
)

var a [10]int32
var addresses []uintptr

func init() {
	addresses = make([]uintptr, 0, 10)
}

func main() {
	// 切片作为参数
	sparam := make([]int32, 0, 10)
	sparam = append(sparam, 1)
	sparam = append(sparam, 2)

	sparam = sliceParam(sparam)
	fmt.Println("====")
	println("sparam addrres\n:", &sparam[0])
	showSliceParam(sparam)

	//aaa := 0xc0000b8018
	//fmt.Println(aaa)

	// 测试 slice int global 变量在内存地址的位置
	tSlice()
	a[0] = 10
	fmt.Printf("globel array's addres:%v\n", &a[0])
	unsafePtr := unsafe.Pointer(&a[0])
	addressCollect(uintptr(unsafePtr))
	fmt.Println("============")
	compareAddress()

	fmt.Println(addresses)
	for {
		select {
		default:

		}
	}

}

func sliceParam(s []int32) []int32 {
	s[0] = 1111
	s = append(s, 100)

	println("s addrres\n:", s)
	showSliceParam(s)
	return s
}

func showSliceParam(s []int32) {
	for i := 0; i < len(s); i++ {
		fmt.Println("showSliceParam：", s[i])
	}
}

/*
	heap：anewint64's addres:0xc0000b8018
	stack：str's addres:0xc00008c220
	globel：globel array's addres:0x117a140


	heap：824634474520
	stack：824634294816
	globel：18325824

*/

func tSlice() {
	//s := make([]int32, 10, 111)
	//s[0] = 1
	//s[1] = 1
	//s[2] = 1
	//fmt.Printf("slice's addres:%v\n", &s[0])
	//
	//unsafePtr := unsafe.Pointer(&s[0])
	//addressCollect(uintptr(unsafePtr))

	anewint64 := new(int64)
	var tanewint64 int64 = 100
	anewint64 = &tanewint64
	fmt.Printf("anewint64's addres:%v\n", &anewint64)
	addressCollect(uintptr(unsafe.Pointer(anewint64)))

	var str string
	str = "1111"
	fmt.Printf("str's addres:%v\n", &str)

	addressCollect(uintptr(unsafe.Pointer(&str)))
}

func addressCollect(a uintptr) {
	addresses = append(addresses, a)
}

// todo 存在bug
func compareAddresses(addresses []uintptr) {
	for i, address := range addresses {
		fmt.Printf("Address %d: %x\n", i, address)
	}

	if len(addresses) < 2 {
		return
	}

	smallest := addresses[0]
	largest := addresses[0]

	for _, address := range addresses {
		if address < smallest {
			smallest = address
		}
		if address > largest {
			largest = address
		}
	}

	fmt.Printf("Smallest address: %x\n", smallest)
	fmt.Printf("Largest address: %x\n", largest)
}

func compareAddress() {
	for i := 0; i < len(addresses); i++ {
		fmt.Println("排序前数据：", addresses[i])
	}

	compareAddresses(addresses)

}
