package main

// 79 word-search 2022-10-19 11:32:38

// leetcode submit region begin(Prohibit modification and deletion)
func exist(board [][]byte, word string) bool {
	rows, cols := len(board), len(board[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if searchWord(board, word, 0, i, j) {
				return true
			}
		}
	}

	return false
}

func searchWord(board [][]byte, word string, index, row, col int) bool {
	if index >= len(word) {
		return true
	}
	if row < 0 || col < 0 || row >= len(board) || col >= len(board[0]) || board[row][col] != word[index] {
		return false
	}
	tmp := board[row][col]
	board[row][col] = ' '
	ok := searchWord(board, word, index+1, row+1, col) ||
		searchWord(board, word, index+1, row-1, col) ||
		searchWord(board, word, index+1, row, col+1) ||
		searchWord(board, word, index+1, row, col-1)
	board[row][col] = tmp

	return ok
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
