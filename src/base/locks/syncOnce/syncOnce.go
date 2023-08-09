package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
	Data string
}

var instance *Singleton
var once sync.Once

func GetInstance(str string) *Singleton {
	once.Do(func() {
		instance = &Singleton{Data: str}
	})
	return instance
}

func main() {
	s1 := GetInstance("Initialized Data 1")
	s2 := GetInstance("Initialized Data 2")

	fmt.Println(s1.Data)
	fmt.Println(s2.Data)
}
