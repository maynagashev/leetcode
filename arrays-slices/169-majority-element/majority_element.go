/*
Given an array nums of size n, return the majority element.

The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.
*/
package main

import "fmt"

func main() {
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2})) // 2
}

/*
Boyer-Moore Voting Algorithm

Time complexity : O(n)
Boyer-Moore performs constant work exactly nn times, so the algorithm runs in linear time.

Space complexity : O(1)
Boyer-Moore allocates only constant additional memory.
*/
func majorityElement(nums []int) int {
	var res int
	vote := 0

	for _, num := range nums {
		if vote == 0 {
			res = num
		}
		if num == res {
			vote++
		} else {
			vote--
		}
	}
	return res
}
