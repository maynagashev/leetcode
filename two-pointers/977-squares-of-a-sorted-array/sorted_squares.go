package main

import "fmt"

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(nums))
	// [0 1 9 16 100]
}

/*
Time Complexity: O(N), where NN is the length of A.
Space Complexity: O(N) if you take output into account and O(1) otherwise.
*/
func sortedSquares(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	left := 0
	right := n - 1

	// filling result from right to left (backwards)
	for i := n - 1; i >= 0; i-- {
		squareLeft := nums[left] * nums[left]
		squareRight := nums[right] * nums[right]
		if squareLeft > squareRight {
			result[i] = squareLeft
			left++
		} else {
			result[i] = squareRight
			right--
		}
	}

	return result
}
