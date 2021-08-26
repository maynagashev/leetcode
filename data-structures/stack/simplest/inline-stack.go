package main

import "fmt"

func main() {
	// Init stack
	stack := []int{1, 2, 3}
	fmt.Println(stack)

	// Pop item
	c := stack[len(stack)-1]     // pick last
	stack = stack[:len(stack)-1] // reslice stack
	fmt.Println("Popped:", c, stack)

	// Push item
	stack = append(stack, 4)
	fmt.Println("Pushed 4", stack)
}
