package main

import (
	"math/rand"
	"testing"
)

const (
	MaxRandLimit = 1000
)

func generateRandomArray(arrayLen int) []int {
	var a []int
	for i := 0; i < arrayLen; i++ {
		a = append(a, rand.Intn(MaxRandLimit))
	}
	return a
}

func benchmarkSortBubble(i int, b *testing.B) {
	a := generateRandomArray(i)
	sortBubble(a)
}


func BenchmarkSortBubble100(b *testing.B) {
	benchmarkSortBubble(100, b)
}

func BenchmarkSortBubble1000(b *testing.B) {
	benchmarkSortBubble(1000, b)
}

func BenchmarkSortBubble10000(b *testing.B) {
	benchmarkSortBubble(10000, b)
}

func BenchmarkSortBubble100000(b *testing.B) {
	benchmarkSortBubble(100000, b)
}

func BenchmarkSortBubble1000000(b *testing.B) {
	benchmarkSortBubble(1000000, b)
}

func BenchmarkSortBubble10000000(b *testing.B) {
	benchmarkSortBubble(10000000, b)
}

