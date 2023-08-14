package main

import (
	"basicsyntax/common"
	"fmt"
	"math"
	"math/cmplx"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"testing"
	"time"
)

/*
	======================		变量数据类型		======================
*/
// 变量声明与赋值
func TestVariable(t *testing.T) {
	// 声明三个变量，皆为 bool 类型
	var c, python, java bool

	// 声明不同类型的变量，并且赋值
	var i bool = true
	var j int = 2

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

// 基本
func TestBasicType(t *testing.T) {
	var b bool = true
	common.TestMinMax(b, false)

	var d_byte byte = 2
	common.TestMinMax(d_byte, false)

	var i32 int32 = -2147483648
	common.TestMinMax(i32, true)

	var iint int = 1
	common.TestMinMax(iint, true)

	var iint8 int8 = 2
	common.TestMinMax(iint8, true)

	{
		a, c := 0b1010, 0x64
		b := 0o144

		fmt.Printf("0b%b, %#o, %#x\n", a, b, c)
		fmt.Println(math.MinInt8, math.MaxInt8)
	}
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

func TestSwitch2(t *testing.T) {}

//### 类型绑定与初始化
//
//Go 中的 type 关键字能够对某个类型进行重命名：
//
//```
//// IntSlice 并不等价于 []int，但是可以利用类型转换进行转换
//type IntSlice []int
//a := IntSlice{1, 2}
//```
//

func TestSwitch1(t *testing.T) {}

// // 类型转换与判断
// str, ok := val.(string);
// ```
//
// ### 基本数据类型
func TestSwitch4(t *testing.T) {}

// ```
func TestSwitch3(t *testing.T) {}

// #### 字符串
//
// ```
// // 多行字符串声明
// hellomsg := `
// "Hello" in Chinese is 你好 ('Ni Hao')
// "Hello" in Hindi is नमस्ते ('Namaste')
// `
// ```
//
// 格式化字符串：
func TestSwitch5(t *testing.T) {}

// ```
// fmt.Println("Hello, 你好, नमस्ते, Привет, ᎣᏏᏲ") // basic print, plus newline
// p := struct { X, Y int }{ 17, 2 }
// fmt.Println( "My point:", p, "x coord=", p.X ) // print structs, ints, etc
// s := fmt.Sprintln( "My point:", p, "x coord=", p.X ) // print to string variable
//
// fmt.Printf("%d hex:%x bin:%b fp:%f sci:%e",17,17,17,17.0,17.0) // c-ish format
// s2 := fmt.Sprintf( "%d %f", 17, 17.0 ) // formatted print to string variable
// ```
//
// ### 序列类型
//
// Array 与 Slice 都可以用来表示序列数据，二者也有着一定的关联。
//
// #### Array
func TestSwitch6(t *testing.T) {}

// 其中 Array 用于表示固定长度的，相同类型的序列对象，可以使用如下形式创建：
//
// ```
// [N]Type
// [N]Type{value1, value2, ..., valueN}
//
// // 由编译器自动计算数目
// [...]Type{value1, value2, ..., valueN}
// ```
//
// 其具体使用方式为：
//
// ```
// // 数组声明
// var a [10]int
//
// // 赋值
// a[3] = 42
//
// // 读取
// i := a[3]
//
// // 声明与初始化
// var a = [2]int{1, 2}
// a := [2]int{1, 2}
// a := [...]int{1, 2}
// ```
//
// Go 内置了 len 与 cap 函数，用于获取数组的尺寸与容量：
//
// ```
// var arr = [3]int{1, 2, 3}
// arr := [...]int{1, 2, 3}
//
// len(arr) // 3
// cap(arr) // 3
// ```
//
// 不同于 C/C++ 中的指针（Pointer）或者 Java 中的对象引用（Object Reference），Go 中的 Array 只是值（Value）。这也就意味着，当进行数组拷贝，或者函数调用中的参数传值时，会复制所有的元素副本，而非仅仅传递指针或者引用。显而易见，这种复制的代价会较为昂贵。
//
// #### Slice
func TestSwitch7(t *testing.T) {}

// Slice 为我们提供了更为灵活且轻量级地序列类型操作，可以使用如下方式创建 Slice:
//
// ```
// // 使用内置函数创建
// make([]Type, length, capacity)
// make([]Type, length)
//
// // 声明为不定长度数组
// []Type{}
// []Type{value1, value2, ..., valueN}
//
// // 对现有数组进行切片转换
// array[:]
// array[:2]
// array[2:]
// array[2:3]
// ```
//
// 不同于 Array，Slice 可以看做更为灵活的引用类型（Reference Type），它并不真实地存放数组值，而是包含数组指针（ptr），len，cap 三个属性的结构体。换言之，Slice 可以看做对于数组中某个段的描述，包含了指向数组的指针，段长度，以及段的最大潜在长度，其结构如下图所示：
//
// ![img](assets/1460000014069224.png)
//
// ```
// // 创建 len 为 5，cap 为 5 的 Slice
// s := make([]byte, 5)
//
// // 对 Slice 进行二次切片，此时 len 为 2，cap 为 3
// s = s[2:4]
//
// // 恢复 Slice 的长度
// s = s[:cap(s)]
// ```
//
// 需要注意的是， 切片操作并不会真实地复制 Slice 中值，只是会创建新的指向原数组的指针，这就保证了切片操作和操作数组下标有着相同的高效率。不过如果我们修改 Slice 中的值，那么其会 真实修改底层数组中的值，也就会体现到原有的数组中：
//
// ```
// d := []byte{'r', 'o', 'a', 'd'}
// e := d[2:]
// // e == []byte{'a', 'd'}
// e[1] = 'm'
// // e == []byte{'a', 'm'}
// // d == []byte{'r', 'o', 'a', 'm'}
// ```
//
// Go 提供了内置的 append 函数，来动态为 Slice 添加数据，该函数会返回新的切片对象，包含了原始的 Slice 中值以及新增的值。如果原有的 Slice 的容量不足以存放新增的序列，那么会自动分配新的内存：
//
// ```
// // len=0 cap=0 []
// var s []int
//
// // len=1 cap=2 [0]
// s = append(s, 0)
//
// // len=2 cap=2 [0 1]
// s = append(s, 1)
//
// // len=5 cap=8 [0 1 2 3 4]
// s = append(s, 2, 3, 4)
//
// // 使用 ... 来自动展开数组
// a := []string{"John", "Paul"}
// b := []string{"George", "Ringo", "Pete"}
// a = append(a, b...) // equivalent to "append(a, b[0], b[1], b[2])"
// // a == []string{"John", "Paul", "George", "Ringo", "Pete"}
// ```
//
// 我们也可以使用内置的 copy 函数，进行 Slice 的复制，该函数支持对于不同长度的 Slice 进行复制，其会自动使用最小的元素数目。同时，copy 函数还能够自动处理使用了相同的底层数组之间的 Slice 复制，以避免额外的空间浪费。
//
// ```
// func copy(dst, src []T) int
//
// // 申请较大的空间容量
// t := make([]byte, len(s), (cap(s)+1)*2)
// copy(t, s)
// s = t
// ```
//
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
