package tree

import "fmt"

func PrintTestTree() {
	tree := &TreeNode{Value: 10}
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(3)
	tree.Insert(7)
	tree.Insert(13)
	tree.Insert(17)
	tree.Insert(1)
	tree.Insert(4)
	tree.Insert(6)
	tree.Insert(8)

	fmt.Println(ClosestValue(tree, 12))
	fmt.Println(ClosestValue(tree, 16))
	fmt.Println(ClosestValue(tree, 18))
	fmt.Println(ClosestValue(tree, 1))
	fmt.Println(ClosestValue(tree, 4))
	fmt.Println(ClosestValue(tree, 6))
	fmt.Println(ClosestValue(tree, 8))
	fmt.Println(SumAllBranches(tree))
}

func ClosestValue(tree *TreeNode, value int) int {
	closest := tree.Value

	if closest > value {
		if tree.Left == nil {
			return closest
		}
		return ClosestValue(tree.Left, value)
	}

	if closest < value {
		if tree.Right == nil {
			return closest
		}
		return ClosestValue(tree.Right, value)
	}
	return closest
}

func SumAllBranches(tree *TreeNode) []int {
	result := []int{}
	sum := 0
	result = append(result, SumBranch(tree, sum, &result)...)
	return result
}

func SumBranch(t *TreeNode, sum int, result *[]int) []int {
	if t == nil {
		return nil
	}
	partialSum := sum + t.Value

	if t.Left == nil && t.Right == nil {
		*result = append(*result, partialSum)
	}
	SumBranch(t.Left, partialSum, result)
	SumBranch(t.Right, partialSum, result)

	return *result
}
