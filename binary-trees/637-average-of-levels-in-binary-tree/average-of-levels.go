package average_of_levels

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
func averageOfLevels(root *TreeNode) []float64 {
	q := []*TreeNode{root}
	var res []float64
	for len(q) > 0 {
		levelSum := 0
		n := len(q)

		for i := 0; i < n; i++ {
			levelSum += q[0].Val

			// queue children
			if q[0].Left != nil {
				q = append(q, q[0].Left)
			}
			if q[0].Right != nil {
				q = append(q, q[0].Right)
			}
			q = q[1:] // dequeue

		}
		res = append(res, float64(levelSum)/float64(n))

	}
	return res
}
