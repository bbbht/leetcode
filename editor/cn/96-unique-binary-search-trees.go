package main

// 96 unique-binary-search-trees 2023-05-05 17:14:38

// leetcode submit region begin(Prohibit modification and deletion)
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-j-1]
		}
	}

	return dp[n]
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
