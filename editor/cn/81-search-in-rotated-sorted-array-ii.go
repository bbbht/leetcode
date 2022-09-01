package main

import (
	"fmt"
)

// 81 search-in-rotated-sorted-array-ii 2022-09-01 09:52:06

// leetcode submit region begin(Prohibit modification and deletion)
func search(nums []int, target int) bool {
	size := len(nums)
	if size == 0 {
		return false
	}
	if target == nums[0] {
		return true
	}

	left, right, mid := 0, size-1, 0
	for left <= right {
		// 相比33 题，只需要额外判断并跳过left==right的情况，以确保 left < right 恒成立即22的情况
		for left < right && nums[left] == nums[right] {
			left += 1
		}
		mid = (left + right) / 2
		if nums[mid] == target {
			return true
		}

		if nums[left] <= nums[mid] { // mid 未超过翻转点
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
				continue
			}
			left = mid + 1
		} else { // mid 超过了翻转点，较小的递增区间
			if nums[right] >= target && target > nums[mid] {
				left = mid + 1
				continue
			}
			right = mid - 1
		}
	}

	return false
}
// leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(search([]int{2,5,6,0,0,1,2}, 2))
}
