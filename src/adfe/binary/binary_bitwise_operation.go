package main

import (
	"fmt"
	"math"
	"strconv"
)

/*
	进制转换 / 位运算相关
*/

const (
	B = 1 << (10 * iota)
	KiB
	MiB
	GiB
	TiB
	PiB
	EiB
)

func testBinaryAndDecimal() {
	/* 这里有几个概念：
		1.次方表示10进制数
		2.二进制
		3.10进制
		运算：
			1. 一般运算：二进制 to 次方 to 十进制
			2. 其他进制转换 类比
	 2^0  1  		00000000001
	 2^1  2			00000000010
	 2^2  4			00000000100
	 2^3  8			00000001000
	 2^10 1024		10000000000

	2^(position-1)
	*/

	binaryToDecimal("00000000001")
	decimalToBinary(1024)
	fmt.Println(KiB)

	// 11位 10000000000
	// 1 --》 1 << 10 --》 10000000000
	// 1 		00000000001
	// 1 << 1  	00000000010
	decimalToBinary(1)
	decimalToBinary(1 << 10)

	// 2^10: 1024
	fmt.Println("2^10:", math.Pow(2, 10))

}

// 01 to 123456789
func binaryToDecimal(binary string) {
	decimal, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		return
	}
	fmt.Println(binary, "十进制：", decimal)
}

// 123456789 to  01
func decimalToBinary(decimal int) string {
	binary := ""
	for decimal > 0 {
		remainder := decimal % 2
		binary = strconv.Itoa(remainder) + binary
		decimal = decimal / 2
	}
	fmt.Println(binary)
	return binary
}

/*
逢x进1

	常用的进制有四种：二进制（binary）、八进制（Octal）、十进制（Decimal）和十六进制（Hexadecimal）。

		二进制（binary）是由0和1组成的进制，用于计算机的内部表示和处理信息。
		八进制（Octal）是由0到7组成的进制，每个八进制位表示三个二进制位。
		十进制（Decimal）是我们平常使用的进制，由0到9组成。
		十六进制（Hexadecimal）是由0到9以及字母A到F（代表10到15）组成的进制，每个十六进制位表示四个二进制位。

	ff 1111 1111
	15 = 2^3 +2^2+2^1+2^0 = 8 + 4 +2 +1  1111
	1byte=8bit

	内存中 ff 1字节
*/

func hexToBinary2(hex string) (string, error) {
	decimal, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		return "", err
	}

	binary := strconv.FormatInt(decimal, 2)
	return binary, nil
}

func testHexToBinary(hex string) {
	binary, _ := hexToBinary2(hex)
	fmt.Printf("十六进制数 %s 转换为二进制数是 %s\n", hex, binary)
}

func main() {
	// 二进制十进制
	//testBinaryAndDecimal()

	// 十六 to 二
	// 十六进制数 ff 转换为二进制数是 11111111

	binaryToDecimal("1111")
	testHexToBinary("ff")

}
