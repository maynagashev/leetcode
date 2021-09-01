/*
https://leetcode.com/problems/find-the-duplicate-number/
https://www.youtube.com/watch?v=FumVLLWUrew
*/
package main

func main() {

}

func findDuplicate(nums []int) int {
	slow := nums[0]
	fast := nums[0]
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	fast = nums[0]
	for fast != slow {
		fast = nums[fast]
		slow = nums[slow]
	}

	return fast
}
