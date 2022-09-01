package main

import (
	"fmt"
	"sort"
)

// 47 permutations-ii 2022-08-31 10:58:06

// leetcode submit region begin(Prohibit modification and deletion)
func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	sort.Ints(nums)
	res := make([][]int, 0, factorial47(len(nums)))

	bfs47(nums, 0, len(nums)-1, &res)

	return res
}

func bfs47(nums []int, begin, end int, res *[][]int) {
	if begin == end {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*res = append(*res, tmp)
		return
	}

	m := make(map[int]struct{}) // 当前位置，如果一个数字已经交换过了，则跳过
	for i := begin; i <= end; i++ {
		if _, ok := m[nums[i]]; ok {
			continue
		}
		m[nums[i]] = struct{}{}
		nums[begin], nums[i] = nums[i], nums[begin]
		bfs47(nums, begin+1, end, res)
		nums[begin], nums[i] = nums[i], nums[begin]
	}
}

func factorial47(n int) int {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}

	return res
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(permuteUnique([]int{1, 2, 2}))
}
