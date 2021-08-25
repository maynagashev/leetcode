package main

type nodeWitDepth struct {
	node  *TreeNode
	depth int
}

type stack []*nodeWitDepth

func (s *stack) Push(node *nodeWitDepth) {
	*s = append(*s, node)
}
func (s *stack) Pop() *nodeWitDepth {
	res := (*s)[len(*s)-1] // last element
	*s = (*s)[:len(*s)-1]
	return res
}
func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxDepth := 0
	s := stack{}
	s.Push(&nodeWitDepth{root, 1})
	for !s.isEmpty() {
		v := s.Pop()
		if v.node == nil {
			continue
		}
		// update max maxDepth
		if v.depth > maxDepth {
			maxDepth = v.depth
		}
		s.Push(&nodeWitDepth{v.node.Left, v.depth + 1})
		s.Push(&nodeWitDepth{v.node.Right, v.depth + 1})
	}

	return maxDepth
}
