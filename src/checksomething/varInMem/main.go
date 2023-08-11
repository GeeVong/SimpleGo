package main

import (
	"fmt"
	"sort"
	"strconv"
)

/*
	. 验证
		int（stack）
		make slice（heap）
		global param
	在内存中地址高低分布情况
*/

var a [10]int32 // 全局数据
var addresses []addressInfo

type addressInfo struct {
	name          string
	address       string
	address10Data int64
}

func init() {
	addresses = make([]addressInfo, 0, 10) // 地址比较排序，用于查看不同类型变量，内存高低
}
func getGlobalIntAddress() {
	a[0] = 10
	fmt.Printf("globel array's addres:%v\n", &a[0])

	addressCollect("global int", &a[0])
}

func getHeapMapAddress() {
	HeapMap := make([]int32, 10, 10)
	HeapMap[0] = 32
	HeapMap[1] = 32

	fmt.Printf("HeapMap[0]'s addres:%v\n", &HeapMap[0])
	addressCollect("HeapMap", &HeapMap[0])
}

func getHeapInt64Address() {
	anewint64 := new(int64)
	var tanewint64 int64 = 100
	anewint64 = &tanewint64
	fmt.Printf("anewint64's addres:%v\n", &anewint64)
	addressCollect("new(int64)", &anewint64)
}

func convertAddress(addresses []addressInfo) {
	for i := 0; i < len(addresses); i++ {
		addresses[i].address10Data = _16to10(addresses[i].name, addresses[i].address)
	}
}

func _16to10(name, hexString string) int64 {
	decimal, err := strconv.ParseInt(hexString, 0, 64)
	if err != nil {
		fmt.Println("转换出错:", err)
		return -1
	}
	fmt.Printf(name)
	fmt.Printf("相关数据,十六进制: %v, 十进制：%d\n", hexString, decimal)
	return decimal
}

func getLocalStrAddress() {
	var str2 string
	str2 = "22221111"
	fmt.Printf("str2's addres:%v\n", &str2)

	addressCollect("str2", &str2)

	var str string
	str = "1111"
	fmt.Printf("str's addres:%v\n", &str)
	addressCollect("str", &str)

}

// 收集不同变量地址
func addressCollect(name string, data interface{}) {
	addressStr := fmt.Sprintf("%p", data)
	addresses = append(addresses, addressInfo{name: name, address: addressStr})
}

func compareAddress() {
	// 实现排序函数
	sort.Slice(addresses, func(i, j int) bool {
		return addresses[i].address10Data < addresses[j].address10Data
	})

	// 打印排序后的内存地址
	fmt.Println("排序后的数据")
	for _, addr := range addresses {
		fmt.Printf("%v\n", addr)
	}

}

func compare16Var(str1, str2 string) {
	// 将十六进制字符串解析为整数
	num1, _ := strconv.ParseInt(str1, 16, 64)
	num2, _ := strconv.ParseInt(str2, 16, 64)

	if num1 > num2 {
		fmt.Printf("%s 更大\n", str1)
	} else if num1 < num2 {
		fmt.Printf("%s 更大\n", str2)
	} else {
		fmt.Println("两个数相等")
	}
}

func main() {
	// 测试 slice int global 变量在内存地址的位置
	fmt.Println("=========== compareAddress =============")
	getLocalStrAddress()
	getGlobalIntAddress()
	getHeapInt64Address()
	getHeapMapAddress()
	convertAddress(addresses)

	compareAddress()

	for {
		select {
		default:

		}
	}

}
