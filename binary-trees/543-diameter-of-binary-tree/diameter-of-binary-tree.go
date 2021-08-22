package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Post-order
*/
func diameterOfBinaryTree(root *TreeNode) int {
	diameter := 0
	longestPath(root, &diameter)
	return diameter
}

func longestPath(node *TreeNode, maxDiameter *int) int {
	if node == nil {
		return 0
	}

	left := longestPath(node.Left, maxDiameter)
	right := longestPath(node.Right, maxDiameter)

	if left+right > *maxDiameter {
		*maxDiameter = left + right
	}

	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}
