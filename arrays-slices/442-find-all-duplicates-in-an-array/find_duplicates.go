package main

func main() {

}

/*
Mark Visited Elements in the Input Array itself

Time complexity : O(n). We iterate over the array twice. Each negation operation occurs in constant time.
Space complexity : No extra space required, other than the space for the output list. We re-use the input array to store visited status.
*/
func findDuplicates(nums []int) []int {
	duplicates := []int{}
	for _, num := range nums {
		if nums[abs(num)-1] < 0 {
			duplicates = append(duplicates, abs(num))
		}
		nums[abs(num)-1] *= -1
	}
	return duplicates
}

func abs(n int) int {
	if n < 0 {
		return -n
	} else {
		return n
	}
}
