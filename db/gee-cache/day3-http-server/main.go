package main

/*
$ curl http://localhost:9999/_geecache/scores/Tom
630

$ curl http://localhost:9999/_geecache/scores/kkk
kkk not exist
*/

import (
	"fmt"
	"geecache"
	"log"
	"net/http"
	"sync"
)

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

var GroupName = "scores"

var once sync.Once

func InitCache() {
	geecache.NewGroup(GroupName, 2<<10, geecache.GetterFunc(
		func(key string) ([]byte, error) { // 实际业务过程中，这里可能是在查数据库
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))
}

func main() {
	once.Do(InitCache)

	// 查询缓存数据
	gcache := geecache.GetGroup(GroupName)
	d, _ := gcache.Get("Tom")
	fmt.Println("Tom score", d.String())
	gcache.Get("Tom") // cache hit

	// 启动一个带缓存的http服务器
	addr := "localhost:9999"
	peers := geecache.NewHTTPPool(addr)
	log.Println("geecache is running at", addr)
	log.Fatal(http.ListenAndServe(addr, peers))
}
