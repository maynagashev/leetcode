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

/*
Since each node in the tree is visited / added to the queue only once, the time complexity is O(n), where n is the number of nodes in the tree.
Space complexity is O(n), since in the worst case, the queue will contain all nodes in one level of the binary tree.
*/
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	stack := []*TreeNode{root}

	for len(stack) > 0 {
		c := stack[len(stack)-1]     // pick last
		stack = stack[:len(stack)-1] // re-slice stack

		c.Left, c.Right = c.Right, c.Left
		if c.Left != nil {
			stack = append(stack, c.Left)
		}
		if c.Right != nil {
			stack = append(stack, c.Right)
		}
	}

	return root
}
