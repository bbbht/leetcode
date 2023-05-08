package main

// 52 n-queens-ii 2023-05-07 21:33:18

// leetcode submit region begin(Prohibit modification and deletion)

func totalNQueens(n int) int {
	res := 0
	// 只需要记下每行对应的列值即可
	queens := make([]int, n)
	for i := range queens {
		queens[i] = -1
	}

	solveNQueens52(queens, 0, &res)

	return res
}

// 每行只能放一个，逐行处理
func solveNQueens52(queens []int, row int, res *int) {
	if len(queens) == row {
		*res++
		return
	}

	for col := 0; col < len(queens); col++ {
		if isValid52(queens, row, col) {
			queens[row] = col
			solveNQueens52(queens, row+1, res)
			queens[row] = -1
		}
	}
}

func isValid52(queens []int, row, col int) bool {
	for i := 0; i < row; i++ {
		// 对角线上的行列差值的绝对值相等
		if col == queens[i] || abs52(row-i) == abs52(col-queens[i]) {
			return false
		}
	}

	return true
}

func abs52(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
