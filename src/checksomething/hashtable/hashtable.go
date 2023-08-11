package main

import (
	"crypto/sha256"
	"fmt"
)

// 哈希表节点的结构体
type Node struct {
	key   string
	value interface{}
	next  *Node
}

// 哈希表结构体
type HashTable struct {
	size    int
	buckets []*Node
}

// 创建一个新的哈希表
func NewHashTable(size int) *HashTable {
	return &HashTable{
		size:    size,
		buckets: make([]*Node, size),
	}
}

/*
开放地址法:生成哈希值
	当 Key3 与已经存入哈希表中的两个键值对 Key1 和 Key2 发生冲突时，Key3 会被写入 Key2 后面的空闲位置。
	当我们再去读取 Key3 对应的值时就会先获取键的哈希并取模，这会先帮助我们找到 Key1，
	找到 Key1 后发现它与 Key 3 不相等，所以会继续查找后面的元素，直到内存为空或者找到目标元素。
*/

func t1() {
	data := []byte("Hello, World!")
	hash := sha256.Sum256(data)
	fmt.Printf("Hash value: %x\n", hash)
}

// 根据键生成哈希值
func hash(key string, size int) int {
	hash := 0
	for _, c := range key {
		hash += int(c)
	}
	return hash % size
}

// 在哈希表中插入键值对
func (ht *HashTable) Insert(key string, value interface{}) {
	// 根据键生成哈希值
	index := hash(key, ht.size)

	// 创建新的节点
	newNode := &Node{
		key:   key,
		value: value,
	}

	// 如果当前桶为空，直接插入
	if ht.buckets[index] == nil {
		ht.buckets[index] = newNode
	} else { // 否则插入到链表的末尾
		current := ht.buckets[index]
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

// 在哈希表中查找键对应的值
func (ht *HashTable) Find(key string) interface{} {
	// 根据键生成哈希值
	index := hash(key, ht.size)

	// 在链表中查找键对应的值
	current := ht.buckets[index]
	for current != nil {
		if current.key == key {
			return current.value
		}
		current = current.next
	}

	// 如果未找到，则返回nil
	return nil
}

func main() {
	ht := NewHashTable(10)

	ht.Insert("apple", 5)
	ht.Insert("banana", 3)
	ht.Insert("cherry", 8)

	fmt.Println(ht.Find("apple"))  // 输出: 5
	fmt.Println(ht.Find("banana")) // 输出: 3
	fmt.Println(ht.Find("cherry")) // 输出: 8
	fmt.Println(ht.Find("durian")) // 输出: nil
}
