package peak_index_in_a_mountain_array

/*
Time Complexity: O(logN), where NN is the length of A.
Space Complexity: O(1).
Runtime: 16 ms, faster than 6.56% of Go online submissions for Peak Index in a Mountain Array.
Memory Usage: 4.8 MB, less than 9.02% of Go online submissions for Peak Index in a Mountain Array.
*/
func peakIndexInMountainArray(arr []int) int {
	start := 0
	end := len(arr) - 1
	mid := 0
	for start <= end {
		mid = start + (end - start)/2
		if arr[mid]<arr[mid+1] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return start
}