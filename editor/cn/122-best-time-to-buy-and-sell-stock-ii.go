package main

// 122 best-time-to-buy-and-sell-stock-ii 2022-11-25 16:52:27

// leetcode submit region begin(Prohibit modification and deletion)
func maxProfit(prices []int) int {
	sub, profit := 0, 0
	// 涨了就前一天买入即可
	for i := 1; i < len(prices); i++ {
		sub = prices[i] - prices[i-1]
		if sub > 0 {
			profit += sub
		}
	}

	return profit
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
