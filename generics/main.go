package main

import (
	"fmt"
	"strings"
)

// Number constraint covers common numeric types
type Number interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

// Map applies fn to every element and returns a new slice
func Map[T, U any](s []T, fn func(T) U) []U {
	out := make([]U, len(s))
	for i, v := range s {
		out[i] = fn(v)
	}
	return out
}

// Filter returns elements for which keep returns true
func Filter[T any](s []T, keep func(T) bool) []T {
	var out []T
	for _, v := range s {
		if keep(v) {
			out = append(out, v)
		}
	}
	return out
}

// Reduce folds a slice into a single value
func Reduce[T, U any](s []T, init U, fn func(U, T) U) U {
	acc := init
	for _, v := range s {
		acc = fn(acc, v)
	}
	return acc
}

// Sum adds all numbers in a slice using the Number constraint
func Sum[T Number](s []T) T {
	var total T
	for _, v := range s {
		total += v
	}
	return total
}

// Stack is a generic LIFO data structure
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(v T) { s.items = append(s.items, v) }
func (s *Stack[T]) Len() int { return len(s.items) }
func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top, true
}

func main() {
	nums := []int{1, 2, 3, 4, 5}

	doubled := Map(nums, func(n int) int { return n * 2 })
	fmt.Println("doubled:", doubled)

	evens := Filter(nums, func(n int) bool { return n%2 == 0 })
	fmt.Println("evens:", evens)

	sum := Reduce(nums, 0, func(acc, n int) int { return acc + n })
	fmt.Println("sum via Reduce:", sum)

	fmt.Println("Sum[float64]:", Sum([]float64{1.1, 2.2, 3.3}))

	words := []string{"go", "is", "great"}
	upper := Map(words, strings.ToUpper)
	fmt.Println("uppercased:", upper)

	var stack Stack[int]
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	fmt.Println("stack len:", stack.Len())
	for {
		v, ok := stack.Pop()
		if !ok {
			break
		}
		fmt.Println("popped:", v)
	}
}
