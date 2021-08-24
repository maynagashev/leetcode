/*
Simple and short stack implementation
https://stackoverflow.com/questions/28541609/looking-for-reasonable-stack-implementation-in-golang
*/
package main

import (
	"fmt"
)

func main() {
	s := stack([]int{})
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println("Stack:", s)
	fmt.Println("Pop:", s.Pop(), s.Pop(), s.Pop())

	if s.isEmpty() {
		fmt.Println("Empty stack")
	} else {
		fmt.Println(s.Pop())
	}
}

type stack []int

func (s *stack) Push(v int) {
	*s = append(*s, v)
}
func (s *stack) Pop() int {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}
func (s *stack) isEmpty() bool {
	return len(*s) == 0
}
