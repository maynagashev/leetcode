// https://leetcode.com/problems/spiral-matrix/
package main

func main() {

}

func spiralOrder(matrix [][]int) []int {
	var res []int
	top := 0
	left := 0
	right := len(matrix[0]) - 1
	bottom := len(matrix) - 1
	direction := 0

	for top <= bottom && left <= right {
		switch direction {
		case 0:
			for i := left; i <= right; i++ {
				res = append(res, matrix[top][i])
			}
			top++
			direction = 1
		case 1:
			for i := top; i <= bottom; i++ {
				res = append(res, matrix[i][right])
			}
			right--
			direction = 2
		case 2:
			for i := right; i >= left; i-- {
				res = append(res, matrix[bottom][i])
			}
			bottom--
			direction = 3
		case 3:
			for i := bottom; i >= top; i-- {
				res = append(res, matrix[i][left])
			}
			left++
			direction = 0
		}
	}
	return res
}
