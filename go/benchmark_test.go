package main

import (
	"testing"
)

func main() {
	helloWorld()
}

func helloWorld() int {
	return 1 + 1 + 1
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		helloWorld()
	}
}
