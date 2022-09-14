package main

import (
	"fmt"
	"sort"
)

// 56 merge-intervals 2022-09-13 17:37:14

// leetcode submit region begin(Prohibit modification and deletion)
func merge(intervals [][]int) [][]int {
	size := len(intervals)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := make([][]int, 0, size/2)
	for i := 0; i < size; i++ {
		tmp := []int{intervals[i][0], intervals[i][1]}
		for i+1 < size {
			// 判断是否能与下一个合并
			if tmp[1] >= intervals[i+1][1] || tmp[1] >= intervals[i+1][0] {
				if intervals[i+1][1] > tmp[1] {
					tmp[1] = intervals[i+1][1]
				}
				i++
				continue
			}
			break
		}

		res = append(res, tmp)
	}

	return res
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(merge([][]int{{1, 4}, {1, 5}}))
	fmt.Println(merge([][]int{{3, 4}, {5, 6}}))
	fmt.Println(merge([][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}))
	fmt.Println(merge([][]int{{1, 4}, {0, 2}, {3, 5}}))
}
