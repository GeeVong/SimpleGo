package main

import (
	"fmt"
	"sync"
)

func worker(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\\n", id)
	// 模拟一些工作
	for i := 0; i < 5; i++ {
		fmt.Printf("Worker %d working...\\n", id)
	}
	fmt.Printf("Worker %d done\\n", id)
}

func main() {
	wg := &sync.WaitGroup{}
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(wg, i)
	}
	wg.Wait()
	fmt.Println("All workers done")
}
