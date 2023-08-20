# Channel && Select source code research




## 内容
  同步原语概述
  channel/select 回顾
  channel 的结构设计
  channel 的初始化行为
  channel 的发送与接收过程及其性能优化
  channel 的回收
  select 的本质及其相关编译器优化




## 并发模型
  从内存的角度看，并行计算只有两种：共享内存、消息通信（拷贝内存）
  目的是为了解决多线程数据的一致性
  基于共享内存的并发模型通常提供互斥锁同步原语
  Tony Hoare 1978 年正式发表基于输入输出（消息通信）程序结构的编程语言的设计（参见 Communication Sequential Processes (CSP) ）
  Go 将其实现为显式的 channel 同步原语
  Go 同时提供 sync.*、atomic.* 的基于共享内存的同步原语



## channel todo
  内存模型 happen-before：
  To specify the requirements of reads and writes,
  we define happens before, a partial order on the execution of 
  memory operations in a Go program. If event e1 happens before 
  event e2, then we say that e2 happens after e1. Also, 
  if e1 does not happen before e2 and does not happen after e2, 
  then we say that e1 and e2 happen concurrently.

  在 channel 中体现为：
  - buffered:   ch <- v   ≤  v := <- ch
    - buffered channel send data happen before rev data
  - unbuffered: v := <- ch  ≤    ch <- v
    buffered channel rev data happen before send data


## select
- 为什么 zero-case 会发生永久阻塞？


- uni-case 情况中，不同类型的 channel 会得到何种不同的结果，会发生什么？



- multi-case 情况中， ch1 和 ch2 的数据如果同时有效，v1 与 v2 之间读取数据 channel 的顺序？


## GMP 调度
  G: Goroutine
  被调度的实体，即用户代码，在本地队列中不断的切换执行
  
  M: Machine
  传统线程实体，即系统线程，负责代码的执行
  
  P: Processor
  处理器抽象，目的是实现串联 G 的本地队列，当 M 持有 P 时，访问的 G 不会出现数据竞争

  1. 系统线程调度本地G队列
  2. 本地队列为空
  3. 如果一个g阻塞，g应该如何处理 todo


##  channel 行为
```golang
  1. make(chan interface{}, size) ⇒ runtime.makechan(interface{}, size)
  
     make(chan interface{})       ⇒ runtime.makechan(interface{}, 0)
  
  2. ch <- v                      ⇒ runtime.chansend1(ch, &v)
  
  3. v := <- ch                   ⇒ runtime.chanrecv1(ch, &v)
  
     v, ok := <- ch               ⇒ ok := runtime.chanrecv2(ch, &v)
  
  4. close(ch)                    ⇒ runtime.closechan(ch)
```



##  channel 的结构设计
```go
  type hchan struct {
      qcount   uint           // total data in the queue
      dataqsiz uint           // size of the circular queue
	  
	  // 环形队列
      buf      unsafe.Pointer // points to an array of dataqsiz elements
      sendx    uint   // send index
      recvx    uint   // receive index

      elemsize uint16       // 元素大小
      closed   uint32       // 是否关闭
      elemtype *_type // element type

	  
	  // 双向队列 g阻塞存储在这里 todo 这里面存的是什么
      recvq    waitq  // list of recv waiters
      sendq    waitq  // list of send waiters
  
      // lock protects all fields in hchan, as well as several
      // fields in sudogs blocked on this channel.
      //
      // Do not change another G's status while holding this lock
      // (in particular, do not ready a G), as this can deadlock
      // with stack shrinking.
      lock mutex
  }
```
```go
  type waitq struct {
      first *sudog
      last  *sudog
  }
```

```go
  type sudog struct {
      // The following fields are protected by the hchan.lock of the
      // channel this sudog is blocking on. shrinkstack depends on
      // this for sudogs involved in channel ops.
  
      g *g
  
      next *sudog
      prev *sudog
      elem unsafe.Pointer // data element (may point to stack)
  
      // The following fields are never accessed concurrently.
      // For channels, waitlink is only accessed by g.
      // For semaphores, all fields (including the ones above)
      // are only accessed when holding a semaRoot lock.
  
      acquiretime int64
      releasetime int64
      ticket      uint32
  
      // isSelect indicates g is participating in a select, so
      // g.selectDone must be CAS'd to win the wake-up race.
      isSelect bool
  
      // success indicates whether communication over channel c
      // succeeded. It is true if the goroutine was awoken because a
      // value was delivered over channel c, and false if awoken
      // because c was closed.
      success bool
  
      parent   *sudog // semaRoot binary tree
      waitlink *sudog // g.waiting list or semaRoot
      waittail *sudog // semaRoot
      c        *hchan // channel
  }
```

## 创建channel
  ```go
    ch := make(chan interface{}, 4)
  ```





#  资料
 - [channel & select 源码分析](https://www.youtube.com/watch?v=d7fFCGGn0Wc&list=PLe5svQwVF1L5bNxB0smO8gNfAZQYWdIpI)
 - [PPT](https://docs.google.com/presentation/d/18_9LcMc8u93aITZ6DqeUfRvOcHQYj2gwxhskf0XPX2U/edit#slide=id.g5ea99f63e9_0_26)
 - [go_channel](https://github.com/halfrost/Halfrost-Field/blob/master/contents/Go/go_channel.md)
