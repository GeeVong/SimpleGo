package main

import (
	"fmt"
	"sync"
	"time"
)

type Queue struct {
	mu    sync.Mutex
	cond  *sync.Cond
	items []int
}

func (q *Queue) Enqueue(item int) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.items = append(q.items, item)
	q.cond.Signal()
}

func (q *Queue) Dequeue() int {
	q.mu.Lock()
	defer q.mu.Unlock()

	for len(q.items) == 0 {
		q.cond.Wait()
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func main() {
	queue := &Queue{
		cond: sync.NewCond(&sync.Mutex{}),
	}

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		item := queue.Dequeue()
		fmt.Println("Dequeued:", item)
	}()

	time.Sleep(1 * time.Second)
	go func() {
		defer wg.Done()
		queue.Enqueue(42)
		fmt.Println("Enqueued: 42")
	}()

	wg.Wait()

	fmt.Println("111111")
}
