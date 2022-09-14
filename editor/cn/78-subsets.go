package main

import (
	"fmt"
)

// 78 subsets 2022-09-09 11:00:28

// leetcode submit region begin(Prohibit modification and deletion)
func subsets(nums []int) [][]int {
	// 全不重复，所以子集值也不重复，总共有 n = 2^len(nums) 种，二进制从 0 - n 正好覆盖所有组合
	n := 1 << len(nums)
	res := make([][]int, n)
	sub := make([]int, 0, n)
	for i := 0; i < n; i++ {
		for j := range nums {
			if (1 << j) & i > 0 {
				sub = append(sub, nums[j])
			}
		}

		tmp := make([]int, len(sub))
		copy(tmp, sub)
		sub = sub[:0]

		res[i] = tmp
	}

	return res
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(subsets([]int{1,2,3}))
}
