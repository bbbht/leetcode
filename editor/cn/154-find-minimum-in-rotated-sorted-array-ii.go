package main

import (
	"fmt"
)

// 154 find-minimum-in-rotated-sorted-array-ii 2022-09-01 11:35:05

// leetcode submit region begin(Prohibit modification and deletion)
func findMin(nums []int) int {
	size := len(nums)
	if nums[0] < nums[size-1] {
		return nums[0]
	}

	mid, left, right := 0, 0, size-1
	for right > 0 && nums[left] == nums[right] {
		right--
	}
	if right == 0  || nums[right] > nums[0]{
		return nums[0]
	}

	for left < right - 1{
		mid = (left + right) / 2
		if nums[mid] >= nums[0] {
			left = mid
		} else {
			right = mid
		}
	}

	return nums[right]
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(findMin([]int{3,1,3,3}))
	fmt.Println(findMin([]int{1,2,1}))
}
