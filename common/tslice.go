package common

import (
	"fmt"
)

type Stack []int

func NewStack() *Stack {
	s := make(Stack, 0, 10)
	return &s
}

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Pop() (int, bool) {
	if len(*s) == 0 {
		return 0, false
	}

	x, n := *s, len(*s)
	/*
		1.x 和 n 的解构赋值：x, n := *s, len(*s)
		*s 是栈对象的切片，通过解引用操作获取栈对象的底层切片。
		len(*s) 获取了栈对象底层切片的长度，即栈中元素的个数。
		x 是一个临时变量，引用了栈对象的底层切片。
		n 是一个变量，保存了栈的长度。

		2.弹出栈顶元素：v := x[n-1]
		n-1 表示栈的顶部元素的索引，我们通过索引操作来获取栈顶元素的值。由于切片索引是从 0 开始的，所以栈顶元素的索引为 n-1。
		x[n-1] 表示获取栈顶元素的值，并将其赋给变量 v。

		3.更新栈对象：*s = x[:n-1]
		x[:n-1] 表示切片操作，从底层切片 x 中截取一个新的切片，包含了除了栈顶元素之外的所有元素。切片操作的语法是 切片[起始索引:结束索引]，如果省略起始索引，则默认为 0；如果省略结束索引，则默认为切片的长度。
		*s = x[:n-1] 表示将截取的新切片赋值给栈对象的切片，即更新了栈对象，将栈顶元素弹出。

	*/
	v := x[n-1]
	*s = x[:n-1]

	return v, true
}

func StackBySlice() {
	s := NewStack()

	// push
	for i := 0; i < 5; i++ {
		s.Push(i + 10)
	}

	// pop
	for i := 0; i < 7; i++ {
		fmt.Println(s.Pop())
	}
}

// ---------------------------
func test() {
	s1 := make([]int, 0, 10)
	s1 = append(s1, 10)
	s1 = append(s1, 110)

	x, n := &s1, len(s1)
	v := (*x)[n-1]
	//*s = x[:n-1]

	fmt.Println(v)
	fmt.Println(s1)
	fmt.Println(&s1[0])
	fmt.Println(*x)
	fmt.Println(&(*x)[0])
}

// --- 切片作为参数

func SliceAsFunctionParam() {
	////  数组产生的切片作为参数
	//a := [...]int{1, 2}
	//s := a[:]
	//
	///*
	//		a's address:0xc000020390
	//		s's address:0xc000010108
	//		a[0]'s address:0xc000020390
	//	-
	//*/
	//fmt.Printf("a's address:%p\n"+
	//	"s's address:%p\n"+
	//	"a[0]'s address:%p\n",
	//	&a,
	//	&s,
	//	&a[0],
	//)
	s := genSliceByMake()
	// s,false, 24, &reflect.SliceHeader{Data:0xc000020390, Len:2, Cap:2}
	GetSliceHeader("s", s)

	fmt.Printf("==== s as param before &s address:%p\n", &s)
	SliceInfo(s)

	sliceParam(s)
	fmt.Println(s)

}

func sliceParam(paramSlice []int) {
	//paramSlice := genSliceByMake()
	//copy(paramSlice, s)
	GetSliceHeader("s", paramSlice)
	fmt.Printf("局部变量 paramSlice  &paramSlice:%p\n", &paramSlice)
	fmt.Printf("局部变量 paramSlice  &paramSlice[0]:%p\n", &paramSlice[0])
	fmt.Printf("局部变量 paramSlice  paramSlice[0]:%d\n", paramSlice[0])

	fmt.Printf("\n")
	paramSlice[0] = 1111
	paramSlice = append(paramSlice, 100)

	fmt.Println("==== s as param after")
	SliceInfo(paramSlice)

	GetSliceHeader("s", paramSlice)
	fmt.Printf("操作之后 ,局部变量 paramSlice  &s:%p\n", &paramSlice)
	fmt.Printf("局部变量 paramSlice  &s[0]:%p\n", &paramSlice[0])
	fmt.Printf("局部变量 paramSlice  s[0]:%d\n", paramSlice[0])
}

func SliceInfo(s []int) {
	for i := 0; i < len(s); i++ {
		fmt.Println("showSliceParam：", s[i])
	}
}

func genSliceByArray() []int {

	/*
		s,false, 24, &reflect.SliceHeader{Data:0xc00010e210, Len:2, Cap:2}
		s,false, 24, &reflect.SliceHeader{Data:0xc00013e080, Len:3, Cap:4}
	*/

	// 扩容，需要拷贝

	a := [...]int{1, 2}
	s := a[:]
	return s
}

func genSliceByMake() []int {
	/*
			s,false, 24, &reflect.SliceHeader{Data:0xc000136000, Len:2, Cap:1000}
			s,false, 24, &reflect.SliceHeader{Data:0xc000136000, Len:3, Cap:1000}
		不需要拷贝
	*/
	s := make([]int, 2, 1000)
	return s
}

func TestSliceAppend() {
	s := genSliceByMake()

	GetSliceHeader("s", s)
	s = append(s, 100)
	GetSliceHeader("s", s)

}
