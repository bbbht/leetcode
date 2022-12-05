package main

import (
	"fmt"
	"sort"
)

// 90 subsets-ii 2022-11-07 15:18:09

// leetcode submit region begin(Prohibit modification and deletion)
func subsetsWithDup(nums []int) [][]int {
	// 排序，将重复的紧邻
	sort.Ints(nums)

	n := 1 << len(nums)
	res := make([][]int, 0, n)
	sub := make([]int, 0, n)
	for i := 0; i < n; i++ {
		skip := false
		for j := range nums {
			// 重复的，只保留最开始连续出现的重复数字（对应二进制标记为1）
			if (i>>j)&1 == 1 {
				// 非首位，且发生了重复，且前一位为 0 （不连续）
				if j > 0 && nums[j] == nums[j-1] && (i>>(j-1)&1) == 0 {
					skip = true
					break
				} else {
					sub = append(sub, nums[j])
				}
			}
		}

		if !skip {
			tmp := make([]int, len(sub))
			copy(tmp, sub)

			res = append(res, tmp)
		}

		sub = sub[:0]
	}

	return res
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}
