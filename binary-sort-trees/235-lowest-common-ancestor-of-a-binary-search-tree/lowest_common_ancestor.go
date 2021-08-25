package main

func main() {

}


type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}


func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		if p.Val>root.Val && q.Val>root.Val { // both right
			root = root.Right
		} else if p.Val<root.Val && q.Val<root.Val { // both left
			root = root.Left
			return lowestCommonAncestor(root, p, q)
		} else {
			return root
		}
	}
	return root
}

func lowestCommonAncestorRecursive(root, p, q *TreeNode) *TreeNode {
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestorRecursive(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestorRecursive(root.Right, p, q)
	}
	return root
}