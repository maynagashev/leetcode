package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type queueOfNodes []*TreeNode

func (queue *queueOfNodes) enqueue(node *TreeNode) *queueOfNodes {
	*queue = append(*queue, node)
	return queue
}
func (queue *queueOfNodes) dequeue() *TreeNode {
	node := (*queue)[0]
	*queue = (*queue)[1:]
	return node
}
func (queue *queueOfNodes) isEmpty() bool {
	return len(*queue) == 0
}


func isSameTree(node1 *TreeNode, node2 *TreeNode) bool {
	queue := queueOfNodes{node1, node2}
	for len(queue) > 0 {
		node1 = queue.dequeue()
		node2 = queue.dequeue()

		if node1 == nil && node2 == nil {
			continue
		}
		if node1 == nil || node2 == nil {
			return false
		}
		if node1.Val != node2.Val {
			return false
		}
		queue.enqueue(node1.Left)
		queue.enqueue(node2.Left)
		queue.enqueue(node1.Right)
		queue.enqueue(node2.Right)
	}
	return true
}
