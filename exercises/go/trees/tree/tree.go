package tree

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func (node *TreeNode) Insert(value int) {
	if value < node.Value {
		if node.Left == nil {
			node.Left = &TreeNode{Value: value}
		} else {
			node.Left.Insert(value)
		}
	} else {
		if node.Right == nil {
			node.Right = &TreeNode{Value: value}
		} else {
			node.Right.Insert(value)
		}
	}
}
