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
func TestString(t *testing.T) {
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

/*
array:

	+---+---+---+----//---+----+
	| 0 | 1 | 2 | ... ... | 99 |   [100]int
	+---+---+---+----//---+----+

-
*/
func TestArray(t *testing.T) {
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

		uobj := [...]user{ // 这里不能省略。
			{1, "zs"}, // 元素类型标签可省略。
			{2, "ls"},
		}
		fmt.Println(uobj)
		common.GetVarType("uobj", uobj)

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
func TestSlice(t *testing.T) {
	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} // array

	s := a[2:6:8]

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

}

// ### 映射类型
func TestSwitch8(t *testing.T) {}

// ```
// var m map[string]int
// m = make(map[string]int)
// m["key"] = 42
//
// // 删除某个键
// delete(m, "key")
//
// // 测试该键对应的值是否存在
// elem, has_value := m["key"]
//
// // map literal
//
//	var m = map[string]Vertex{
//	   "Bell Labs": {40.68433, -74.39967},
//	   "Google":    {37.42202, -122.08408},
//	}
//
// ```
//
// ## Struct & Interface: 结构体与接口
//
// ### Struct: 结构体
func TestSwitch9(t *testing.T) {}

// Go 语言中并不存在类的概念，只有结构体，结构体可以看做属性的集合，同时可以为其定义方法。
//
// ```
// // 声明结构体
//
//	type Vertex struct {
//	   // 结构体的属性，同样遵循大写导出，小写私有的原则
//	   X, Y int
//	   z bool
//	}
//
// // 也可以声明隐式结构体
//
//	point := struct {
//	   X, Y int
//	}{1, 2}
//
// // 创建结构体实例
// var v = Vertex{1, 2}
//
// // 读取或者设置属性
// v.X = 4;
//
// // 显示声明键
// var v = Vertex{X: 1, Y: 2}
//
// // 声明数组
// var v = []Vertex{{1,2},{5,2},{5,5}}
// ```
//
// 方法的声明也非常简洁，只需要在 func 关键字与函数名之间声明结构体指针即可，该结构体会在不同的方法间进行复制：
//
// ```
//
//	func (v Vertex) Abs() float64 {
//	   return math.Sqrt(v.X*v.X + v.Y*v.Y)
//	}
//
// // Call method
// v.Abs()
// ```
//
// 对于那些需要修改当前结构体对象的方法，则需要传入指针：
//
// ```
//
//	func (v *Vertex) add(n float64) {
//	   v.X += n
//	   v.Y += n
//	}
//
// ```
//
// ```
// var p *Person = new(Person) // pointer of type Person
// ```
//
// ### Pointer: 指针
func TestSwitch10(t *testing.T) {}

// ```
// // p 是 Vertex 类型
// p := Vertex{1, 2}
//
// // q 是指向 Vertex 的指针
// q := &p
//
// // r 同样是指向 Vertex 对象的指针
// r := &Vertex{1, 2}
//
// // 指向 Vertex 结构体对象的指针类型为 *Vertex
// var s *Vertex = new(Vertex)
// ```
//
// ### Interface: 接口
func TestSwitch11(t *testing.T) {}

//
//Go 允许我们通过定义接口的方式来实现多态性：
//
//```
//// 接口声明
//type Awesomizer interface {
//    Awesomize() string
//}
//
//// 结构体并不需要显式实现接口
//type Foo struct {}
//
//// 而是通过实现所有接口规定的方法的方式，来实现接口
//func (foo Foo) Awesomize() string {
//    return "Awesome!"
//}
//```
//
//```
//type Shape interface {
//   area() float64
//}
//
//func getArea(shape Shape) float64 {
//   return shape.area()
//}
//
//type Circle struct {
//   x,y,radius float64
//}
//
//type Rectangle struct {
//   width, height float64
//}
//
//func(circle Circle) area() float64 {
//   return math.Pi * circle.radius * circle.radius
//}
//
//func(rect Rectangle) area() float64 {
//   return rect.width * rect.height
//}
//
//func main() {
//   circle := Circle{x:0,y:0,radius:5}
//   rectangle := Rectangle {width:10, height:5}
//
//   fmt.Printf("Circle area: %f\n",getArea(circle))
//   fmt.Printf("Rectangle area: %f\n",getArea(rectangle))
//}
////Circle area: 78.539816
////Rectangle area: 50.000000
//```
//
//惯用的思路是先定义接口，再定义实现，最后定义使用的方法：
//
//```
//package animals
//
//type Animal interface {
//    Speaks() string
//}
//
//// implementation of Animal
//type Dog struct{}
//func (a Dog) Speaks() string { return "woof" }
//
///** 在需要的地方直接引用 **/
//
//package circus
//
//import "animals"
//
//func Perform(a animal.Animal) { return a.Speaks() }
//```
//
//Go 也为我们提供了另一种接口的实现方案，我们可以不在具体的实现处定义接口，而是在需要用到该接口的地方，该模式为：
//
//```
//func funcName(a INTERFACETYPE) CONCRETETYPE
//```
//
//定义接口：
//
//```
//package animals
//
//type Dog struct{}
//func (a Dog) Speaks() string { return "woof" }
//
///** 在需要使用实现的地方定义接口 **/
//package circus
//
//type Speaker interface {
//    Speaks() string
//}
//
//func Perform(a Speaker) { return a.Speaks() }
//```
//

// ### Embedding
func TestSwitch12(t *testing.T) {}

// Go 语言中并没有子类继承这样的概念，而是通过嵌入（Embedding）的方式来
// 实现类或者接口的组合。
//
// ```
// // ReadWriter 的实现需要同时满足 Reader 与 Writer
//
//	type ReadWriter interface {
//	   Reader
//	   Writer
//	}
//
// // Server 暴露了所有 Logger 结构体的方法
//
//	type Server struct {
//	   Host string
//	   Port int
//	   *log.Logger
//	}
//
// // 初始化方式并未受影响
// server := &Server{"localhost", 80, log.New(...)}
//
// // 却可以直接调用内嵌结构体的方法，等价于 server.Logger.Log(...)
// server.Log(...)
//
// // 内嵌结构体的名词即是类型名
// var logger *log.Logger = server.Logger
// ```
//
// ## 并发编程
func TestSwitch13(t *testing.T) {}

// ### Goroutines
//
// Goroutines 是轻量级的线程，可以参考[并发编程导论](https://parg.co/UnK)一文中的进程、线程与协程的讨论；Go 为我们提供了非常便捷的 Goroutines 语法：
//
// ```
// // 普通函数
// func doStuff(s string) {
// }
//
//	func main() {
//	   // 使用命名函数创建 Goroutine
//	   go doStuff("foobar")
//
//	   // 使用匿名内部函数创建 Goroutine
//	   go func (x int) {
//	       // function body goes here
//	   }(42)
//	}
//
// ```
//
// ### Channels
func TestSwitch14(t *testing.T) {}

// 信道（Channel）是带有类型的管道，可以用于在不同的 Goroutine 之间传递消息，其基础操作如下：
//
// ```
// // 创建类型为 int 的信道
// ch := make(chan int)
//
// // 向信道中发送值
// ch <- 42
//
// // 从信道中获取值
// v := <-ch
//
// // 读取，并且判断其是否关闭
// v, ok := <-ch
//
// // 读取信道，直至其关闭
//
//	for i := range ch {
//	   fmt.Println(i)
//	}
//
// ```
//
// 譬如我们可以在主线程中等待来自 Goroutine 的消息，并且输出：
//
// ```
// // 创建信道
// messages := make(chan string)
//
// // 执行 Goroutine
// go func() { messages <- "ping" }()
//
// // 阻塞，并且等待消息
// msg := <-messages
//
// // 使用信道进行并发地计算，并且阻塞等待结果
// c := make(chan int)
// go sum(s[:len(s)/2], c)
// go sum(s[len(s)/2:], c)
// x, y := <-c, <-c // 从 c 中接收
// ```
//
// 如上创建的是无缓冲型信道（Non-buffered Channels），其是阻塞型信道；当没有值时读取方会持续阻塞，而写入方则是在无读取时阻塞。我们可以创建缓冲型信道（Buffered Channel），其读取方在信道被写满前都不会被阻塞：
//
// ```
// ch := make(chan int, 100)
//
// // 发送方也可以主动关闭信道
// close(ch)
// ```
//
// Channel 同样可以作为函数参数，并且我们可以显式声明其是用于发送信息还是接收信息，从而增加程序的类型安全度：
//
// ```
// // ping 函数用于发送信息
//
//	func ping(pings chan<- string, msg string) {
//	   pings <- msg
//	}
//
// // pong 函数用于从某个信道中接收信息，然后发送到另一个信道中
//
//	func pong(pings <-chan string, pongs chan<- string) {
//	   msg := <-pings
//	   pongs <- msg
//	}
//
//	func main() {
//	   pings := make(chan string, 1)
//	   pongs := make(chan string, 1)
//	   ping(pings, "passed message")
//	   pong(pings, pongs)
//	   fmt.Println(<-pongs)
//	}
//
// ```
//
// ### 同步
func TestSwitch15(t *testing.T) {}

//同步，是并发编程中的常见需求，这里我们可以使用 Channel 的阻塞特性来实现 Goroutine 之间的同步：
//
//```
//func worker(done chan bool) {
//    time.Sleep(time.Second)
//    done <- true
//}
//
//func main() {
//    done := make(chan bool, 1)
//    go worker(done)
//
//    // 阻塞直到接收到消息
//    <-done
//}
//```
//
//Go 还为我们提供了 select 关键字，用于等待多个信道的执行结果：
//
//```
//// 创建两个信道
//c1 := make(chan string)
//c2 := make(chan string)
//
//// 每个信道会以不同时延输出不同值
//go func() {
//    time.Sleep(1 * time.Second)
//    c1 <- "one"
//}()
//go func() {
//    time.Sleep(2 * time.Second)
//    c2 <- "two"
//}()
//
//// 使用 select 来同时等待两个信道的执行结果
//for i := 0; i < 2; i++ {
//    select {
//    case msg1 := <-c1:
//        fmt.Println("received", msg1)
//    case msg2 := <-c2:
//        fmt.Println("received", msg2)
//    }
//}
//```
//

// unit testing

// bench testing
