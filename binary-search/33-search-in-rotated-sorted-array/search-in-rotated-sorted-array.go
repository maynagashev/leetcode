package _33_search_in_rotated_sorted_array

import "fmt"

// [4,5,6,7,0,1,2] target = 0
func search(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low <= high {

		mid := (low + high) / 2
		// fmt.Printf("mid: %v val: %v\n", mid, nums[mid])
		if nums[mid] == target {
			return mid
		}

		// if first half is sorted
		if nums[low] <= nums[mid] {
			if target < nums[mid] && target >= nums[low] {
				// fmt.Printf("target in sorted first half\n")
				high = mid - 1
			} else {
				fmt.Printf("target in unsorted second half, drop left\n")
				low = mid + 1
			}
			// else if second hald is sorted
		} else if nums[high] >= nums[mid] {
			// if target in sorted second half, drop first half
			if target > nums[mid] && target <= nums[high] {
				// fmt.Printf("target in sorted second half\n")
				low = mid + 1
			} else {
				// fmt.Printf("target in unsorted left half, drop right\n")
				high = mid - 1
			}
		}
		// fmt.Printf("new range: %v (%v) - %v (%v)\n", nums[low], low, nums[high], high)
	}

	// Replace this placeholder return statement with your code
	return -1
}
