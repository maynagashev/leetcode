/*
Package binary_search

Time complexity : O(logN).
Space complexity: O(1)
Both: 32 ms	6.9 MB
*/
package binary_search

import "sort"

func BinarySearchClassic(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	var mid int
	for start <= end {
		mid = (start+end)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			end = mid-1
		} else {
			start = mid+1
		}
	}
	return -1
}

// BinarySearch using standard lib function https://pkg.go.dev/sort#SearchInts
func BinarySearch(nums []int, target int) int {
	i := sort.SearchInts(nums, target)
	if i < len(nums) && nums[i]== target {
		return i
	}
	return -1
}