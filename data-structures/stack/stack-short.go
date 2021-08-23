/*
Simple and short stack implementation
https://stackoverflow.com/questions/28541609/looking-for-reasonable-stack-implementation-in-golang
*/
package main

import "fmt"

func main() {
	s := stack([]int{})
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println("Stack:", s)
	fmt.Println("Pop:", s.Pop(), s.Pop(), s.Pop())

	fmt.Println(s.Pop()) // panic: runtime error: index out of range [-1]x
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
