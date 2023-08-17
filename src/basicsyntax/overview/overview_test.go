package main

import (
	"basicsyntax/common"
	"basicsyntax/common/pb"
	"fmt"
	"github.com/golang/protobuf/proto"
	"math/cmplx"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"testing"
	"time"
	"unicode/utf8"
	"unsafe"
)

/*
	======================		变量数据类型		======================
*/
// 变量声明与赋值
func TestVariable(t *testing.T) {
	// 声明三个变量，皆为 bool 类型
	var c, python, java bool

	// 声明不同类型的变量，并且赋值
	var i = true
	var j = 2

	// 复杂变量声明
	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)

	// 短声明变量
	{
		c, python, java := true, false, "no!"
		fmt.Println(c, java, python, i, j, ToBe, MaxInt, z)
	}

	// 声明常量
	const constant = "This is a constant"

	fmt.Println(c, java, python, i, j, ToBe, MaxInt, z, constant)
}

func TestBasicType(t *testing.T) {
	//var b = true
	//common.TestMinMax(b, false)
	//
	//var d_byte byte = 2
	//common.TestMinMax(d_byte, false)
	//
	//var i32 int32 = -2147483648
	//common.TestMinMax(i32, true)

	var iint int = 1
	common.TestMinMax(iint, true)

	var iuint8 uint8 = 2
	common.TestMinMax(iuint8, false)

	//var iint8 int8 = 2
	//common.TestMinMax(iint8, true)

	//{
	//	a, c := 0b1010, 0x64
	//	b := 0o144
	//
	//	fmt.Printf("0b%b, %#o, %#x\n", a, b, c)
	//	fmt.Println(math.MinInt8, math.MaxInt8)
	//}
}

// 引用 todo

// 类型转换 todo

// self define type todo

// 类型比较 reflect.DeepEqual
func TestReflectDeepEqual(t *testing.T) {
	m1 := map[string]int{
		"a": 1,
		"b": 2,
	}
	m2 := map[string]int{
		"a": 1,
		"b": 2,
	}
	m3 := map[string]int{
		"a": 1,
		"b": 3,
	}
	fmt.Println(reflect.DeepEqual(m1, m2))
	fmt.Println(reflect.DeepEqual(m1, m3))
}

/*
	======================		语句		======================
*/
// 选择-[条件判断 if else]
func TestIf(t *testing.T) {
	// 基础形式
	var x, b, c int = 0, 2, 3
	x, _ = strconv.Atoi(os.Args[0])
	b, _ = strconv.Atoi(os.Args[1])
	c, _ = strconv.Atoi(os.Args[2])
	if x > 0 {
		fmt.Println(x)
	} else {
		fmt.Println(x)
	}

	// 条件判断之前添加自定义语句
	if a := b + c; a < 42 {
		fmt.Println(a)
	} else {
		fmt.Println(a - 42)
	}

	// 常用的类型判断
	var val interface{}
	val = 111

	if str, ok := val.(string); ok {
		fmt.Println(str)
	}

	// 利用反射判断变量类型
	tt := reflect.TypeOf(val)
	if tt.String() == "int" {
		fmt.Println("val是 int 类型")
	} else {
		fmt.Println("val不是 int 类型")
	}

}

// 选择-[多路执行 switch]
func TestSwitch(t *testing.T) {
	operatingSystem := runtime.GOOS
	switch operatingSystem {
	case "darwin":
		fmt.Println("Mac OS Hipster")
		// 默认 break，不需要显式声明
	case "linux":
		fmt.Println("Linux Geek")
	default:
		// Windows, BSD, ...
		fmt.Println("Other")
	}

	// 类似于 if，可以在条件之前添加自定义语句
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Mac OS Hipster")
		// 默认 break，不需要显式声明
	case "linux":
		fmt.Println("Linux Geek")
	default:
		// Windows, BSD, ...
		fmt.Println("Other")
	}

	// 使用 switch 语句进行类型判断：
	var anything interface{}
	anything = os.Args[5]
	switch v := anything.(type) {
	case string:
		fmt.Println(v)
	case int32, int64:
		fmt.Println(v)
	default:
		fmt.Println("unknown")
	}

	// 比较，条件运算
	number := 42
	number, _ = strconv.Atoi(os.Args[0])
	switch {
	case number < 42:
		fmt.Println("Smaller")
	case number == 42:
		fmt.Println("Equal")
	case number > 42:
		fmt.Println("Greater")
	}

	// 多条件匹配
	var char byte = '?'
	switch char {
	case ' ', '?', '&', '=', '#', '+', '%':
		fmt.Println("Should escape")
	}
}

