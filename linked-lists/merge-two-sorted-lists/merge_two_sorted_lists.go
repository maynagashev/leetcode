package merge_two_sorted_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	current := result
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	// if l1 longer, handle l1's tail
	if l1 != nil {
		current.Next = l1
		l1 = l1.Next
		current = current.Next
	}

	// if l2 longer, handle l2's tail
	if l2 != nil {
		current.Next = l2
		l2 = l2.Next
		current = current.Next
	}

	return result.Next
}
