/*
Given a binary tree, find its minimum depth.
The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.
Note: A leaf is a node with no children.

https://leetcode.com/problems/minimum-depth-of-binary-tree/
*/
package main

import (
	"container/list"
)

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Queue (BFS) implementation.
Possible memory leak when q = q[1:]
*/
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	q := []*TreeNode{root}
	depth := 1
	for len(q) > 0 {
		n := len(q)
		for i := 0; i < n; i++ {

			// dequeue
			node := q[0]
			q = q[1:]

			// if found any node without children, got a result
			if node.Left == nil && node.Right == nil {
				return depth
			}

			// queue children
			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		depth++
	}

	return depth
}

/*
With proper memory leak handling.
*/
func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := list.New()
	q.PushBack(root)
	depth := 1
	for q.Len() > 0 {
		n := q.Len()
		for i := 0; i < n; i++ {

			// dequeue
			first := q.Front()
			node := first.Value
			q.Remove(first)

			switch node := node.(type) {
			case *TreeNode:
				// fmt.Printf("%+v %+v \n", node.Left, node.Right)
				// if found any node without children, got a result
				if node.Left == nil && node.Right == nil {
					return depth
				}

				// queue children
				if node.Left != nil {
					q.PushBack(node.Left)
				}

				if node.Right != nil {
					q.PushBack(node.Right)
				}
			}

		}
		depth++
	}

	return depth
}
