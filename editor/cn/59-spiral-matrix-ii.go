package main

import (
	"fmt"
)

// 59 spiral-matrix-ii 2022-09-14 14:02:22

// leetcode submit region begin(Prohibit modification and deletion)
func generateMatrix(n int) [][]int {
	up, down, left, right, count := 0, n-1, 0, n-1, 1
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}

	for {
		for i := left; i <= right; i++ {
			res[up][i] = count
			count++
		}
		up++
		if up > down {
			break
		}

		for i := up; i <= down; i++ {
			res[i][right] = count
			count++
		}
		right--
		if right < left {
			break
		}

		for i := right; i >= left; i-- {
			res[down][i] = count
			count++
		}
		down--
		if down < up {
			break
		}

		for i := down; i >= up; i-- {
			res[i][left] = count
			count++
		}
		left++
		if left > right {
			break
		}
	}
	return res
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(generateMatrix(1))
	fmt.Println(generateMatrix(3))
	fmt.Println(generateMatrix(5))
}
