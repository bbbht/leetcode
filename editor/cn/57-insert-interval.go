package main

import (
	"fmt"
)

// 57 insert-interval 2022-09-14 10:21:04

// leetcode submit region begin(Prohibit modification and deletion)
func insert(intervals [][]int, newInterval []int) [][]int {
	pos := len(intervals)
	res := make([][]int, 0, pos+1)

	for i := range intervals {
		if newInterval[0] < intervals[i][0] {
			pos = i
			break
		}
	}

	intervals = append(intervals[:pos], append([][]int{newInterval}, intervals[pos:]...)...)

	for _, interval := range intervals {
		if len(res) == 0 || interval[0] > res[len(res)-1][1] {
			res = append(res, interval)
		} else {
			res[len(res)-1][1] = max57(res[len(res)-1][1], interval[1])
		}
	}
	return res
}

func max57(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))
	fmt.Println(insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}))
}