// 循环-for / [支持for，没有while/until]
func TestLoop(t *testing.T) {
	for i := 1; i < 10; i++ {
	}

	// while - loop
	i := 1
	for i < 10 {
		i++
	}
	fmt.Println(i)

	// 阻塞，单条件情况下可以忽略分号
	for i < 11 {
		fmt.Println(i)
	}

	// 永远阻塞
	// select{} 替代
	for {
	}
}

// 循环-range / [遍历切片，map]
func TestRange(t *testing.T) {
	// loop over an array/a slice
	slice := []int{1, 23, 34, 11, 1, 1123, 113}
	arr := [10]int{1, 23, 34, 11, 1, 1123, 113}
	common.GetVarType("slice", slice)
	common.GetVarType("arr", arr)

	fmt.Println("== slice index ,value:")
	for i, v := range slice {
		// i 表示下标，e 表示元素
		fmt.Printf("  %d %d\n", i, v)
	}

	// 仅需要元素
	fmt.Println("== slice value:")
	for _, e := range slice {
		// e is the element
		fmt.Printf("  %d\n", e)
	}

	// 或者仅需要下标
	fmt.Println("== slice index:")
	for i := range slice {
		fmt.Printf("   %d\n", i)
	}

	fmt.Println()

	// map
	m1 := make(map[int]int32)
	m1[1] = 20
	m1[2] = 30
	common.GetVarType("m1", m1)

	fmt.Println("== m1 k,v:")
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	// 定时执行
	for range time.Tick(time.Second) {
		if time.Now().Unix() < 1691981967 {
			fmt.Println(time.Now().Unix()) //1691981807
		} else {
			break
		}
	}
}

// 跳转-goto
func TestGoto(t *testing.T) {
	common.TestOuter()
	/*
		break: 终止 switch/for/select
		continue：仅仅用于for，进入下一层循环
	*/
	common.TestContinue()
}

/*
	======================		函数		======================
*/

// 匿名函数 todo

func TestFunction(t *testing.T) {
	// 无参
	common.Function1()

	// 多参数
	common.Function2("function2", 1)

	// 不定参
	common.Function3(10, "11", 1090.0001, []int{10, 20})

	// 函数变量，返回值
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println("add(1,2)=", add(1, 2))

	// todo 切片，map，channel  作为函数参数传递会怎么样呢

	// 1.函数作为参数，执行回调
	/*
		var fn func(int) bool
		fn = func(i int) bool {
			if i < 15 {
				fmt.Printf("%d lee then 15\n", i)
			}
			return i < 15
		}
		fmt.Println(Filter([]int{10, 20}, fn))
	*/

	//  2.简化写法[函数作为参数，执行回调]
	fmt.Println(common.Filter([]int{10, 20}, func(i int) bool {
		if i < 15 {
			fmt.Printf("%d lee then 15\n", i)
		}
		return i < 15
	}))

	// 多返回值 todo

	// 闭包
	f, d := common.Scope()
	fmt.Println("Closure 使用:", f(), d)

}

// 延迟调用 defer
func TestDefer(t *testing.T) {
	defer common.TestDefer()

	var i int
	for range time.Tick(time.Second) {
		for {
			if time.Now().Unix() < time.Now().Unix()+10 && i < 5 {
				i++
				fmt.Println(i)
				fmt.Println(time.Now().Unix()) //1691981807
			} else {
				fmt.Println(time.Now().Unix())
				break
			}
		}
		break
	}
}

