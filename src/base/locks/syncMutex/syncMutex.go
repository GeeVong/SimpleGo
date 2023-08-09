package main

import (
	"fmt"
	"sync"
)

// A Mutex is a mutual exclusion lock.
// The zero value for a Mutex is an unlocked mutex.
//
// A Mutex must not be copied after first use.

type Counter struct {
	mu         sync.Mutex
	count      int
	stringData string
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *Counter) dataChange(s string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.stringData = s
}

func main() {
	counter := &Counter{}

	counter.dataChange("aa")
	// 处于占用状态
	go counter.Increment()

	counter.dataChange("aa")

	fmt.Println("stringData:", counter.stringData, "count:", counter.count)
	// stringData: aa count: 0
}
