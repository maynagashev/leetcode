package length_of_longest_substring

import "fmt"

/*
Optimized sliding window, where we store last seen position of symbol - in hash map. Stored position helps us to move left pointer (i).

Time complexity : O(n). Index j will iterate n times.
Space complexity (HashMap) : O(min(m,n)). We need O(k) space for the sliding window, where k is the size of the Set. The size of the Set is upper bounded by the size of the string n and the size of the charset/alphabet m.
Runtime: 4 ms, faster than 90.73% of Go online submissions for Longest Substring Without Repeating Characters.
Memory Usage: 3.1 MB, less than 57.14% of Go online submissions for Longest Substring Without Repeating Characters.
*/
func lengthOfLongestSubstring(s string) int {

	// map with lastSeenIndex + 1
	mp := make(map[byte]int)
	maxLen := 0
	i := 0
	for j := 0; j < len(s); j++ {
		if lastSeenIndexPlus, seen := mp[s[j]]; seen && lastSeenIndexPlus > i {
			i = lastSeenIndexPlus
		}
		if j-i+1 > maxLen {
			maxLen = j - i + 1
		}
		mp[s[j]] = j + 1
	}
	return maxLen
}

func printAsAscii(str string) {
	runes := []rune(str)
	var result []int

	for i := 0; i < len(runes); i++ {
		result = append(result, int(runes[i]))
	}

	fmt.Println(result)
}
