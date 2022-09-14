package main

import (
	"fmt"
)

// 54 spiral-matrix 2022-09-13 15:19:59

// leetcode submit region begin(Prohibit modification and deletion)
func spiralOrder(matrix [][]int) []int {
	rows, cols := len(matrix), len(matrix[0])
	size := rows * cols
	res := make([]int, 0, size)

	left, right, up, down := 0, cols-1, 0, rows-1
	for {
		for i := left; i <= right; i++ { // ➡
			res = append(res, matrix[up][i])
			if len(res) == size {
				return res
			}
		}
		up++
		for i := up; i <= down; i++ { // ⬇
			res = append(res, matrix[i][right])
			if len(res) == size {
				return res
			}
		}
		right--
		for i := right; i >= left; i-- { // ️⬅
			res = append(res, matrix[down][i])
			if len(res) == size {
				return res
			}
		}
		down--
		for i := down; i >= up; i-- { // ⬆
			res = append(res, matrix[i][left])
			if len(res) == size {
				return res
			}
		}
		left++
	}
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}
