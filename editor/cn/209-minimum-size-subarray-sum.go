package main

import (
	"fmt"
)

// 209 minimum-size-subarray-sum 2022-09-01 17:34:52

// leetcode submit region begin(Prohibit modification and deletion)
func minSubArrayLen(target int, nums []int) int {
	sum, size, left, right := 0, len(nums), 0, 0
	minN := size + 1
	// 窗口法，保留窗口内部数据，移动时加减变动的值
	for right < size {
		for ; sum < target && right < size; right++ { // 增加长度以满足
			if nums[right] >= target {
				return 1
			}
			sum += nums[right]
		}

		for ; sum >= target; left++ { // 缩短直至不满足
			// 取最小值，right是预加 1 的，所以直接减
			minN = min209(minN, right-left)
			sum -= nums[left]
		}
	}

	if minN > size {
		return 0
	}

	return minN
}

func min209(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(minSubArrayLen(11, []int{12}))
}
