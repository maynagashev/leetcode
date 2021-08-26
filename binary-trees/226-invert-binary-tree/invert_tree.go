package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTreeRecursive(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = invertTreeRecursive(root.Right), invertTreeRecursive(root.Left)
	return root
}
