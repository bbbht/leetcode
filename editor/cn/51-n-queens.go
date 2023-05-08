package main

import (
	"strings"
)

// 51 n-queens 2023-05-07 20:53:02

// leetcode submit region begin(Prohibit modification and deletion)
func solveNQueens(n int) [][]string {
	var res [][]string
	queens := make([][]string, n)
	for i := range queens {
		queens[i] = make([]string, n)
		for j := range queens[i] {
			queens[i][j] = "."
		}
	}

	solveNQueens51(queens, 0, &res)

	return res
}

// 每行只能放一个，逐行处理
func solveNQueens51(queens [][]string, row int, res *[][]string) {
	if len(queens) == row {
		*res = append(*res, generate51(queens))
		return
	}

	for i := 0; i < len(queens); i++ {
		if isValid51(queens, row, i) {
			queens[row][i] = "Q"
			solveNQueens51(queens, row+1, res)
			queens[row][i] = "."
		}
	}
}

func generate51(queens [][]string) []string {
	res := make([]string, len(queens))
	for i := 0; i < len(queens); i++ {
		res[i] = strings.Join(queens[i], "")
	}

	return res
}

func isValid51(queens [][]string, row, col int) bool {
	// 上方列
	for i := 0; i < row; i++ {
		if queens[i][col] == "Q" {
			return false
		}
	}

	// 左上方斜线
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if queens[i][j] == "Q" {
			return false
		}
	}

	// 右上方斜线
	for i, j := row-1, col+1; i >= 0 && j < len(queens); i, j = i-1, j+1 {
		if queens[i][j] == "Q" {
			return false
		}
	}

	return true
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