// 错误处理 error
func TestError(t *testing.T) {
	result, err := common.GetNumber(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

// panic
func TestPanic(t *testing.T) {
	result, err := common.GetNumber(11)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(result)
	}

}

/*
	======================		数据		======================
*/
// type rename, HPPCb  means handle callback function
// and HPPCb is a type of function
type HPPCb func(msg string, p proto.Message) bool
type ProtoMgr struct {
	ProtoMap map[string]HPPCb
}

// initialize ProtoMgr
func (self *ProtoMgr) Init(param ...interface{}) bool {
	self.ProtoMap = make(map[string]HPPCb, 0)
	// do something else
	return true
}
func (self *ProtoMgr) sayHello(protocol string, data proto.Message) bool {
	fmt.Printf("protocol:%v data:%v\n", protocol, data)
	fmt.Println("say hello function called!")
	return true
}

func (self *ProtoMgr) RegisterHPPCb(proto string, handler HPPCb) bool {
	self.ProtoMap[proto] = handler
	return true
}
func TestTypeDefine(t *testing.T) {
	pm := ProtoMgr{}
	if !pm.Init() {
		return
	}
	pm.RegisterHPPCb("sayhello", pm.sayHello)

	time.Sleep(time.Second)

	handler := pm.ProtoMap["sayhello"]
	handler("111", &pb.Pb{})

}

// type conversion and type checking
func TestCvAndCk(t *testing.T) {
	var data interface{} = 12

	// get data type
	common.GetVarType("dataName", data)
	str, ok := data.(string) // 判断myInterface是否为string类型
	if ok {
		fmt.Println("myInterface is a string:", str)
	} else {
		fmt.Println("myInterface is not a string")
	}

}

func TestString(t *testing.T) {
	/*
		string:

			data struct string mem struct is like this:

			 +-----------+          +---+---+---+---+---+
			 |  pointer -|--------> | h | e | l | l | o |
			 +-----------+          +---+---+---+---+---+
			 |  len = 5  |
			 +-----------+          [...]byte, UTF-8
				header

				// runtime/string.go
				type stringStruct struct {
					str unsafe.Pointer     // str *int   8 byte
					len int
				}

				type stringStructDWARF struct {
					str *byte				// str *uint8  1byte
					len int
				}

				int  	字节大小:8  最小值:-9223372036854775808  最大值:9223372036854775807  2^(64-1)-1:  	64位有符号整数的数据范围: -9223372036854775808 到 9223372036854775807
				uint8  	字节大小:1  最小值:0  					  最大值:255  				 2^8:  			8位无符号整数的数据范围: 0 到 255

				编码 UTF-8，无 NULL 结尾，默认值 ""。
				使用 `raw string` 定义原始字符串。
				支持 !=、==、<、<=、>=、>、+、+=。
				索引访问字节数组（非字符），不能获取元素地址。
				切片返回子串，依旧指向原数组。
				内置函数 len 返回字节数组长度。

		-
	*/

	{
		/*
			"雨痕"：这是一个包含中文字符的普通字符串，没有使用任何转义。
			"\x61"：这是一个ASCII转义序列，表示小写字母'a'的十六进制编码，它在字符串中的位置对应'a'字符。
			"\142"：这是一个八进制转义序列，表示小写字母'b'的八进制编码，它在字符串中的位置对应'b'字符。
			"\u0041"：这是一个Unicode转义序列，表示大写字母'A'的Unicode码值，它在字符串中的位置对应'A'字符。
		*/

		s := "雨痕\x61\142\u0041"

		bs := []byte(s)
		rs := []rune(s) // rune/int32: unicode code point
		fmt.Println(s)
		fmt.Printf("% X, %d\n", s, len(s))
		fmt.Printf("% X, %d\n", bs, utf8.RuneCount(bs))
		fmt.Printf("%U, %d\n", rs, utf8.RuneCountInString(s))
	}

	{
		s := "雨痕abc"

		fmt.Printf("%X\n", s[1]) // 9B

		// println(&s[1])
		// invalid operation: cannot take address of s[1]
	}

	{
		var s string
		println(s == "") // true

		// println(s == nil)
		// invalid operation: mismatched types string and untyped nil
	}

	{
		s := `line\r\n,
  line 2`

		println(s) // raw string
	}

	{
		s := "ab" + // 跨行时，加法操作符必须在上行结尾。
			"cd"

		println(s == "abcd") // true
		println(s > "abc")   // true
	}

	{
		s := "hello, world!"
		s2 := s[:4]

		common.GetStringHeader(s)
		common.GetStringHeader(s2)
		p1 := (*reflect.StringHeader)(unsafe.Pointer(&s))
		p2 := (*reflect.StringHeader)(unsafe.Pointer(&s2))

		fmt.Printf("%#v, %#v\n", p1, p2)
		//&reflect.StringHeader{Data:0x1232a51, Len:13},
		//&reflect.StringHeader{Data:0x1232a51, Len:4}
	}

	// 遍历
	{
		s := "雨痕"

		// byte
		for i := 0; i < len(s); i++ {
			fmt.Printf("%d: %X\n", i, s[i])
		}

		// rune
		for i, c := range s {
			fmt.Printf("%d: %U\n", i, c)
		}
	}

	// slice operator copy append
	{
		s := "de"

		bs := make([]byte, 0)
		bs = append(bs, "abc"...)
		bs = append(bs, s...)

		buf := make([]byte, 5)
		copy(buf, "abc")
		copy(buf[3:], s)

		fmt.Printf("%s\n", bs)  // abcde
		fmt.Printf("%s\n", buf) // abcde
	}

	// 标准库相关 todo
	// 转换
	{

	}

}

func TestArray(t *testing.T) {
	/*
		array:

			+---+---+---+----//---+----+
			| 0 | 1 | 2 | ... ... | 99 |   [100]int
			+---+---+---+----//---+----+

		-
	*/

	{
		var d1 [2]int // 编译期计算。
		var d2 [2]int // 元素类型相同，长度不同！

		d1 = d2
		fmt.Println(d1)
		//~~ cannot use [2]int as type [8]int in assignment
	}

	{ // 初始化
		var a [4]int // 元素自动初始化为零。

		b := [4]int{2, 5}         // 未提供初始值的元素自动初始化为 0。
		c := [4]int{5, 3: 10}     // 可指定索引位置初始化。
		d := [...]int{1, 2, 3}    // 编译器按初始化值数量确定数组长度。
		e := [...]int{10, 3: 100} // 支持索引初始化，数组长度与此有关。

		fmt.Println(a, b, c, d, e)

		type user struct {
			id   int
			name string
		}

		uObj := [...]user{ // 这里不能省略。
			{1, "zs"}, // 元素类型标签可省略。
			{2, "ls"},
		}
		fmt.Println(uObj)
		common.GetVarType("uObj", uObj)

		// 多数组初始化
		x := [...][2]int{
			{1, 2},
			{3, 4},
		}

		y := [...][3][2]int{
			{
				{1, 2},
				{3, 4},
				{5, 6},
			},
			{
				{10, 11},
				{12, 13},
				{14, 15},
			},
		}

		fmt.Println(x, len(x), cap(x)) // 2, 2
		fmt.Println(y, len(y), cap(y)) // 2, 2
	}

	{ // if support element compare then the array is always support
		var a, b [2]int
		println(a == b) // true

		c := [2]int{1, 2}
		d := [2]int{0, 1}
		println(c == d) // false
	}

	{
		/*
			array pointer:
					&array,point to the entire array's pointer
			element's pointer:
					&array[0],the pointer point to the array's element
			pointer array
					[...]*int{&a,&b}

		*/
		d := [...]int{0, 1, 2, 3}

		var p *[4]int = &d  // 数组指针。
		var pe *int = &d[1] // 元素指针。

		p[0] += 10 // 相当于 (*p)[0]
		*pe += 20

		fmt.Println(d)

		{
			a, b := 1, 2

			d := [...]*int{&a, &b} // 指针数组。
			*d[1] += 10            // d[1] 返回 b 指针。

			fmt.Println(d)
			fmt.Println(a, b)
		}
	}

	common.ArrayOpt()

}

func TestSlice(t *testing.T) {
	/*
		slice:

				  +---------+            +---+---+----//---+----+
				  |  array -|----------> | 0 | 1 | ... ... | 99 |
				  +---------+            +---+---+----//---+----+
				  |  len    |
				  +---------+            array
				  |  cap    |
				  +---------+
			     	header

		-

				// runtime/slice.go
				type slice struct {
					array unsafe.Pointer
					len   int
					cap   int
				}

			s := a[2:6:8]
			+---+---+---+---+---+---+---+---+---+---+
			| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 |   array
			+---+---+---+---+---+---+---+---+---+---+
					|               |       |           slice: [low : high : max]
					|<--- s.len --->|       |
					|                       |           len = high - low
					|<------- s.cap ------->|           cap = max  - low
	*/

	//slice 内存结构
	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} // array

	s := a[2:6:8]
	common.GetVarType("S", s)

	fmt.Println(s)
	fmt.Println(len(s), cap(s))

	// 引用原数组。

	fmt.Printf("a: %p ~ %p\n", &a[0], &a[len(a)-1])
	fmt.Printf("s: %p ~ %p\n", &s[0], &s[len(s)-1])

	for i, x := range s {
		fmt.Printf("s[%d]: %d\n", i, x)
	}

	/*

		# 切片引用原数组，此图只为方便理解。
		+---+---+---+---+---+---+---+---+---+---+
		| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 |   a: [10]int
		+---+---+---+---+---+---+---+---+---+---+
		        .                       .
		        +---+---+---+---+---+---+
		        | 2 | 3 | 4 | 5 |   |   |           s: a[2:6:8]
		        +---+---+---+---+---+---+
		        0   1   2   3


		s[0]: 2
		s[1]: 3
		s[2]: 4
		s[3]: 5
	*/

	// len range 范围， cap 截取范围

	{ //构造slice
		a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		s := a[2:5:7]

		fmt.Println(s, len(s), cap(s)) //

		// 按初始化值，自动分配底层数组。
		s1 := []int{0, 1, 2, 3}
		fmt.Println(s1, len(s1), cap(s1))

		// 自动创建底层数组。
		s2 := make([]int, 5)
		fmt.Println(s2, len(s2), cap(s2))

		s3 := make([]int, 0, 5)
		fmt.Println(s3, len(s3), cap(s3))

		/*
			s0: [2 3 4],	 len= 3, cap = 5
			s1: [0 1 2 3],   len = 4, cap = 4
			s2: [0 0 0 0 0], len = 5, cap = 5
			s3: [],          len = 0, cap = 5
		*/
	}

	{ // slice 初始化
		p := func(s []int) {
			fmt.Printf("%t, %d, %#v\n",
				s == nil,
				unsafe.Sizeof(s),
				(*reflect.SliceHeader)(unsafe.Pointer(&s)))
		}
		// 仅分配 header 内存，未初始化。
		var s1 []int
		// s1[0] = 111   it wil be panic
		//s1 = append(s1, 111)

		// 初始化。
		s2 := []int{}
		//s2 = append(s2, 111)

		// 调用 makeslice 初始化。
		s3 := make([]int, 0)

		p(s1)
		p(s2)
		p(s3)

		/*
			true,  24, &reflect.SliceHeader{Data:0x0, Len:0, Cap:0}
			false, 24, &reflect.SliceHeader{Data:0x14b4d20, Len:0, Cap:0}
			false, 24, &reflect.SliceHeader{Data:0x14b4d20, Len:0, Cap:0}
		*/
	}

	{ // slice can convert to array/array pointer directly

		/*
			a 数组，s切片 两个数据都是指向同一块内存地址，引用
		*/
		var a [4]int = [...]int{0, 1, 2, 3}

		// array -> slice: 指向原数组
		//var s []int = a[:]
		s := a[:]

		//a 是 array 类型
		//s a[:]是 slice 类型
		common.GetVarType("a", a)
		common.GetVarType("s a[:]", s)
		println(&s[0] == &a[0]) // true

		/*
			a2 数组，s转换数组，值拷贝
		*/
		// slice -> array: 复制底层数组（片段）
		a2 := [4]int(s)
		println(&a2[0] == &a[0]) // false

		/*
			slice -> *array
				var i *int  // 0xc000020450
				b := 10
				i = &b

				var dd *[4]int
				dd = (*[4]int)(s)
				fmt.Println(dd == &a) // true
		*/
		// slice -> *array: 返回底层数组（片段）指针
		p2 := (*[4]int)(s)

		println(p2 == &a) // true
	}

}

