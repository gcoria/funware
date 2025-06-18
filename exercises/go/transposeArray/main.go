package main

import "fmt"

func main() {
	arr := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(transposeArray(arr))
}

func transposeArray(arr [][]int) [][]int {
	transposed := make([][]int, len(arr[0]))
	for i := range transposed {
		transposed[i] = make([]int, len(arr))
	}

	for i := range arr {
		for j := range arr[i] {
			transposed[j][i] = arr[i][j]
		}
	}

	return transposed
}
