package old_examples

import (
	"testing"
)

var collections = make(map[int][]int)
var max = 1000

func getArray(max int) []int {
	if col, ok := collections[max]; ok {
		return col
	}
	collection := make([]int, max)
	for i := 0; i < max; i++ {
		collection[i] = i
	}
	collections[max] = collection
	return collection
}

func InArrayInt(search int, values []int) bool {
	for _, v := range values {
		if search == v {
			return true
		}
	}
	return false
}

func InArrayGeneric[T comparable](search T, values []T) bool {
	for _, v := range values {
		if search == v {
			return true
		}
	}
	return false
}

func InArrayInterface(search interface{}, values []interface{}) bool {
	switch search.(type) {
	case int:
		for _, v := range values {
			if search == v {
				return true
			}
		}
	case string:
		for _, v := range values {
			if search == v {
				return true
			}
		}
	case int64:
		for _, v := range values {
			if search == v {
				return true
			}
		}
	case int32:
		for _, v := range values {
			if search == v {
				return true
			}
		}
	}
	return false
}

func BenchmarkInArrayInt1000(b *testing.B) {
	col := getArray(max)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		InArrayInt(max, col)
	}
}

func BenchmarkInArrayInterface(b *testing.B) {
	col := getArray(max)
	arr := make([]interface{}, 0)
	for _, num := range col {
		arr = append(arr, num)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		// we have to change slice type
		InArrayInterface(max, arr)
	}
}

func BenchmarkInArrayGen(b *testing.B) {
	col := getArray(max)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		InArrayGeneric(max, col)
	}
}

func BenchmarkInArrayGenInit(b *testing.B) {
	col := getArray(max)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		InArrayGeneric[int](max, col)
	}
}
