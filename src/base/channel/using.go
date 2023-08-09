//package main.go
//
//import (
//	"fmt"
//	"sync"
//)
//
//type Task struct {
//	id int
//}
//
//var wg sync.WaitGroup
//
//func worker(id int, taskQueue <-chan Task) {
//	defer wg.Done()
//
//	for task := range taskQueue {
//		fmt.Printf("Worker %d processing Task %d\n", id, task.id)
//		// 实际任务处理逻辑
//	}
//}
//
//func main.go() {
//	numWorkers := 5
//	numTasks := 10
//
//	taskQueue := make(chan Task)
//
//	// 创建工作 Goroutine
//	wg.Add(numWorkers)
//	for i := 0; i < numWorkers; i++ {
//		go worker(i+1, taskQueue)
//	}
//
//	// 向任务队列中添加任务
//	for i := 0; i < numTasks; i++ {
//		taskQueue <- Task{id: i + 1}
//	}
//	close(taskQueue)
//
//	// 等待所有工作 Goroutine 结束
//	wg.Wait()
//}
