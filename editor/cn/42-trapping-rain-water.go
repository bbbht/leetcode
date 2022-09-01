package main

import (
	"fmt"
)

// 接雨水 2022-07-29 16:43:02

// leetcode submit region begin(Prohibit modification and deletion)
func trap(height []int) int {
	result := 0
	if len(height) <= 1 {
		return 0
	}
	// 重点在判断第二高，最高端点不动，水量为第二高（短板）减去此处高度
	secHigh, left, right := 0, 0, len(height)-1

	for left < right {
		if height[left] < height[right] {
			secHigh = max42(secHigh, height[left])
			result += secHigh - height[left]
			left++
		} else {
			secHigh = max42(secHigh, height[right])
			result += secHigh - height[right]
			right--
		}
	}

	return result
}

func max42(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(trap([]int{0,1,0,2,1,5,1,3,2,1,4,1}))
}
