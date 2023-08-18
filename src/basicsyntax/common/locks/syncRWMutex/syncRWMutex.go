package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Cache struct {
	mu    sync.RWMutex
	cache map[string]string
}

func (c *Cache) Get(key string) string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.cache[key]
}

func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = value
}

// 测试读锁占用，写操作
/*
	==== 设置初值
	========show the cache info ========
	k: 2 value: 65
	==== 并发读
	========show the cache info ========
	k: 2 value: 65
	Exiting.
结果：3 67 并没有被写入 ，由于读锁占用，写锁被排斥
*/
func testWriteWhenCacheRead() {
	cache := &Cache{
		cache: make(map[string]string),
	}

	randomInt := rand.Intn(101)
	rand.Seed(time.Now().UnixNano())
	fmt.Println("==== 设置初值")
	cache.Set("2", fmt.Sprintf("%d", randomInt))
	showCacheInfo(cache) // 当前状态只有一个2 65

	// 并发读
	fmt.Println("==== 并发读")
	//testRead(cache)

	go cache.Set("3", fmt.Sprintf("%d", randomInt+2))
	showCacheInfo(cache) // 当前状态只有一个2 65  3 67
}

//测试写锁占用，读操作
func testReadWhenCacheWrite() {
	cache := &Cache{
		cache: make(map[string]string),
	}

	randomInt := rand.Intn(101)
	rand.Seed(time.Now().UnixNano())

	cache.Set("2", fmt.Sprintf("%d", randomInt))
	fmt.Println("========2show the cache info ========")
	for key, value := range cache.cache {
		fmt.Println("k:", key, "value:", value)
	}

	go func() {
		value := cache.Get(fmt.Sprintf("%d", "2"))
		fmt.Println("===cache.Get====", "2", value)
	}()

}

// 测试并发写
func testWrite() {
	cache := &Cache{
		cache: make(map[string]string),
	}

	randomInt := rand.Intn(101)
	rand.Seed(time.Now().UnixNano())
	cache.Set("2", fmt.Sprintf("%d", randomInt))

	//showCacheInfo(cache)

	//time.Sleep(1 * time.Second)

	ws := sync.WaitGroup{}
	ws.Add(1)
	go cache.Set("2", fmt.Sprintf("%d", randomInt+111))
	ws.Done()
	ws.Wait()
	showCacheInfo(cache)
}

// 测试并发读
func testRead(cache *Cache) {
	var wg sync.WaitGroup
	fmt.Println("========show the cache info ========")
	for key, value := range cache.cache {
		fmt.Println("k:", key, "value:", value)
	}

	wg.Wait()
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(i int) {
			value := cache.Get(fmt.Sprintf("%d", i))
			fmt.Println("k,value", i, value)
			wg.Done()
		}(i)
	}

	wg.Wait()
}

func showCacheInfo(cache *Cache) {
	var wg sync.WaitGroup
	fmt.Println("========show the cache info ========")
	for key, value := range cache.cache {
		fmt.Println("k:", key, "value:", value)
	}
	wg.Wait()
}

func main() {
	//testWriteWhenCacheRead()
	//testReadWhenCacheWrite()
	testWrite()
}
