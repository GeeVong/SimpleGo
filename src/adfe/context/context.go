package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 创建根上下文
	ctx := context.Background()

	// 启动goroutine并传递上下文
	go doSomething(ctx)

	// 等待一段时间，然后取消上下文
	time.Sleep(time.Second * 2)
	fmt.Println("Canceling the context")
	// 手动取消上下文
	cancelCtx := context.TODO()
	_, cancel := context.WithCancel(cancelCtx)
	defer cancel()
	cancel()

	// 等待goroutine执行完毕
	time.Sleep(time.Second)
	fmt.Println("Exiting...")
}

func doSomething(ctx context.Context) {
	for {
		select {
		case <-time.After(time.Second):
			fmt.Println("Doing something...")
		case <-ctx.Done():
			fmt.Println("Context is canceled")
			return
		}
	}
}
