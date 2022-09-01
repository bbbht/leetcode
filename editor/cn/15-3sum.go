package main

import (
	"sort"
)

// 三数之和 2022-07-20 10:05:24

// leetcode submit region begin(Prohibit modification and deletion)
// func main() {
// 	fmt.Println(threeSum([]int{1, -1, -1, 0}))
// 	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
// }
func threeSum(nums []int) [][]int {
	var res [][]int

	size, sum := len(nums), 0
	if size < 3 {
		return res
	}

	// 排序
	sort.Ints(nums)

	for i := range nums {
		l, r := i+1, size-1
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for l < r {
			sum = nums[i] + nums[l] + nums[r]

			if sum == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				// 避免出现重复
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				for l < r && nums[l] == nums[l+1] {
					l++
				}
			}

			if sum >= 0 {
				r--
			}

			if sum <= 0 {
				l++
			}
		}
	}
	return res
}

// leetcode submit region end(Prohibit modification and deletion)
