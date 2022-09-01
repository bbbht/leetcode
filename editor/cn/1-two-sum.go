package main

// 两数之和 2022-07-06 15:31:41

// leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, n := range nums {
		if j, ok := m[target - n]; ok {
			return []int{i, j}
		}

		m[n] = i
	}

	return []int{}
}

// leetcode submit region end(Prohibit modification and deletion)
