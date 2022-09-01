package main

import (
	"sort"
)

// 组合总和 2022-07-28 16:53:50

// leetcode submit region begin(Prohibit modification and deletion)
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	path, res := make([]int, 0), make([][]int, 0)

	combination39(0, target, candidates, path, &res)

	return res
}

func combination39(index, target int, candidates []int, path []int, res *[][]int) {
	// 此次遍历完成
	if target == 0 {
		// [:]的行为太复杂，直接copy最安全
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}

	// 肯定不会有组合了
	if target < 0 {
		return
	}

	// 回溯了，所有可能来一遍
	for i := index; i < len(candidates); i++ {
		path = append(path, candidates[i])
		combination39(i, target-candidates[i], candidates, path, res)
		path = path[:len(path)-1]
	}
}

// leetcode submit region end(Prohibit modification and deletion)
