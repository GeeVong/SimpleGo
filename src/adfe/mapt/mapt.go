package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

func main() {
	defer pprof.StopCPUProfile()

	f, err := os.Create("s.cpuprof")
	if err != nil {
		log.Printf("start cpu profile failed: %v", err)
	}
	log.Print("start cpu profile")
	pprof.StartCPUProfile(f)

	for {
		go func() {
			m := make(map[int]string)

			m[1] = "aa"
			fmt.Println(m[1])

			m = nil
			fmt.Println(m[1])
		}()
	}

}
