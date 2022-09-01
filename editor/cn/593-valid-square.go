package main

import (
	"sort"
)

// 有效的正方形 2022-07-29 16:46:55

// leetcode submit region begin(Prohibit modification and deletion)
func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	// 计算任意两点的距离（实际是平方），应该是四个边等长，两个对角线等长
	diaLens := make([]int, 6)
	diaLens[0] = diaLen(p1, p2)
	diaLens[1] = diaLen(p1, p3)
	diaLens[2] = diaLen(p1, p4)
	diaLens[3] = diaLen(p2, p3)
	diaLens[4] = diaLen(p2, p4)
	diaLens[5] = diaLen(p3, p4)
	sort.Ints(diaLens)

	// 如果前四条边相等，后两条边相等，且前面的边长小于对角线长，即可组成正方形。
	return diaLens[0] == diaLens[1] && diaLens[0] == diaLens[2] && diaLens[0] == diaLens[3] && diaLens[4] == diaLens[5] && diaLens[3] < diaLens[4]
}

// 计算边长，未开方
func diaLen(p1, p2 []int) int {
	return (p1[0]-p2[0])*(p1[0]-p2[0]) + (p1[1]-p2[1])*(p1[1]-p2[1])
}

// leetcode submit region end(Prohibit modification and deletion)
