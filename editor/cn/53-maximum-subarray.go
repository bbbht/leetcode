package main

import (
	"fmt"
)

// 53 maximum-subarray 2022-08-31 17:04:49

// leetcode submit region begin(Prohibit modification and deletion)
func maxSubArray(nums []int) int {
	sum, res := 0, nums[0]
	for i := range nums {
		// 不断的滚动累加，直到累加值小于当前值，从头开始
		// 判断依据为如果加上当前值反而变小了，说明前面的累加已经是负向的，可以丢弃从当前值重新开始
		sum = max53(sum+nums[i], nums[i])
		// 每一次累加后判断是否出现了更大值
		res = max53(sum, res)
	}

	return res
}

func max53(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(maxSubArray([]int{-1,1,2,3,-4,1,5}))
}