func TestSlicePointer(t *testing.T) {
	a := [...]int{0, 1, 2, 3}
	s := a[:]
	s1 := a[0:2:3]

	/*
		the first element of a's address: 0xc00001e100
		s,false, 24, &reflect.SliceHeader{Data:0xc00001e100, Len:4, Cap:4}
		s1,false, 24, &reflect.SliceHeader{Data:0xc00001e100, Len:2, Cap:3}
		-
	*/
	p := &s    // 切片指针
	e := &s[1] // 元素指针
	fmt.Println("the first element of a's address:", &a[0])
	common.GetSliceHeader("s", s)
	common.GetSliceHeader("s1", s1)

	/*
			Go has pointers. A pointer holds the memory address of a value.
				内存 地址和指针 的问题
					地址：内存中每个字节单位的编号
					指针：指针是实体，需要分配内存空间，专门用来保持地址的整形变量
				p 是一个指针变量，
						1.有内存地址，地址为 0xc0000100d8
						2.存储s切片header/也就是s切片底层数组a的首元素地址

													p指针变量内存
			9									  	+---------------------+
			9                              	 		|0xc0000100d8(内存编号)|
			9									  	+---------------------+
			9							 p指针    	0xc000012058

							slice
		 	9 			  0xc0000100d8
				          +---------+            +---+---+----//---+----+
						  |  array -|----------> | 0 | 1 | ... ... | 99 |
						  +---------+            +---+---+----//---+----+
						  |  len    |			 0xc00001e100
						  +---------+            	array
						  |  cap    |
						  +---------+
					     	header

				s切片的地址：0xc0000100d8
				s:[0 1 2 3]
				&s:&[0 1 2 3]
				*(&s):[0 1 2 3]


				*p,false, 24, &reflect.SliceHeader{Data:0xc00001e100, Len:4, Cap:4}
				p:0xc0000100d8
				*p:0xc00001e100
				&(*p)[0]:0xc00001e100
				&p:0xc000012058
				*(&p):0xc0000100d8
				e:0xc00001e108

								数组指针直接指向元素所在内存。
					切片指针指向 header 内存,也就是底层数组首元素地址 。
	*/
	fmt.Printf("s切片的地址：%p\n", &s)
	fmt.Printf("s:%v \n", s)
	fmt.Printf("&s:%v \n", &s)
	fmt.Printf("*(&s):%v \n\n", *(&s))

	// p切片指针
	common.GetSliceHeader("*p", *p)
	fmt.Printf("p:%p \n", p) // 也就是切片s的地址 0xc0000100d8
	fmt.Printf("*p:%p \n", *p)
	fmt.Printf("&(*p)[0]:%v \n", &(*p)[0])
	fmt.Printf("&p:%p \n", &p) //  指针变量 p 自身的地址 0xc000012058
	fmt.Printf("*(&p):%p \n", *(&p))

	fmt.Printf("e:%v \n", e)
	// _ = p[1]
	//     ~~~~ invalid: cannot index p (variable of type *[]int)

	_ = (*p)[1]

	// 元素指针指向数组。

	*e += 100

	fmt.Println(e == &a[1]) // true
	fmt.Println(a)          // [0 101 2 3]
}

