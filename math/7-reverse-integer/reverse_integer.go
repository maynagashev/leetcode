package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(123, reverse(123))
}

func reverse(x int) int {
	s := 0
	for x != 0 {
		t := x % 10
		x = x / 10
		s =  s * 10 + t
	}

	//fmt.Println(math.MinInt32, math.MaxInt32)
	//fmt.Println(-2147483648, 2147483647)
	if s < math.MinInt32 || s > math.MaxInt32 {
		return 0
	}
	return s
}