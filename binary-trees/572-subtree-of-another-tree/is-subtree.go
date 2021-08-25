package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}

	if isSameTree(root, subRoot) {
		return true
	}

	// recursively traversing both branches and check is next node is subtree
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

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