// 映射类型 字典
func TestSliceOther(t *testing.T) {
	var a [4]int = [...]int{0, 1, 2, 3}

	// 基于数组指针创建切片。
	//var p *[4]int = &a
	//var s []int = p[:]

	s := a[:]

	println(&s[2] == &a[2]) // true

	// 基于非数组指针创建切片
	p2 := (*byte)(unsafe.Pointer(&a[0])) // 元素指针
	var s2 []byte = unsafe.Slice(p2, 8)

	fmt.Println(s2)

	// 交错数组功能
	s3 := [][]int{
		{1, 2},
		{10, 20, 30},
		{100},
	}
	common.GetVarType("s3", s3)

	// 切片作为参数
	common.SliceAsFunctionParam()

	// slice append
	common.TestSliceAppend()
}

// 映射类型 字典
func TestMap(t *testing.T) {
	/*
			   pointer              header

		  +-------+          +-----------+
		  |  map -|--------> |  hmap     |
		  +-------+          +-----------+
		                     |  ...      |
		                     +-----------+       +-----//-----+
		                     |  buckets -|-----> | ...    ... |   array
		                     +-----------+       +-----//-----+

	*/

	// 创建map方式
	m := make(map[string]int) // make
	m1 := map[string]int{}    // 初始化空map
	m4 := map[int]string{     // 使用字面量初始化
		1: "value1",
	}
	fmt.Println(m4)

	var m3 map[string]int // nil
	if nil == m3 {
		fmt.Println("m3 is nil")
	} else {
		fmt.Println("m3 not nil")
	}

	m1["age"] = 10086
	m["dog"] = 13
	m["salary"] = 100231

	// 遍历
	fmt.Println("m1 数据状态")
	for k, v := range m1 {
		fmt.Printf("k:%s ,value:%d \n", k, v)
	}

	//增，改
	m["1"] = 1
	m["1"] = 1111

	fmt.Println("m 数据状态")
	for k, v := range m {
		fmt.Printf("k:%s ,value:%d \n", k, v)
	}
	// 查
	v, f := m["ag1e"]
	if !f {
		fmt.Println("data not found")
	} else {
		fmt.Println(v)
	}

	// 删除
	delete(m, "1")

}

