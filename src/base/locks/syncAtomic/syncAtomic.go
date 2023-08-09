package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64

	// 并发递增计数器
	for i := 0; i < 10; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
		}()
	}

	// 等待所有goroutine执行完毕
	for atomic.LoadInt64(&counter) < 10 {
	}

	fmt.Println("Counter:", counter)

	ExampleValue_readMostly()
}

func ExampleValue_readMostly() {
	type Map map[string]string
	var m atomic.Value
	m.Store(make(Map))
	var mu sync.Mutex // used only by writers
	// read function can be used to read the data without further synchronization
	read := func(key string) (val string) {
		m1 := m.Load().(Map)
		return m1[key]
	}
	// insert function can be used to update the data without further synchronization
	insert := func(key, val string) {
		mu.Lock() // synchronize with other potential writers
		defer mu.Unlock()
		m1 := m.Load().(Map) // load current value of the data structure
		m2 := make(Map)      // create a new value
		for k, v := range m1 {
			m2[k] = v // copy all data from the current object to the new one
		}
		m2[key] = val // do the update that we need
		m.Store(m2)   // atomically replace the current object with the new one
		// At this point all new readers start working with the new version.
		// The old version will be garbage collected once the existing readers
		// (if any) are done with it.
	}
	//_, _ = read, insert

	insert("1", "212321dsasadasda")
	fmt.Println(read("1"))
}
