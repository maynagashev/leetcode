/*
https://leetcode.com/problems/product-of-array-except-self/
*/
package main

func main() {

}

/*
Time complexity : O(N) where NN represents the number of elements in the input array. We use one iteration to construct the array L, one to update the array answer.
Space complexity : O(1) since don't use any additional array for our computations. The problem statement mentions that using the answer array doesn't add to the space complexity.
*/
func productExceptSelf0(nums []int) []int {
	n := len(nums)

	left := make([]int, n)
	left[0] = 1
	for i := 1; i < n; i++ {
		left[i] = left[i-1] * nums[i-1]
	}

	right := nums[n-1]
	for i := n - 2; i >= 0; i-- {
		left[i] *= right
		right *= nums[i]
		// fmt.Println(i, "left[i]", left[i], "right", right)
	}

	// fmt.Println(left)
	return left
}

func productExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, n)

	for i := range res {
		res[i] = 1
	}

	leftProduct, rightProduct := 1, 1
	l, r := 0, n-1

	for l < n && r > -1 {
		res[l] *= leftProduct
		res[r] *= rightProduct

		leftProduct *= nums[l]
		rightProduct *= nums[r]

		l++
		r--
	}

	return res
}
