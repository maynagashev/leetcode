package min_depth

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	q := []*TreeNode{root}
	depth := 1
	for len(q) > 0 {
		n := len(q)
		for i := 0; i < n; i++ {

			// dequeue
			node := q[0]
			q = q[1:]

			// if found any node without children, got a result
			if node.Left == nil && node.Right == nil {
				return depth
			}

			// queue children
			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		depth++
	}

	return depth
}
