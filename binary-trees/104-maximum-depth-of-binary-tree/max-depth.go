package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

/*
DFS strategy with recursion.

Time complexity : we visit each node exactly once, thus the time complexity is O(N), where N is the number of nodes.
Space complexity : in the worst case, the tree is completely unbalanced, e.g. each node has only left child node,
the recursion call would occur N times (the height of the tree), therefore the storage to keep the call stack would be O(N).
But in the best case (the tree is completely balanced), the height of the tree would be log(N). Therefore, the space complexity in this case would be O(log(N)).
 */
func maxDepthRecursive(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepthRecursive(root.Left)
	right := maxDepthRecursive(root.Right)
	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}
