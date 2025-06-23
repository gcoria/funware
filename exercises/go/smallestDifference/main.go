package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println("Hello, 世界")
	fmt.Println(SmallestDifference([]int{1, 10, 5, 12, -19, 3, 8}))
	fmt.Println(SmallestDifferenceNlogN([]int{1, 10, 5, 12, -19, 3, 8}))
}

func SmallestDifference(numbers []int) float64 {
	smallest := float64(1000)
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if i == j {
				continue
			}
			current := math.Abs(float64(numbers[i] - numbers[j]))
			if current < smallest {
				smallest = current
			}
		}
	}
	return smallest
} // 0(N^2)

func SmallestDifferenceNlogN(numbers []int) int {
	sort.Ints(numbers)
	smallest := math.MaxInt32

	for i := 1; i < len(numbers); i++ {
		current := math.Abs(float64(numbers[i] - numbers[i-1]))
		if current < float64(smallest) {
			smallest = int(current)
		}
	}
	return smallest
}
