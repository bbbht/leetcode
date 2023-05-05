package main

// 322 coin-change 2023-05-04 17:51:14

// leetcode submit region begin(Prohibit modification and deletion)
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0

	size := len(coins)
	for i := 1; i <= amount; i++ {
		for j := 0; j < size; j++ {
			if coins[j] <= i {
				dp[i] = min322(dp[i], dp[i-coins[j]]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}

func min322(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
