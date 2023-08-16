package common

import (
	"errors"
	"fmt"
	"reflect"
	"unsafe"
)

// todo 类型待补充
func GetVarType(varName string, data interface{}) {
	tt := reflect.TypeOf(data)
	switch tt.Kind() {
	case reflect.Slice:
		fmt.Println(varName, "是 slice 类型")
	case reflect.Array:
		fmt.Println(varName, "是 array 类型")
	case reflect.Map:
		fmt.Println(varName, "是 map 类型")
	case reflect.String:
		fmt.Println(varName, "是 str 类型")
	case reflect.Interface:
		fmt.Println(varName, "是 Interface 类型")
	case reflect.Int:
		fmt.Println(varName, "是 int 类型")
	}
}

// 简单函数定义
func Function1() {
	fmt.Println("define a func called function1")
}

// 含参函数定义
func Function2(param1 string, param2 int) {
	fmt.Println("define a func called function2 with params")
}

// 不定参
func Function3(param1 int, param ...interface{}) int {
	fmt.Println("define a func called function2 with params")

	fmt.Println(param...)
	return param1
}

func TestDefer() {
	fmt.Println("func called testDefer")
}

// 回调
func Filter(s []int, fn func(int) bool) []int {
	var p []int // == nil
	for _, v := range s {
		if fn(v) {
			p = append(p, v)
		}
	}
	return p
}

func Scope() (func() int, int) {
	outer_var := 2
	foo := func() int {
		if outer_var < 5 {
			outer_var++
		} else {
			return outer_var
		}
		fmt.Println("outer_var1:", outer_var)
		return outer_var
	}
	fmt.Println("outer_var2:", outer_var)
	return foo, outer_var
}

// test error/panic
func GetNumber(num int32) (int32, error) {
	arr := [5]int32{1, 23, 41}
	for _, v := range arr {
		if v == num {
			return num, nil
		}
	}
	return -1, errors.New("num is not found")
}

/*
	reflect.SliceHeader 是 reflect 包中定义的结构体，
	它提供了关于底层数组的指针地址、长度和容量等底层信息。

-
*/
func GetSliceHeader(name string, s []int) {
	fmt.Printf("%s,%t, %d, %#v\n",
		name,
		s == nil,
		unsafe.Sizeof(s),
		(*reflect.SliceHeader)(unsafe.Pointer(&s)))
}

func GetStringHeader(s string) {
	fmt.Printf("%t, %d, %#v\n",
		s == "",
		unsafe.Sizeof(s),
		(*reflect.StringHeader)(unsafe.Pointer(&s)))
}
