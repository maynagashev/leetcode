/*
You are given a 0-indexed integer array nums, where nums[i] represents the score of the ith student. You are also given an integer k.
Pick the scores of any k students from the array so that the difference between the highest and the lowest of the k scores is minimized.
Return the minimum possible difference.
*/
package main

import (
	"math"
	"sort"
)

func main() {

}

/*
Sliding window
1. Sort array
2. Moving start & end of k elements, calculating difference
 */
func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	n := len(nums)
	res := math.MaxInt32
	// fmt.Println(nums)
	for i := 0; i < n-k+1; i++ {
		min := nums[i]
		max := nums[i+k-1]
		// fmt.Println(min, max)
		if max-min < res {
			res = max - min
		}
	}
	return res
}
