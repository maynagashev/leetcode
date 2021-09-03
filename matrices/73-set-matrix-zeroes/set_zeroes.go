/*
Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's, and return the matrix.
You must do it in place.

https://leetcode.com/problems/set-matrix-zeroes/
*/
package main

func main() {

}

/*
Marking first row + column, to store what cols & rows need to be changed.

Time Complexity : O(M×N)
Space Complexity : O(1)

1. We iterate all the cells and mark the corresponding first row/column cell in case of a cell with zero value.
2. We iterate the matrix we got from the above steps and mark respective cells zeroes.
 */
func setZeroes(matrix [][]int) {

	needToResetFirstRow := false
	needToResetFirstCol := false

	rows := len(matrix)
	cols := len(matrix[0])

	for i := 0; i < rows; i++ {
		if matrix[i][0] == 0 {
			needToResetFirstCol = true
		}
	}
	for i := 0; i < cols; i++ {
		if matrix[0][i] == 0 {
			needToResetFirstRow = true
		}
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if needToResetFirstCol {
		for i := 0; i < rows; i++ {
			matrix[i][0] = 0
		}
	}
	if needToResetFirstRow {
		for i := 0; i < cols; i++ {
			matrix[0][i] = 0
		}
	}
}
