/*
You are given two binary trees root1 and root2.

Imagine that when you put one of them to cover the other, some nodes of the two trees are overlapped while the others are not. You need to merge the two trees into a new binary tree. The merge rule is that if two nodes overlap, then sum node values up as the new value of the merged node. Otherwise, the NOT null node will be used as the node of the new tree.

Return the merged tree.

Note: The merging process must start from the root nodes of both trees.
*/
package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Pre-order traversing (NLR)

Time complexity : O(m). A total of m nodes need to be traversed. Here, mm represents the minimum number of nodes from the two given trees.
Space complexity : O(m). The depth of the recursion tree can go upto m in the case of a skewed tree. In average case, depth will be O(logm).
*/
func mergeTreesRecursive(node1 *TreeNode, node2 *TreeNode) *TreeNode {

	if node1 == nil {
		return node2
	}
	if node2 == nil {
		return node1
	}

	// if both ok
	node1.Val += node2.Val

	node1.Left = mergeTreesRecursive(node1.Left, node2.Left)
	node1.Right = mergeTreesRecursive(node1.Right, node2.Right)

	return node1
}
