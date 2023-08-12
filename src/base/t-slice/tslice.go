package main

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

func stackBySlice() {
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

func main() {
	//test()
	stackBySlice()
}
