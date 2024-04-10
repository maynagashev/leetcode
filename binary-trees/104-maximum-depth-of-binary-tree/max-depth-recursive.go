package main

import (
	"fmt"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	// If the root is nil, return 0.
	if root == nil {
		return 0
	}

	// Recursively calculate the maximum depth of left and right subtrees.
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	// Return the maximum of the depths of left and right subtrees, plus 1 (for the current node).
	return 1 + max(leftDepth, rightDepth)
}

// Helper function to return the maximum of two integers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Example usage:
	// Construct a binary tree:      3
	//                             /   \
	//                            9     20
	//                                 /  \
	//                                15   7
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	fmt.Println("Maximum Depth:", maxDepth(root)) // Output: 3
}
