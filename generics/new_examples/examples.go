package main

import (
	"fmt"
	"sort"
)

// Slide 5: Swap example
func Swap[T any](a, b *T) {
	*a, *b = *b, *a
}

func swapExample() {
	a, b := 1, 2
	Swap(&a, &b)
	fmt.Println(a, b)

	s1, s2 := "Hello", "World"
	Swap(&s1, &s2)
	fmt.Println(s1, s2)
}

// Slide 6: Stack example
type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Push(value T) {
	s.data = append(s.data, value)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.data) == 0 {
		var zero T
		return zero, false
	}
	value := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return value, true
}

func stackExample() {
	intStack := Stack[int]{}
	intStack.Push(1)
	intStack.Push(2)
	fmt.Println(intStack.Pop())

	stringStack := Stack[string]{}
	stringStack.Push("Hello")
	stringStack.Push("World")
	fmt.Println(stringStack.Pop())
}

// Slide 7: Map example
func Map[T, R any](list []T, mapper func(T) R) []R {
	result := make([]R, len(list))
	for i, v := range list {
		result[i] = mapper(v)
	}
	return result
}

func mapExample() {
	numbers := []int{1, 2, 3, 4, 5}
	squares := Map(numbers, func(n int) int { return n * n })
	fmt.Println(squares)

	words := []string{"Hello", "Golang", "Generics"}
	lengths := Map(words, func(s string) int { return len(s) })
	fmt.Println(lengths)
}

// Slide 8: Filter example
func Filter[T any](list []T, predicate func(T) bool) []T {
	result := []T{}
	for _, v := range list {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func filterExample() {
	numbers := []int{1, 2, 3, 4, 5}
	evenNumbers := Filter(numbers, func(n int) bool { return n%2 == 0 })
	fmt.Println(evenNumbers)

	words := []string{"", "Hello", "", "Golang", "Generics", ""}
	nonEmptyWords := Filter(words, func(s string) bool { return s != "" })
	fmt.Println(nonEmptyWords)
}

// Slide 9: Sort example
type SortInterface[T any] interface {
	sort.Interface
}

func Sort[T any](data SortInterface[T]) {
	sort.Sort(data)
}

// sort/sort.go
type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func sortExample() {
	numbers := IntSlice{3, 1, 4, 1, 5, 9, 2, 6}
	Sort(numbers)
	fmt.Println(numbers)
}

func main() {

}
