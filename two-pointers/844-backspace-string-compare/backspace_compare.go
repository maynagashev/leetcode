/*
Given two strings s and t, return true if they are equal when both are typed into empty text editors. '#' means a backspace character.

Note that after backspacing an empty text, the text will continue empty.
*/
package main

import "fmt"

func main() {
	fmt.Println(backspaceCompare("a##c", "#a#c"))
}

/*
Two pointers.

Time Complexity: O(M+N), where M,N are the lengths of S and T respectively.
Space Complexity: O(1).
*/
func backspaceCompare(s string, t string) bool {
	i := len(s) - 1
	j := len(t) - 1
	// iterating from end
	for i >= 0 || j >= 0 {
		i = nextIndexToKeep(s, i)
		j = nextIndexToKeep(t, j)

		if i < 0 && j < 0 {
			return true
		}
		if i < 0 || j < 0 {
			return false
		}
		if s[i] != t[j] {
			return false
		}
		i--
		j--
	}
	return true

}

func nextIndexToKeep(s string, i int) int {
	skip := 0
	for i >= 0 {
		if s[i] == '#' {
			skip++
		} else if skip > 0 {
			skip--
		} else {
			break
		}
		i--
	}
	return i
}
