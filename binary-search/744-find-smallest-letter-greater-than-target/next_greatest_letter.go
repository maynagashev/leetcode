package next_greatest_letter

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Find Smallest Letter Greater Than Target.
Memory Usage: 2.7 MB, less than 100.00% of Go online submissions for Find Smallest Letter Greater Than Target.
Time Complexity: O(logN), where NN is the length of letters. We peek only at logN elements in the array.
Space Complexity: O(1), as we maintain only pointers.
*/
func nextGreatestLetter(letters []byte, target byte) byte {
	start := 0
	end := len(letters) - 1
	mid := 0
	for start <= end {
		mid = start + (end-start)/2
		if letters[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return letters[start%len(letters)]
}
