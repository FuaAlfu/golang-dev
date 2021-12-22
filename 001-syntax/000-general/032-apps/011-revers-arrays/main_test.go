package main

import "testing"

func BenchmarkRevOne(b *testing.B) {
	s := []int{1, 2, 3}
	for i := 0; i < b.N; i++ {
		reverseOne(s)
	}
}

func BenchmarkRevTwo(b *testing.B) {
	s := []int{1, 2, 3}
	for i := 0; i < b.N; i++ {
		reverseTwo(s)
	}
}

func BenchmarkRevThree(b *testing.B) {
	s := []int{1, 2, 3}
	for i := 0; i < b.N; i++ {
		reverseThree(s)
	}
}
