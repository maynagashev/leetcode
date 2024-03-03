package main

import "fmt"

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Printf("%s\n%v\n", s, isPalindrome(s))
}

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1

	for left < right {
		if !isAlphanumeric(s[left]) {
			left++
			continue
		}
		if !isAlphanumeric(s[right]) {
			right--
			continue
		}
		if lower(s[left]) != lower(s[right]) {
			return false
		}
		left++
		right--
	}
	return true
}

func isAlphanumeric(u uint8) bool {
	return (u >= 'a' && u <= 'z') || (u >= 'A' && u <= 'Z') || (u >= '0' && u <= '9')
}

func lower(i byte) byte {
	if i >= 'A' && i <= 'Z' {
		return i + 32
	}
	return i
}
