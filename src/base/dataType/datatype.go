package main

import (
	"fmt"
	"math"
	"reflect"
)

/*
		==== type ==== | = len = | = default = | ====== comment =============

	 bool               1         false
	 byte               1           0         uint8

	 int,   uint        8           0         x86:4, x64:8
	 int8,  uint8       1           0         -128 ~ 127,     0 ~ 255
	 int16, uint16      2           0         -32768 ~ 32767, 0 ~ 65535
	 int32, uint32      4           0
	 int64, uint64      8           0

	 float32            4          0.0
	 float64            8          0.0

	 complex64          8
	 complex128        16

	 rune               4           0         unicode code point, int32
	 uintptr            8           0         uint

	 string                        ""         len()
	 array                                    len() == cap()
	 struct

	 function                      nil
	 interface                     nil

	 map                           nil        make(), len()
	 slice                         nil        make(), len(), cap()
	 channel                       nil        make(), len(), cap()

*/

func calculateMaxValue(nbit int, isUnsigned bool) {
	if !isUnsigned {
		// 计算有符号整数的数据范围
		minSigned := -int(math.Pow(2, float64(nbit-1)))
		maxSigned := int(math.Pow(2, float64(nbit-1))) - 1
		fmt.Printf("  %d位有符号整数的数据范围: %d 到 %d\n", nbit, minSigned, maxSigned)
	} else {
		// 计算无符号整数的数据范围
		maxUnsigned := uint(math.Pow(2, float64(nbit))) - 1
		fmt.Printf("  %d位无符号整数的数据范围: 0 到 %d\n", nbit, maxUnsigned)

	}
}

func testMinMax(i interface{}, isUnsigned bool) {
	size, minValue, maxValue := getTypeDetails(i)
	fmt.Print(reflect.ValueOf(i).Kind())
	fmt.Print("  字节大小:", size)
	fmt.Print("  最小值:", minValue)
	fmt.Print("  最大值:", maxValue)
	if 0 != size && isUnsigned {
		fmt.Printf("  2^%d:", size*8)
	} else {
		fmt.Printf("  2^(%d-1)-1:", size*8)
	}

	calculateMaxValue(size*8, isUnsigned)

}

func getTypeDetails(i interface{}) (int, interface{}, interface{}) {
	value := reflect.ValueOf(i)
	kind := value.Kind()

	size := typeSize(value.Type())
	switch kind {
	case reflect.Int32:
		minValue := int32(math.MinInt32)
		maxValue := int32(math.MaxInt32)
		return size, minValue, maxValue
	case reflect.Bool:
		return size, "bool范围为true/false", "bool 范围为 true/false"

	case reflect.Uint8:
		minValue := byte(0)
		maxValue := byte(math.MaxUint8)
		return size, minValue, maxValue
	case reflect.Int:
		minValue := math.MinInt
		maxValue := math.MaxInt
		return size, minValue, maxValue
	default:
		return -1, nil, nil
	}
}

func typeSize(typ reflect.Type) int {
	size := typ.Size()
	return int(size)
}

func main() {

	//var b bool = true
	//testMinMax(b, true)
	//
	//var d_byte byte = 2
	//testMinMax(d_byte, true)
	//
	//var i32 int32 = -2147483648
	//testMinMax(i32, false)

	var iint int = 1
	testMinMax(iint, false)

}
