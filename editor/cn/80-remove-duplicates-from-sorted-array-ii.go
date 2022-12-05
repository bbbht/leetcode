package main

import (
	"fmt"
)

// 80 remove-duplicates-from-sorted-array-ii 2022-11-03 17:20:41

// leetcode submit region begin(Prohibit modification and deletion)
func removeDuplicates(nums []int) int {
	cur := 0
	for i := range nums {
		if cur < 2 || nums[i] > nums[cur-2] {
			nums[cur] = nums[i]
			cur++
		}
	}

	return cur
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(removeDuplicates([]int{1,1,1,2,2,3}))
	// fmt.Println(removeDuplicates([]int{0,0,1,1,1,1,2,3,3}))
}
