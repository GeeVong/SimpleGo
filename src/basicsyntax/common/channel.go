package common

import (
	"fmt"
	"sync"
	"time"
)

/*
	c := make(chan bool) // 创建一个无缓冲的 bool 型 Channel
	c <- x        // 向一个 Channel 发送一个值
	<- c          // 从一个 Channel 中接收一个值
	x = <- c      // 从 Channel c 接收一个值并将其存储到 x 中
	x, ok = <- c  // 从 Channel 接收一个值，如果 channel 关闭（close）了
					或没有数据，那么 ok 将被置为 false
*/

/*
	c1 := make(chan []byte)                // 无缓冲
	c2 := male(chan []byte, 1024)           // 有缓冲
*/

/*
对于无缓冲区channel：

	发送的数据如果没有被接收方接收，那么发送方阻塞；

	如果一直接收不到发送方的数据，接收方阻塞；
*/
func noBufferChan() {
	ch := make(chan int) // 创建一个无缓冲通道
	go func() {
		fmt.Println("====222111")
		val, ok := <-ch // 接收数据，接收方会被阻塞，由于数据还没接收到，
		// 直到发送方发送数据
		if !ok {
			fmt.Println("数据发送失败")
		} else {
			fmt.Println("接收到数据:", val)
		}
	}()

	fmt.Println("====111")
	ch <- 10 // 发送数据，发送方会被阻塞直到接收方接收数据
	fmt.Println("数据发送成功")

}

/*
有缓冲的channel：

	发送方在缓冲区满的时候阻塞，接收方不阻塞；

	接收方在缓冲区为空的时候阻塞，发送方不阻塞。
*/
func bufferChan() {
	ch := make(chan int, 2) // 创建一个有缓冲大小为2的通道

	go func() {
		fmt.Println("====1")
		ch <- 10 // 发送数据，有缓冲区未满，发送方不阻塞
		ch <- 20 // 发送数据，有缓冲区未满，发送方不阻塞
		fmt.Println("====2")
		ch <- 30 // 发送数据，缓冲区已满，发送方阻塞
		fmt.Println("====3")
		fmt.Println("发送数据完成")
	}()

	time.Sleep(1 * time.Second)
	val := <-ch // 接收数据，有缓冲区非空，接收方不阻塞
	fmt.Println("接收到数据:", val)

	time.Sleep(1 * time.Second)
	val = <-ch // 接收数据，有缓冲区非空，接收方不阻塞
	fmt.Println("接收到数据:", val)

	val = <-ch // 接收数据，缓冲区为空，接收方阻塞
	fmt.Println("接收到数据:", val)
}

func createChannel() {
	type A struct {
		a int
		b float32
		d string
	}
	ch := make(chan A, 10) // make 操作
	a := A{
		a: 312,
		b: 0.001,
		d: "this is string",
	}
	ch <- a

	msg := <-ch
	fmt.Printf("%+v", msg)
	//fmt.Printf("%d %f %s", msg.a, msg.b, msg.d)
}

// 从一个 nil channel 中接收数据会一直被 block
func readDataFromNilCh() {
	var ch chan int // 未初始化的通道（nil 通道）
	msg := <-ch
	fmt.Println(msg)
}

// 向一个已经被 close 的 channel 中继续发送数据会导致 run-time panic
func writeDataFromCloseCh() {
	ch := make(chan int)
	close(ch)
	ch <- 10
	fmt.Println(<-ch)
}

// select 的 case 中建议不要出现非阻塞的条件
// （如等待某个 channel 读事件，其他 routine 将此 channel 关闭，导致该 case 永远为真）
func t() {
	ch := make(chan int)

	go func() {
		time.Sleep(time.Second)
		close(ch)
	}()

	select {
	case val := <-ch:
		fmt.Println("Received:", val)
	default:
		fmt.Println("Channel closed")
	}
}

// 控制携程数量
func controlGor() {
	var wg sync.WaitGroup
	ch := make(chan struct{}, 1024)
	for i := 0; i < 20000; i++ {
		wg.Add(1)
		//ch <- struct{}{} // 发送数据，当ch 缓冲区慢了，将会阻塞当前 goroutine，达到控制线程并发数量的目的
		ch <- struct{}{}
		go func() {
			defer wg.Done()
			<-ch
		}()
	}
	wg.Wait()

}

// 正确的从cha中读取数据的方式
func rightWayReadCh() {
	ch := make(chan int, 3)
	ch <- 10
	ch <- 10
	ch <- 10

	select {
	case val := <-ch:
		fmt.Println(val)
	default:
		fmt.Println("===default")
	}
}

func rightWayWriteCh() {
	ch := make(chan int, 2)

	for i := 0; i < 5; i++ {
		select {
		case ch <- i: // 非阻塞发送操作
			fmt.Println("写入数据:", i)
		default:
			fmt.Println("通道已满，无法写入数据")
		}
	}

	close(ch) // 关闭通道，表示发送方不再发送数据

	for val := range ch {
		fmt.Println("读取数据:", val)
	}
}

func TestForSelect() {
	chan1 := make(chan int)
	chan2 := make(chan int)
	chan3 := make(chan string)
	z := "11111"
	for {
		select {
		case x := <-chan1:
			// … 使用 x 进行一些操作
			fmt.Println(x)

		case y, ok := <-chan2:
			// 使用 y 进行一些操作，检查 ok 值判断 chan2 是否已经关闭
			if ok {
				fmt.Println(y)
			}

		case chan3 <- z:
		// z 值被成功发送到 chan3 上时

		default:
			// 上面 case 均无法执行时，执行此分支（防止程序逻辑阻塞的范式）
			//go func() {
			//	// 等待一段时间
			//	time.Sleep(2 * time.Second)
			//	// 向通道发送数据
			//	chan1 <- 20
			//}()
		}
	}

}

// 利用channel 构建任务队列
type Task struct {
	ID int
}

func worker(id int, taskQueue <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range taskQueue {
		fmt.Printf("Worker %d processing Task %d\n", id, task.ID)
		// 实际任务处理逻辑
	}
}

func taskQueue() {
	numWorkers := 5
	numTasks := 10

	taskQueue := make(chan Task, numTasks)
	wg := sync.WaitGroup{}

	// 创建工作 Goroutine
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker(i+1, taskQueue, &wg)
	}

	// 向任务队列中添加任务
	for i := 0; i < numTasks; i++ {
		taskQueue <- Task{ID: i + 1}
	}
	close(taskQueue)

	// 等待所有工作 Goroutine 结束
	wg.Wait()
}

func main() {
	// createChannel()

	// chan死锁: fatal error: all goroutines are asleep - deadlock!
	//readDataFromNilCh()

	// panic: send on closed channel
	//writeDataFromCloseCh()

	// 有无buffer chan的区别
	noBufferChan()
	bufferChan()

	// 正确读写chan的方式.
	//rightWayReadCh()
	//rightWayWriteCh()

	// chan 常使用方式 for + select
	// TestForSelect()

	taskQueue()
}
