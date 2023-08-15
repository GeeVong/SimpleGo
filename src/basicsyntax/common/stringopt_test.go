package common

import (
	"strings"
	"testing"
	"unsafe"
)

var S = strings.Repeat("a", 100)

func normalConv() bool {
	b := []byte(S)
	s2 := string(b)
	//s2 := S
	return s2 == S
}

func unsafeConv() bool {

	// []byte(s)
	b := unsafe.Slice(unsafe.StringData(S), len(S))

	// string(b)
	s2 := unsafe.String(unsafe.SliceData(b), len(b))

	return s2 == S
}

// ----------------------------------

func BenchmarkNormal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if !normalConv() {
			b.Fatal()
		}
	}
}

func BenchmarkUnsafe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if !unsafeConv() {
			b.Fatal()
		}
	}
}
