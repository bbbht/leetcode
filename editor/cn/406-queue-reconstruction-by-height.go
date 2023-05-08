package main

import (
	"fmt"
	"sort"
)

// 406 queue-reconstruction-by-height 2023-05-07 22:06:38

// leetcode submit region begin(Prohibit modification and deletion)
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		return people[i][0] > people[j][0] || people[i][0] == people[j][0] && people[i][1] < people[j][1]
	})

	// fmt.Println(people)

	res := make([][]int, 0, len(people))
	for _, pair := range people {
		idx := pair[1]
		res = append(res[:idx], append([][]int{pair}, res[idx:]...)...)
	}

	return res
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(reconstructQueue([][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}))
}
