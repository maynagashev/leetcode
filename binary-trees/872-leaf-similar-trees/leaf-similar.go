package leaf_similar_trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Function to traverse the tree and collect leaf node values into a slice.
func leafSequence(root *TreeNode, leaves *[]int) {
	if root == nil {
		return
	}

	// Check if the node is a leaf (both left and right children are nil).
	if root.Left == nil && root.Right == nil {
		*leaves = append(*leaves, root.Val)
		return
	}

	// Traverse left and right subtrees.
	leafSequence(root.Left, leaves)
	leafSequence(root.Right, leaves)
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	// Collect leaf node values into slices for both trees.
	var leaves1, leaves2 []int
	leafSequence(root1, &leaves1)
	leafSequence(root2, &leaves2)

	// Check if the leaf sequences are equal.
	if len(leaves1) != len(leaves2) {
		return false
	}

	for i := 0; i < len(leaves1); i++ {
		if leaves1[i] != leaves2[i] {
			return false
		}
	}

	return true
}
