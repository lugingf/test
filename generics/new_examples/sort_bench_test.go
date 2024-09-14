package main

import (
	"math/rand"
	"sort"
	"testing"
)

// Non-generic sorting
func sortInts(data []int) {
	sort.Slice(data, func(i, j int) bool { return data[i] < data[j] })
}

// Generic sorting
type SortInterfaceT[T any] interface {
	sort.Interface
}

func SortT[T any](data SortInterfaceT[T]) {
	sort.Sort(data)
}

type IntSliceT []int

func (p IntSliceT) Len() int           { return len(p) }
func (p IntSliceT) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSliceT) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func randomIntSlice(size int) []int {
	data := make([]int, size)
	for i := range data {
		data[i] = rand.Intn(size)
	}
	return data
}

func BenchmarkSortInts(b *testing.B) {
	data := randomIntSlice(1000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sortInts(data)
	}
}

func BenchmarkSortGenerics(b *testing.B) {
	data := IntSlice(randomIntSlice(1000))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SortT(data)
	}
}

//go test -bench .
