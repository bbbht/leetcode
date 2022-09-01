package main

import (
	"fmt"
)

// 46 permutations 2022-08-30 17:00:44

// leetcode submit region begin(Prohibit modification and deletion)
func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	res := make([][]int, 0, factorial46(len(nums)))

	bfs46(nums, 0, len(nums)-1, &res)

	return res
}

func bfs46(nums []int, begin, end int, res *[][]int) {
	if begin == end {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*res = append(*res, tmp)
		return
	}

	for i := begin; i <= end; i++ {
		nums[begin], nums[i] = nums[i], nums[begin]
		bfs46(nums, begin+1, end, res)
		nums[begin], nums[i] = nums[i], nums[begin]
	}
}

func factorial46(n int) int {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}

	return res
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(permute([]int{1,2,3}))
}
