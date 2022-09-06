package main

import (
	"fmt"
)

// 240 search-a-2d-matrix-ii 2022-09-02 14:00:30

// leetcode submit region begin(Prohibit modification and deletion)
func searchMatrix(matrix [][]int, target int) bool {
	rows, cols := len(matrix), len(matrix[0])
	if target < matrix[0][0] || target > matrix[rows-1][cols-1] {
		return false
	}
	// 从左下角开始判断，左下角是列的最大值，行的最小值，向右上角移动靠近目标值
	row, col := rows-1, 0
	for row >= 0 && col < cols {
		if matrix[row][col] > target {
			row--
		} else if matrix[row][col] < target {
			col++
		} else {
			return true
		}
	}

	return false
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(searchMatrix([][]int{{-5}}, -5))
}
