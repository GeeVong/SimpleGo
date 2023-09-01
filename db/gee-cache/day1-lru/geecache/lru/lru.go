package lru

import (
	"container/list"
)

/*
FIFO(First In First Out)
		先进先出，也就是淘汰缓存中最老(最早添加)的记录。FIFO 认为，最早添加的记录，
		其不再被使用的可能性比刚添加的可能性大。这种算法的实现也非常简单，创建一个队列，
		新增记录添加到队尾，每次内存不够时，淘汰队首。但是很多场景下，部分记录虽然是最早添加但也最常被访问，
		而不得不因为呆的时间太长而被淘汰。这类数据会被频繁地添加进缓存，又被淘汰出去，导致缓存命中率降低。
-
*/

/*
LFU(Least Frequently Used)
		最少使用，也就是淘汰缓存中访问频率最低的记录。LFU 认为，如果数据过去被访问多次，
		那么将来被访问的频率也更高。LFU 的实现需要维护一个按照访问次数排序的队列，每次访问，
		访问次数加1，队列重新排序，淘汰时选择访问次数最少的即可。LFU 算法的命中率是比较高的，
		但缺点也非常明显，维护每个记录的访问次数，对内存的消耗是很高的；
		另外，如果数据的访问模式发生变化，LFU 需要较长的时间去适应，也就是说 LFU 算法受历史数据的影响比较大。
		例如某个数据历史上访问次数奇高，但在某个时间点之后几乎不再被访问，但因为历史访问次数过高，而迟迟不能被淘汰。
-
*/

/*
LRU(Least Recently Used)
		最近最少使用，相对于仅考虑时间因素的 FIFO 和仅考虑访问频率的 LFU，
		LRU 算法可以认为是相对平衡的一种淘汰算法。LRU 认为，如果数据最近被访问过，
		那么将来被访问的概率也会更高。LRU 算法的实现非常简单，维护一个队列，
		如果某条记录被访问了，则移动到队尾，那么队首则是最近最少访问的数据，淘汰该条记录即可。

-
*/

// Cache is a LRU cache. It is not safe for concurrent access.
type Cache struct {
	maxBytes int64                    // 允许使用的最大内存
	nbytes   int64                    // 当前已使用的内存
	ll       *list.List               // 双向链表
	cache    map[string]*list.Element // 数据字典
	// optional and executed when an entry is purged.
	OnEvicted func(key string, value Value) // OnEvicted 是某条记录被移除时的回调函数
}

type entry struct {
	key   string
	value Value
}

// Value use Len to count how many bytes it takes
type Value interface {
	Len() int
}

// New is the Constructor of Cache
func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

// Add adds a value to the cache.
func (c *Cache) Add(key string, value Value) {
	// 如果键存在，则更新对应节点的值，并将该节点移到列表最前面。
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.nbytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		// 字典中中没有找到
		// 向链表表头，插入数据
		// 向字典添加key，并存储对应的链表节点
		// 计算内存大小
		ele := c.ll.PushFront(&entry{key, value})
		c.cache[key] = ele
		c.nbytes += int64(len(key)) + int64(value.Len())
	}

	// 是否超过最大缓存大小，如果大于则删除链表尾节点
	// 如果来一个非常大的数据，大于最大cache容量，那会造成所有缓存不可以用 todo
	for c.maxBytes != 0 && c.maxBytes < c.nbytes {
		//l := c.ll
		//log.LogInfo("TestOnEvicted", "nbytes:%d ,maxBytes%d", c.nbytes, c.maxBytes)
		//for e := l.Front(); e != nil; e = e.Next() {
		//	elementSize := int64(len(e.Value.(*entry).key)) + int64(e.Value.(*entry).value.Len())
		//	fmt.Printf("ll's element:{%s,%s} ", e.Value.(*entry).key, e.Value.(*entry).value)
		//	fmt.Printf("c.nbytes:%d,elementSize:%d \n", c.nbytes, elementSize)
		//}

		c.RemoveOldest()
	}
}

// Get look ups a key's value
func (c *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}

// RemoveOldest removes the oldest item
func (c *Cache) RemoveOldest() {
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		c.nbytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

// Len the number of cache entries
// 为了方便测试，我们实现 Len() 用来获取添加了多少条数据。
func (c *Cache) Len() int {
	return c.ll.Len()
}
