package main

import (
	"fmt"
	"math"
	"strconv"
)

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

func main() {
	const (
		B = 1 << (10 * iota)
		KiB
		MiB
		GiB
		TiB
		PiB
		EiB
	)

	// 1<<10
	// 01 10 10000000000
	// 1  2  1024
	// 2^0 2^1

	fmt.Println("2^9:", math.Pow(2, 10))
	decimalToBinary(1)
	decimalToBinary(2)
	decimalToBinary(1024)
	fmt.Println(KiB)

}
