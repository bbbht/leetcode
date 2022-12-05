package main

// 63 unique-paths-ii 2022-10-13 16:24:41

// leetcode submit region begin(Prohibit modification and deletion)
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	rows, cols := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, rows+1)
	for i := range dp {
		dp[i] = make([]int, cols+1)
	}

	if obstacleGrid[0][0] != 1 {
		dp[1][1] = 1
	}

	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			if i == 1 && j == 1 {
				continue
			}
			if obstacleGrid[i-1][j-1] != 1 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[rows][cols]
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
