package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

}


func indexPairs(text string, words []string) [][]int {
	res := [][]int{}
	for _, word := range words {

		start := strings.Index(text, word)
		tail := text
		shift := 0
		for start != -1 {
			end := start + len(word)-1
			res = append(res,[]int{shift + start, shift + end})
			// fmt.Println(tail, word, start, end)
			tail = tail[start+1:] // cut
			shift += start+1
			fmt.Println("tail", tail)
			start = strings.Index(tail, word) // search again
		}
	}

	sort.Slice(res, func(i, j int) bool {
		if res[i][0] == res[j][0] {
			return res[i][1] <res[j][1]
		}
		return res[i][0] < res[j][0]
	})

	fmt.Println(res)

	return res
}