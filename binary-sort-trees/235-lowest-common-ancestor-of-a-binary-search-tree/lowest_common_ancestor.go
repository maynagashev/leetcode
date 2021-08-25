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