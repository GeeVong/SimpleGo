package lru

type LRUCache struct {
	head, tail *Node
	Keys       map[int]*Node
	Cap        int
}

type Node struct {
	Key, Val   int
	Prev, Next *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{Keys: make(map[int]*Node), Cap: capacity}
}

// 获取数据
func (this *LRUCache) Get(key int) int {
	if node, ok := this.Keys[key]; ok {
		this.Remove(node)
		this.Add(node)
		return node.Val
	}
	return -1
}

// 向cache 写入数据
func (this *LRUCache) Put(key int, value int) {
	// 1. 字典里面发现存在数据，
	// 2. 将原来节点删除
	// 3. 新数据节点移动到双向链表表头
	if node, ok := this.Keys[key]; ok {
		node.Val = value
		this.Remove(node)
		this.Add(node)
		return
	} else {
		// 字典中没有找到，新数据节点插入到双向链表表头
		node = &Node{Key: key, Val: value}
		this.Keys[key] = node
		this.Add(node)
	}
	// 超过cache容量，会先做插入动作《这里其实以及超过cache容量了》，移除双向链表尾节点
	if len(this.Keys) > this.Cap {
		delete(this.Keys, this.tail.Key)
		this.Remove(this.tail)
	}
}

// 向链表头插入数据
func (this *LRUCache) Add(node *Node) {
	node.Prev = nil
	node.Next = this.head
	if this.head != nil {
		this.head.Prev = node
	}
	this.head = node
	if this.tail == nil {
		this.tail = node
		this.tail.Next = nil
	}
}

// 移除最后节点
func (this *LRUCache) Remove(node *Node) {
	if node == this.head {
		this.head = node.Next
		node.Next = nil
		return
	}
	if node == this.tail {
		this.tail = node.Prev
		node.Prev.Next = nil
		node.Prev = nil
		return
	}
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}
