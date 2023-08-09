package main

import (
	"fmt"
	"reflect"
)

type reflectStr struct {
	a int32
	b string
}

func DeepEqual(x, y any) bool {
	if x == nil || y == nil {
		return x == y
	}
	v1 := reflect.ValueOf(x)
	v2 := reflect.ValueOf(y)
	if v1.Type() != v2.Type() {
		return false
	}
	return false
}

// 比较两个结构体类型
func compareStruct() {
	s1 := reflectStr{
		a: 12,
		b: "123",
	}
	s2 := reflectStr{
		a: 13,
		b: "123",
	}
	s3 := reflectStr{
		a: 13,
		b: "123",
	}
	DeepEqual(s1, s2)
	if reflect.DeepEqual(s1, s2) {
		fmt.Println("DeepEqual s1 same with s2")

	} else {
		fmt.Println("DeepEqual s1 not  same with s2")
	}

	if reflect.TypeOf(s1) == reflect.TypeOf(s2) {
		fmt.Println("s1 same with s2")

	} else {
		fmt.Println("s1 not  same with s2")
	}

	if reflect.TypeOf(s3) == reflect.TypeOf(s2) {
		fmt.Println("s3 same with s2")
	} else {
		fmt.Println("s3 not same with s2")
	}

}

func main() {
	compareStruct()
}
