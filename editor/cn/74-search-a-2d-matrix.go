package main

import (
	"fmt"
)

// 74 search-a-2d-matrix 2022-08-31 17:07:19

// leetcode submit region begin(Prohibit modification and deletion)
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	size := len(matrix[0])

	// 确定行
	row, mid, up, down := -1, -1, 0, len(matrix)-1
	for up <= down {
		mid = (up + down) / 2
		if matrix[mid][0] <= target && matrix[mid][size-1] >= target {
			row = mid
			break
		}
		if matrix[mid][0] < target {
			up = mid + 1
		} else {
			down = mid - 1
		}
	}
	fmt.Println(row)
	if row < 0 {
		return false
	}

	left, right := 0, size-1
	for left <= right {
		mid = (left + right) / 2
		if matrix[row][mid] == target {
			return true
		}
		if matrix[row][mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(searchMatrix([][]int{{1,3,5,7},{10,11,16,20},{23,30,34,60}}, 13))
}
