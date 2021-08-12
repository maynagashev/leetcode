/*
Package add_two_numbers

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.
You may assume the two numbers do not contain any leading zero, except the number 0 itself.
*/
package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
Runtime: 4 ms, faster than 98.38% of Go online submissions for Add Two Numbers.
Memory Usage: 4.8 MB, less than 93.38% of Go online submissions for Add Two Numbers.
Time complexity: O(max(m,n)). Assume that mm and nn represents the length of l1l1 and l2l2 respectively, the algorithm above iterates at most max(m,n) times.
Space complexity: O(max(m,n)). The length of the new list is at most max(m,n)+1.
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	current := &ListNode{}
	result := current
	var carry, x, y, sum int

	for l1 != nil || l2 != nil {
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		} else {
			x = 0
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		} else {
			y = 0
		}

		sum = x + y + carry
		carry = sum / 10

		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}
	if carry > 0 {
		current.Next = &ListNode{Val: carry}
	}
	return result.Next
}
