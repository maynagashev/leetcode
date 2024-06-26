package main

import "fmt"

func main() {
	fmt.Printf("Primes (up to 100): %#v\n", eratosthenes(100))
}

func eratosthenes(n int) []bool {
	numbers := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		numbers[i] = true
	}
	for num := 2; num < n; num++ {
		if numbers[num] {
			for j := 2 * num; j <= n; j += num {
				numbers[j] = false
			}
		}
	}
	numbers[0] = false
	numbers[1] = false
	return numbers
}
