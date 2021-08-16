/*
Package same_tree

Given the roots of two binary trees p and q, write a function to check if they are the same or not.
Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.
*/
package same_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Same Tree.
Memory Usage: 2.2 MB, less than 18.80% of Go online submissions for Same Tree.
Time complexity : O(N) since each node is visited exactly once.
Space complexity : O(log(N)) in the best case of completely balanced tree and O(N) in the worst case of completely unbalanced tree, to keep a deque.
*/
func isSameTree(p *TreeNode, q *TreeNode) bool {
	queue := []*TreeNode{p, q}

	for len(queue) > 0 {
		p = queue[0]
		q = queue[1]
		queue = queue[2:]

		if p == nil && q == nil {
			continue
		}

		if p == nil || q == nil {
			return false
		}

		if p.Val != q.Val {
			return false
		}

		queue = append(queue, p.Left)
		queue = append(queue, q.Left)
		queue = append(queue, p.Right)
		queue = append(queue, q.Right)
	}

	return true
}
