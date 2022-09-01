package main

// 移除元素 2022-07-25 17:12:21

// leetcode submit region begin(Prohibit modification and deletion)
func removeElement(nums []int, val int) int {
	num := 0
	for _, item := range nums {
		if item != val {
			nums[num] = item
			num++
		}
	}

	return num
}

// leetcode submit region end(Prohibit modification and deletion)
