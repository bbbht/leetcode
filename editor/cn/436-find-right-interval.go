package main

import (
	"fmt"
	"sort"
)

// 436 find-right-interval 2022-09-07 10:24:25

// leetcode submit region begin(Prohibit modification and deletion)
func findRightInterval(intervals [][]int) []int {
	size := len(intervals)

	starts, res := make([][2]int, size), make([]int, 0, size)
	// 创建一个新结构，保存start与对应的索引
	for i := range intervals {
		starts[i] = [2]int{intervals[i][0], i}
	}

	sort.Slice(starts, func(i, j int) bool {
		return starts[i][0] < starts[j][0]
	})

	for i := range intervals {
		res = append(res, findBigger436(starts, intervals[i][1]))
	}

	return res
}

func findBigger436(starts [][2]int, n int) int {
	left, right, mid := 0, len(starts)-1, 0
	for left < right {
		mid = (left + right) / 2
		if starts[mid][0] >= n {
			right = mid
		} else {
			left = mid + 1
		}
	}

	if starts[left][0] >= n {
		return starts[left][1]
	}

	return -1
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(findRightInterval([][]int{{1,2}}))
}
