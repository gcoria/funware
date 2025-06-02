package main

import (
	"fmt"
	"strconv"
)

func RiverSizes(matrix [][]int) []int {
	var sizes []int
	visited := make(map[string]bool)

	for i, row := range matrix {
		for j, value := range row {
			if value == 1 && !visited[coordKey(i, j)] {
				size := exploreRiver(matrix, i, j, visited)
				sizes = append(sizes, size)
			}
		}
	}
	return sizes
}

func exploreRiver(matrix [][]int, i, j int, visited map[string]bool) int {
	stack := [][2]int{{i, j}}
	size := 0

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		x, y := current[0], current[1]
		key := coordKey(x, y)

		if x < 0 || y < 0 || x >= len(matrix) || y >= len(matrix[0]) || matrix[x][y] == 0 || visited[key] {
			continue
		}

		visited[key] = true
		size++

		stack = append(stack, [2]int{x + 1, y})
		stack = append(stack, [2]int{x - 1, y})
		stack = append(stack, [2]int{x, y + 1})
		stack = append(stack, [2]int{x, y - 1})
	}

	return size
}

func coordKey(x, y int) string {
	return strconv.Itoa(x) + "," + strconv.Itoa(y)
}

func main() {
	matrix := [][]int{
		{1, 0, 0, 1, 0},
		{1, 0, 1, 0, 0},
		{0, 0, 1, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 1, 1, 0},
	}
	fmt.Println(RiverSizes(matrix))
}
