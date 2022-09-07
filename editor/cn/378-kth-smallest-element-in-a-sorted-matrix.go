package main

import (
	"fmt"
)

// 378 kth-smallest-element-in-a-sorted-matrix 2022-09-07 14:09:50

// leetcode submit region begin(Prohibit modification and deletion)
func kthSmallest(matrix [][]int, k int) int {
	size := len(matrix)
	// 左上角最小，右下角最大
	left, right, mid := matrix[0][0], matrix[size-1][size-1], 0
	for left <= right {
		mid = (left + right) / 2
		if countLess378(matrix, mid) < k {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}

func countLess378(matrix [][]int, target int) int {
	n := len(matrix)
	x, y, res := n-1, 0, 0
	for x >= 0 && y < n {
		if matrix[x][y] <= target {
			res += x + 1
			y++
		} else {
			x--
		}
	}

	return res
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	// fmt.Println(countLess378([][]int{{1,5,9},{10,11,13},{12,13,15}}, 8))
	// fmt.Println(kthSmallest([][]int{{1,5,9},{10,11,13},{12,13,15}}, 8))
	fmt.Println(countLess378([][]int{{-5, -4}, {-5, -4}}, 2))
	fmt.Println(kthSmallest([][]int{{-5, -4}, {-5, -4}}, 2))
}
