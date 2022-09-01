package main

import (
	"sort"
)

// 四数之和 2022-07-22 15:02:39

// leetcode submit region begin(Prohibit modification and deletion)
func fourSum(nums []int, target int) [][]int {
	size := len(nums)
	if size < 4 {
		return nil
	}

	sort.Ints(nums)

	res := make([][]int, 0)
	// 复用三数之和逻辑，即在三数之和外加一层遍历即可，求三数之和为 target-n

	sum, newTarget := 0, 0
	for k := range nums {
		newTarget = target - nums[k]
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}

		for i := k + 1; i < size-2; i++ {
			l, r := i+1, size-1
			if i > k+1 && nums[i] == nums[i-1] {
				continue
			}
			for l < r {
				sum = nums[i] + nums[l] + nums[r]

				if sum == newTarget {
					res = append(res, []int{nums[k], nums[i], nums[l], nums[r]})
					// 避免出现重复
					for l < r && nums[r] == nums[r-1] {
						r--
					}
					for l < r && nums[l] == nums[l+1] {
						l++
					}
				}

				if sum >= newTarget {
					r--
				}

				if sum <= newTarget {
					l++
				}
			}
		}
	}

	return res
}

// leetcode submit region end(Prohibit modification and deletion)
