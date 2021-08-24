/*
FIFO - first in, first out
*/
package main

import "fmt"

func main() {
	queue := queueOfInt{1, 2, 3}
	fmt.Println("Queue: ", queue)
	fmt.Println("Dequeue: ", queue.dequeue())
	fmt.Println("Dequeue: ", queue.dequeue())
	fmt.Println("Enqueue: 4", queue.enqueue(4))
	fmt.Println("Queue: ", queue)
	fmt.Println("Dequeue: ", queue.dequeue())
	fmt.Println("Queue: ", queue)
}

type queueOfInt []int

func (q *queueOfInt) enqueue(val int) *queueOfInt {
	*q = append(*q, val)
	return q
}
func (q *queueOfInt) dequeue() int {
	val := (*q)[0]
	*q = (*q)[1:]
	return val
}
func (q *queueOfInt) isEmpty() bool {
	return len(*q) == 0
}
