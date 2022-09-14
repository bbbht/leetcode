package main

import (
	"math"
)

// 363 max-sum-of-rectangle-no-larger-than-k 2022-09-08 09:46:44

// leetcode submit region begin(Prohibit modification and deletion)
func maxSumSubmatrix(matrix [][]int, k int) int {
	rows, cols, res := len(matrix), len(matrix[0]), math.MinInt64
	// 计算每一行进行逐列累加后的和，覆盖原数据（最终得到每行的前缀和）
	for i := 0; i < rows; i++ {
		for j := 1; j < cols; j++ {
			matrix[i][j] += matrix[i][j-1]
		}
	}
	for i := 0; i < cols; i++ {
		for j := i; j < cols; j++ {
			sum := 0
			for r := 0; r < rows; r++ {
				if i == 0 {
					sum += matrix[r][j]
				} else {
					sum += matrix[r][j] - matrix[r][i]
				}
				if sum == k {
					return sum
				}
				if sum < k {
					res = max363(res, sum)
				}
			}
		}
	}

	return res
}

func max363(a, b int) int {
	if a > b  {
		return a
	}
	return b
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
