package main

// 删除有序数组中的重复项 2022-07-25 17:11:16

// leetcode submit region begin(Prohibit modification and deletion)
func removeDuplicates(nums []int) int {
	num := 0
	for index, item := range nums {
		if index == 0 || nums[index-1] != item {
			nums[num] = item
			num++
		}
	}

	return num
}

// leetcode submit region end(Prohibit modification and deletion)
