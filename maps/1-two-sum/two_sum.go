/*
Package two_sum

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.

1) Brute Force:
Time complexity: O(n^2). For each element, we try to find its complement by looping through the rest of the array which takes O(n) time. Therefore, the time complexity is O(n^2).
Space complexity: O(1). The space required does not depend on the size of the input array, so only constant space is used.

2) Hash Table (generally we care more about "time" than "space", because it is easy to increase the size of memory, but we can never buy more time):
Time complexity: O(n) - gets better.
Space complexity: O(n) - gets worse. 
*/
package two_sum

/*
One-pass Hash Table

Runtime: 4 ms, faster than 97.54% of Go online submissions for Two Sum.
Memory Usage: 4.3 MB, less than 21.91% of Go online submissions for Two Sum.
Time complexity: O(n). We traverse the list containing nn elements only once. Each lookup in the table costs only O(1) time.
Space complexity: O(n). The extra space required depends on the number of items stored in the hash table, which stores at most n elements.
*/
func twoSum(nums []int, target int) []int {
	seenNums := make(map[int]int)
	for index, thisNum := range nums {
		if seenIndex, ok := seenNums[target - thisNum]; ok {
			return []int{seenIndex, index}
		}
		seenNums[thisNum] = index
	}
	return []int{0, 0} // Should not happen
}