// 结构体
func TestStruct(t *testing.T) {
	type node struct {
		id   int    `field:"uid"  type:"integer"`
		name string `field:"name" type:"text"`
		next *node  // 自身指针类型。
	}

	n := node{
		id:   2,
		name: "abc", // 注意结尾逗号 !!!
	}

	n.next = &n
	fmt.Println(n.next.name)

	{
		type User struct {
			id   int    `field:"uid"  type:"integer"`
			name string `field:"name" type:"text"`
		}

		t := reflect.TypeOf(User{})

		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			fmt.Println(f.Name, f.Tag.Get("field"), f.Tag.Get("type"))
		}
	}

	// 内存对齐 todo

}

// 指针 todo
func TestPointer(t *testing.T) {
	/*
			Go has pointers. A pointer holds the memory address of a value.
					内存 地址和指针 的问题
						地址：内存中每个字节单位的编号
						指针：指针是实体，需要分配内存空间，专门用来保持地址的整形变量
					p 是一个指针变量，
							1.有内存地址，地址为 0xc0000100d8
							2.存储s切片header/也就是s切片底层数组a的首元素地址
				-


		            x: int          p: *int
		   -------+---------------+--------------+-------
		     ...  | 100           | 0xc000000000 |  ...    memory
		   -------+---------------+--------------+-------
		          0xc000000000    0xc000000008             address

	*/
	// 空指针也会分配内存。
	var p *int
	println(unsafe.Sizeof(p)) // 8

	/*

	   p: *int
	   +--------------+
	   | 0            |
	   +--------------+
	   0xc000000008

	*/

	var x int

	/*

	   p: *int                x: int
	   +---------------+      +---------------+
	   | 0             |      | 0             |
	   +---------------+      +---------------+
	   0xc000000008           0xc000000000

	*/

	// 二级指针，指针的指针。
	var pp **int = &p
	*pp = &x

	/*

	                               +--------- *p -----------+
	                               |                        |
	   pp: **int              p    |                 x      v
	   +---------------+      +---------------+      +---------------+
	   | 0xc000000008 -|----->| 0xc000000000 -|----->| 0             |
	   +---------------+      +---------------+      +---------------+
	   0xc000000010           0xc000000008   ^       0xc000000000   ^
	         |                               |                      |
	         +-------- *pp ------------------+                      |
	         |                                                      |
	         +-------- **pp ----------------------------------------+
	*/

	*p = 100
	**pp += 1

	println(**pp, *p, x) // 101, 101, 101

}

/*
	======================		method		======================
	对象实例绑定的特殊函数

-
*/
func TestMethod(t *testing.T) { // 在切片中使用指针
	common.StackBySlice()
}

/*
	======================		interface		======================

-
*/
func TestInterface(t *testing.T) {
	common.Test()
}

/*
	======================		generic		======================

-
*/
func TestGeneric(t *testing.T) {}

/*
	======================		concurrency		======================
	goroutine
	channel
	sync

-
*/

func tt(x int) {
	println(x)
}
func TestConcurrency(t *testing.T) {
	// 打包并发任务（函数 + 参数）。
	// 并非立即执行。

	go println("abc")
	go tt(123)
	go func(x int) { println(x) }(123)

	// 上述并发任务会被其他饥饿线程取走。
	// 执行时间未知，下面这行可能先输出。

	println("main")
	time.Sleep(time.Second)
}

/*
	======================		package		======================

-
*/
func TestPackage(t *testing.T) {}

/*
	======================		ASM/CGO		======================

-
*/
func TestASM(t *testing.T) {}
func TestCGO(t *testing.T) {}

/*
	======================		testing		======================

-
*/
// unit testing
func TestUintTesting(t *testing.T) {}

// bench testing
func TestBenchTesting(t *testing.T) {}
