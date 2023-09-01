package lru

import (
	"fmt"
	"github.com/GeeVong/SimpleGo/log"
	"reflect"
	"testing"
)

type String string

func (d String) Len() int {
	return len(d)
}

func TestGet(t *testing.T) {
	log.SetLogLevel(3)
	log.InitZeroLogCfg()

	lru := New(int64(0), nil)
	lru.Add("key1", String("1234"))
	v, ok := lru.Get("key1")
	if !ok {
		log.LogError("lru TestGet", " key1 is not found")
	} else {
		log.LogInfo("lru TestGet", "cache hit , value is %v", v)
	}

	if string(v.(String)) != "1234" {
		log.LogError("lru TestGet", "the value is not 1234")
	}
	_, ok = lru.Get("key2")
	if !ok {
		log.LogError("lru TestGet", "miss key2")
	}
}

func TestRemoveoldest(t *testing.T) {
	k1, k2, k3 := "key1", "key2", "k3"
	v1, v2, v3 := "value1", "value2", "v3"
	cap := len(k1 + k2 + v1 + v2)
	lru := New(int64(cap), nil)
	lru.Add(k1, String(v1))
	lru.Add(k2, String(v2))
	lru.Add(k3, String(v3))

	if _, ok := lru.Get("key1"); ok || lru.Len() != 2 {
		t.Fatalf("Removeoldest key1 failed")
	}
}

func TestOnEvicted(t *testing.T) {
	log.SetLogLevel(3)
	log.InitZeroLogCfg()

	keys := make([]string, 0)
	callback := func(key string, value Value) {
		keys = append(keys, key)
	}
	lru := New(int64(10), callback)

	lru.Add("k1", String("v1"))
	getListInfo(lru)

	lru.Add("k2", String("v2"))
	getListInfo(lru)

	/*
		2023-08-29T11:57:09+08:00 INF lru.go:90 > nbytes:33 ,maxBytes10 [tag]=TestOnEvicted
		ll's element:{k3,v2222288888888888888888} c.nbytes:33,elementSize:25
		ll's element:{k2,v2} c.nbytes:33,elementSize:4
		ll's element:{k1,v1} c.nbytes:33,elementSize:4
		2023-08-29T11:59:03+08:00 INF lru.go:90 > nbytes:29 ,maxBytes10 [tag]=TestOnEvicted
		ll's element:{k3,v2222288888888888888888} c.nbytes:29,elementSize:25
		ll's element:{k2,v2} c.nbytes:29,elementSize:4
		2023-08-29T11:59:30+08:00 INF lru.go:90 > nbytes:25 ,maxBytes10 [tag]=TestOnEvicted
		ll's element:{k3,v2222288888888888888888} c.nbytes:25,elementSize:25

	*/

	lru.Add("k3", String("v2222288888888888888888"))
	getListInfo(lru)
	lru.Add("k2", String("v2222288888888888888888"))
	getListInfo(lru)

	expect := []string{"key1", "k2"}

	if !reflect.DeepEqual(expect, keys) {
		t.Fatalf("Call OnEvicted failed, expect keys equals to %s", expect)
	}
}

func TestAdd(t *testing.T) {
	lru := New(int64(0), nil)
	lru.Add("key", String("1"))
	lru.Add("key", String("111"))

	if lru.nbytes != int64(len("key")+len("111")) {
		t.Fatal("expected 6 but got", lru.nbytes)
	}
}

func typeSize(typ reflect.Type) int {
	size := typ.Size()
	return int(size)
}

func getListInfo(c *Cache) {
	l := c.ll
	log.LogInfo("TestOnEvicted", "nbytes:%d ,maxBytes%d", c.nbytes, c.maxBytes)
	for e := l.Front(); e != nil; e = e.Next() {
		elementSize := int64(len(e.Value.(*entry).key)) + int64(e.Value.(*entry).value.Len())
		fmt.Printf("ll's element:{%s,%s} ", e.Value.(*entry).key, e.Value.(*entry).value)
		fmt.Printf("c.nbytes:%d,elementSize:%d \n", c.nbytes, elementSize)
	}
	fmt.Println()
}

func dataSize(d any) {
	log.LogInfo("dataSize", "value d's size %d", typeSize(reflect.TypeOf(d)))
}
