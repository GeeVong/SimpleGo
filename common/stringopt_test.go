package common

import (
	"strings"
	"testing"
)

var S = strings.Repeat("a", 100)

func normalConv() bool {
	b := []byte(S)
	s2 := string(b)
	//s2 := S
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
