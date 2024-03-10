package three_sum

import (
	"fmt"
	"sort"
)

// FindSumOfThree returns true if there are three elements in nums such that their sum equals target
// Time complexity : O(n^2). We have outer and inner loops, each going through n elements.
// Sorting the array takes O(nlogn), so overall complexity is O(nlogn + n^2) = O(n^2).
// Space complexity : from O(logn) to O(n), depending on the implementation of the sorting algorithm.
func findSumOfThree(nums []int, target int) bool {
	sort.Ints(nums)
	fmt.Printf("nums: %v\n", nums)

	for i := 0; i < len(nums)-2; i++ {
		low := i + 1
		high := len(nums) - 1
		fmt.Printf("i: %d, low: %d, high: %d\n", i, low, high)
		// двигаем low и high навстречу друг другу,
		// в случае если сумма меньше target,
		// то увеличиваем low, иначе уменьшаем high
		for low < high {
			sum := nums[i] + nums[low] + nums[high]
			fmt.Printf("sum: %d\n", sum)
			if sum == target {
				return true
			}
			if sum < target {
				low++
			}
			if sum > target {
				high--
			}
		}

	}
	return false
}
