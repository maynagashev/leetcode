// https://leetcode.com/problems/word-search/
package main

func main() {

}

func exist(board [][]byte, word string) bool {
	for i:=0; i<len(board); i++ {
		for j:=0; j<len(board[0]); j++ {
			if board[i][j] == word[0] {
				if backtrack(board, word, i, j) == true {
					return true
				}
			}
		}
	}
	return false
}

func backtrack(board [][]byte, word string, i int, j int) bool {
	if len(word) == 0 {
		return true
	}

	if i < 0 || i>=len(board) || j<0 || j>=len(board[0]) {
		return false
	}

	if board[i][j] != word[0] {
		return false
	}

	originByte := board[i][j]
	board[i][j] = '_'

	match := backtrack(board, word[1:], i+1, j) || backtrack(board, word[1:], i-1, j) ||
		backtrack(board, word[1:], i, j+1) || backtrack(board, word[1:], i, j-1)


	board[i][j] = originByte

	return match
}