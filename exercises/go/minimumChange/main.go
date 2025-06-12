package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MinimumChange([]int{5, 7, 1, 1, 2, 3, 22}))
	fmt.Println(MinimumChange([]int{1, 1, 1, 1, 1}))
	fmt.Println(MinimumChange([]int{1, 5, 1, 1, 1, 10, 15, 20, 100}))
	fmt.Println(MinimumChange([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(MinimumChange([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
}

func MinimumChange(coins []int) int {
	sort.Ints(coins)
	change := 0
	for _, coin := range coins {
		if coin > change+1 {
			break
		}
		change += coin
	}
	return change + 1
}
