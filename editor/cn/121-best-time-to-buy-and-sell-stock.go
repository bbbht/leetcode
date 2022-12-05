package main

// 121 best-time-to-buy-and-sell-stock 2022-11-25 14:23:31

// leetcode submit region begin(Prohibit modification and deletion)
func maxProfit(prices []int) int {
	max, buy := 0, 0
	for sell := 0; sell < len(prices); sell++ {
		// 当前价格更小了，更新 buy
		if prices[sell] < prices[buy] {
			buy = sell
		} else {
			max = max121(max, prices[sell]-prices[buy])

		}
	}
	return max
}

func max121(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
