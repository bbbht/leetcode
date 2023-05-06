package main

import (
	"sort"
)

// 215 kth-largest-element-in-an-array 2023-05-06 12:16:41

// leetcode submit region begin(Prohibit modification and deletion)
func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
