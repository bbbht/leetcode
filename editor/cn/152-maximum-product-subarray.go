package main

import (
	"fmt"
)

// 152 maximum-product-subarray 2023-05-06 11:51:41

// leetcode submit region begin(Prohibit modification and deletion)
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxNum, minNum, res := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		max, min := maxNum, minNum
		maxNum = max152(nums[i], max152(max*nums[i], min*nums[i]))
		minNum = min152(nums[i], min152(max*nums[i], min*nums[i]))

		res = max152(res, maxNum)
	}

	return res
}

func max152(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min152(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(maxProduct([]int{-4, -3, -2}))
}
