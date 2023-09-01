package geecache

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"testing"
)

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func TestGetter(t *testing.T) {
	var f Getter = GetterFunc(func(key string) ([]byte, error) {
		return []byte(key), nil
	})

	expect := []byte("key")
	if v, _ := f.Get("key"); !reflect.DeepEqual(v, expect) {
		t.Fatal("callback failed")
	}
}

func TestGet(t *testing.T) {
	loadCounts := make(map[string]int, len(db))
	gee := NewGroup("scores", 2<<10, GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				if _, ok := loadCounts[key]; !ok {
					loadCounts[key] = 0
				}
				loadCounts[key]++
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))

	for k, v := range db {
		if view, err := gee.Get(k); err != nil || view.String() != v {
			t.Fatal("failed to get value of Tom")
		}
		if _, err := gee.Get(k); err != nil || loadCounts[k] > 1 {
			t.Fatalf("cache %s miss", k)
		}
	}

	if view, err := gee.Get("unknown"); err == nil {
		t.Fatalf("the value of unknow should be empty, but %s got", view)
	}
}

func TestGetGroup(t *testing.T) {
	groupName := "scores"
	NewGroup(groupName, 2<<10, GetterFunc(
		func(key string) (bytes []byte, err error) { return }))
	if group := GetGroup(groupName); group == nil || group.name != groupName {
		t.Fatalf("group %s not exist", groupName)
	}

	if group := GetGroup(groupName + "111"); group != nil {
		t.Fatalf("expect nil, but %s got", group.name)
	}
}

type Info struct {
	Name string // `json:"name"`
	Age  int    // `json:"age"`
}

func Test_AddDate(t *testing.T) {
	groupName := "userCache"
	name := "geeVon"
	user := &Info{Name: name, Age: 12}

	g := NewGroup(groupName, 2<<10, GetterFunc(
		func(key string) (bytes []byte, err error) {

			if key == user.Name {
				data, err := json.Marshal(user)
				if err != nil {
					fmt.Println("转换失败：", err)
					return nil, fmt.Errorf("%s not exist", key)
				}
				return data, nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))

	//g.Get(name)
	view, err := g.Get(name)
	if nil != err {
		t.Fatal("failed to get value of geeVon")
	}
	var Msg Info
	err = json.Unmarshal(view.b, &Msg)
	if nil != err {
		fmt.Println("Unmarshal:", err)
	}
	fmt.Println(view.b)
	fmt.Println(Msg)
}

// 结构体必须大写服了
func TestMarshal(t *testing.T) {
	type nfo struct {
		name string // `json:"name"`
		Age  int    // `json:"age"`
	}

	name := "geeVon"
	user := &nfo{name: name, Age: 12}
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("转换失败：", err)
	} else {
		fmt.Println(data)
	}

	var Msg nfo
	err = json.Unmarshal(data, &Msg)
	if nil != err {
		fmt.Println("Unmarshal:", err)
	} else {
		fmt.Println(Msg)
	}

}
