package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"runtime/pprof"
)

var ProfilePath = "./src/simplesyntax/pprof/profile/"

func fib(n int) int {
	if n <= 1 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

// cpu pprof
func cpu_pprof() {
	os.MkdirAll(ProfilePath, os.ModePerm)
	f, _ := os.OpenFile(ProfilePath+"cpu.profile", os.O_CREATE|os.O_RDWR, 0644)
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	n := 10
	for i := 1; i <= 5; i++ {
		fmt.Printf("fib(%d)=%d\n", n, fib(n))
		n += 3 * i
	}
}

const Letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func generate(n int) string {
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		buf.WriteByte(Letters[rand.Intn(len(Letters))])
	}
	return buf.String()
}

func repeat(s string, n int) string {
	var result string
	for i := 0; i < n; i++ {
		result += s
	}

	return result
}

func memory_pprof() {
	f, _ := os.OpenFile(ProfilePath+"mem.profile", os.O_CREATE|os.O_RDWR, 0644)
	//defer f.Close()
	for i := 0; i < 100; i++ {
		repeat(generate(100), 100)
	}

	pprof.Lookup("heap").WriteTo(f, 0)
}

func main() {
	cpu_pprof()
	memory_pprof()
}
