package main

import (
	"fmt"
)

// 162 find-peak-element 2022-09-01 14:01:37

// leetcode submit region begin(Prohibit modification and deletion)
func findPeakElement(nums []int) int {
	size := len(nums)
	// 特殊情况
	if size == 2 {
		if nums[0] > nums[1] {
			return 0
		}
		return 1
	}

	left, right, mid := 0, size-1, 0
	for left < right {
		mid = (left + right) / 2
		// 到头了，且约束相邻数字不同
		if mid == 0 || mid == size-1 || nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
			return mid
		}

		// 如果小于前值，那么峰值肯定在前面
		if nums[mid] < nums[mid-1] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println([]int{1, 2, 1, 3, 5, 6, 4})
}
