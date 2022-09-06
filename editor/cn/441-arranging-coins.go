package main

import (
	"fmt"
)

// 441 arranging-coins 2022-09-06 14:12:09

// leetcode submit region begin(Prohibit modification and deletion)
func arrangeCoins(n int) int {
	left, right, mid := 0, n, 0
	for left <= right {
		mid = (left + right) / 2
		// 等差数列求和公式
		if mid*(mid+1)/2 <= n {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return right
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(1, arrangeCoins(1))
	fmt.Println(5, arrangeCoins(5))
	fmt.Println(10, arrangeCoins(10))
	fmt.Println(18, arrangeCoins(18))
}
