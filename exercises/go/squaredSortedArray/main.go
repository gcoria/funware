package main

import (
	"fmt"
	"sort"
)

func Map[T any, Q any](arr []T, f func(T) Q) []Q {
	result := make([]Q, len(arr))
	for i, v := range arr {
		result[i] = f(v)
	}
	return result
} // move to a lib future me

func SortedSquareArray(numbers []int) []int {
	squared := Map(numbers, func(n int) int { return n * n })
	sort.Ints(squared)
	return squared
}

func main() {
	numbers := []int{-2, -1, 0, 2, 3}
	sorted := SortedSquareArray(numbers)
	fmt.Println(sorted)
}
