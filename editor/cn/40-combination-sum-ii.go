package main

import (
	"fmt"
	"sort"
)

// 组合总和 II 2022-07-29 15:14:00

// leetcode submit region begin(Prohibit modification and deletion)
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	path, res := make([]int, 0), make([][]int, 0)

	combination40(0, target, candidates, path, &res)

	return res
}

func combination40(index, target int, candidates []int, path []int, res *[][]int) {
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
		combination40(i+1, target-candidates[i], candidates, path, res)
		path = path[:len(path)-1]

		for i+1 < len(candidates) && candidates[i] == candidates[i+1] {
			i++
		}
	}
}

func main() {
	fmt.Println(combinationSum2([]int{1, 1, 1, 1, 2, 2, 2, 3, 4, 5}, 3))
}

// leetcode submit region end(Prohibit modification and deletion)
