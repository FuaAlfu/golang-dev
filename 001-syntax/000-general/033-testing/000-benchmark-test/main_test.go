package main

import "testing"

func BenchmarkFibo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci(10)
	}
}
