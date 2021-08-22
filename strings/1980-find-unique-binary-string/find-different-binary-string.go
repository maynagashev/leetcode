package main

import (
	"fmt"
	"strconv"
)

func main() {
	cases := [][]string{
		{"01", "10"},
		{"00", "01"},
		{"111", "011", "001"},
	}
	findDifferentBinaryString(cases[2])
}

/*
Contest version.
*/
func findDifferentBinaryString(nums []string) string {
	n := len(nums[0])
	seen := make(map[string]bool)
	for _, v := range nums {
		seen[v] = true
	}

	current := make([]int, n)
	currentStr := toStr(current)

	for i := 0; i < n; i++ {
		for j := 0; j <= 1; j++ {
			current[i] = j
			currentStr = toStr(current)
			fmt.Println(current, currentStr, i, j)
			if _, ok := seen[currentStr]; !ok {
			    return currentStr
			}
		}
	}
	return currentStr
}

func toStr(nums []int) string {
	res := ""
	for _, v := range nums {
		res += strconv.Itoa(v)
	}
	return res
}
