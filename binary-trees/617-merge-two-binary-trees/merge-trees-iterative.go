package main

func mergeTrees(node1 *TreeNode, node2 *TreeNode) *TreeNode {
	s := stack{}
	s.Push([]*TreeNode{node1, node2})
	for !s.isEmpty() {
		t := s.Pop()
		if t[0] == nil || t[1] == nil {
			continue
		}
		t[0].Val += t[1].Val

		if t[0].Left == nil  {
			t[0].Left = t[1].Left
		} else {
			s.Push([]*TreeNode{t[0].Left, t[1].Left})
		}

		if t[0].Right == nil {
			t[0].Right = t[1].Right
		} else {
			s.Push([]*TreeNode{t[0].Right, t[1].Right})
		}
	}

	return node1
}

type stack [][]*TreeNode

func (s *stack) Push(v []*TreeNode) {
	*s = append(*s, v)
}
func (s *stack) Pop() []*TreeNode {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}
func (s *stack) isEmpty() bool {
	return len(*s) == 0
}
