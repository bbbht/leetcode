package main

// 64 minimum-path-sum 2022-10-18 13:57:17

// leetcode submit region begin(Prohibit modification and deletion)
func minPathSum(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	dp := make([][]int, rows)
	for i := range dp {
		dp[i] = make([]int, cols)
	}

	// 初始化
	dp[0][0] = grid[0][0]
	// 第一列，只能从上至下
	for i := 1; i < rows; i++ {
		dp[i][0] = grid[i][0] + dp[i-1][0]
	}
	// 第一行，只能从左至右
	for i := 1; i < cols; i++ {
		dp[0][i] = grid[0][i] + dp[0][i-1]
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			dp[i][j] = grid[i][j] + min64(dp[i-1][j], dp[i][j-1])
		}
	}

	return dp[rows-1][cols-1]
}

func min64(a, b int) int {
	if a > b {
		return b
	}

	return a
}
// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
