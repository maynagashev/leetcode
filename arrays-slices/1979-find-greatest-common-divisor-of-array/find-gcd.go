/*
Given an integer array nums, return the greatest common divisor of the smallest number and largest number in nums.
The greatest common divisor of two numbers is the largest positive integer that evenly divides both numbers.
*/
package main

import "fmt"

func main() {
	nums := []int{2,3,4,6}
	fmt.Println(nums, "GCD =>", findGCD(nums));
}

func findGCD(nums []int) int {
	// pick first element as min & max value
	min, max := nums[0], nums[0]

	// find min, max
	for _, num := range nums {
		if num>max {
			max = num
		}
		if num<min {
			min = num
		}
	}

	// find greates common divisor
	gcd := 1
	for i:= min; i>0; i-- {
		if min % i ==0 && max % i == 0 && i>gcd {
			gcd = i
		}
	}
	return gcd
}