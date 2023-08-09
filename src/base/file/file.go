package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	data := make([]byte, 11)
	count, err := file.Read(data)
	if err != nil {
		return
	}
	fmt.Println("字节数据：", data)          // [72 101 108 108 111 32 119 111 114 108 100]
	fmt.Println("字符串数据：", string(data)) // Hello world
	fmt.Println("所获取字节的长度：", count)     // 11
}
