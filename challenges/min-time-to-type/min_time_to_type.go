package min_time_to_type

import "fmt"

type ListNode struct {
	Next *ListNode
	Prev *ListNode
	Val  rune
}

func MinTimeToType(word string) int {
	time := 0
	current := NewLinkedList() // a

	for _, r := range word {
		fmt.Printf("Letter: %c \n", r)
		leftTime, left := search("left", r, current)
		rightTime, right := search("right", r, current)

		if leftTime <= rightTime {
			time += leftTime
			current = left
		} else {
			time += rightTime
			current = right
		}
		fmt.Printf("lefTime: %d, rightTime: %d, current: %c\n", leftTime, rightTime, current.Val)
	}
	fmt.Printf("Result time: %d\n", time)
	return time
}

func search(dir string, r rune, current *ListNode) (int, *ListNode) {
	time:=0
	for current.Val != r {
		switch dir {
		case "left":
			current = current.Prev
		case "right":
			current = current.Next
		}
		time++ // move
	}
	time++ // pick
	return time, current
}

func NewLinkedList() *ListNode {

	sentinelNode := &ListNode{
		Prev: nil,
		Next: nil,
		Val:  0,
	}

	current := sentinelNode
	for r := 'a'; r <= 'z'; r++ {
		next := &ListNode{
			Prev: current,
			Next: nil,
			Val:  r,
		}
		current.Next = next
		current = next
	}
	sentinelNode.Next.Prev = current // a.Prev = z
	current.Next = sentinelNode.Next // z.Next = a


	current = sentinelNode.Next // a
	for i := 0; i < 26; i++ {
		fmt.Printf("%c (%c, %c)", current.Val, current.Prev.Val, current.Next.Val)
		current = current.Next
	}
	fmt.Println("---")
	return sentinelNode.Next
}