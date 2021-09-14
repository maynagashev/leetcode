package main

import "fmt"

func main() {

}

var level = 0

func letterCasePermutation(s string) []string {
	level++
	fmt.Printf("letterCasePermutation [%d]: %s\n", level, s)
	if len(s) == 0 {
		return []string{""}
	}
	var res []string
	c := s[len(s)-1]
	for _, str := range letterCasePermutation(s[:len(s)-1]) {
		fmt.Println(str)
		// add original char
		res = append(res, str+string(c))

		// add inverted if needed
		if c >= 'a' && c <= 'z' {
			res = append(res, str+string(c-32)) // add str + last UPPER
		} else if c >= 'A' && c <= 'Z' {
			res = append(res, str+string(c+32)) // add str + last LOWER
		}
	}
	fmt.Printf("result letterCasePermutation [%d]: %s\n%v\n-\n", level, s, res)
	level--
	return res

}