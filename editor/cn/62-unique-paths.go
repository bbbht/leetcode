package main

// 62 unique-paths 2022-10-12 17:40:03

// leetcode submit region begin(Prohibit modification and deletion)
func uniquePaths(m int, n int) int {
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] += dp[j-1]
		}
	}

	return dp[n-1]
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
