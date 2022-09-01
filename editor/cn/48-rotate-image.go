package main

import (
	"fmt"
)

// 48 rotate-image 2022-08-31 09:58:05

// leetcode submit region begin(Prohibit modification and deletion)
func rotate(matrix [][]int) {
	size := len(matrix)
	if size == 0 {
		return
	}
	// 先上下翻转
	for row := 0; row < size/2; row++ { // 行只需要遍历一半
		for col := 0; col < size; col++ { // 遍历每一列
			matrix[row][col], matrix[size-row-1][col] = matrix[size-row-1][col], matrix[row][col]
		}
	}
	// 沿左上-右下对角线翻转
	for row := 0; row < size; row++ {
		for col := 0; col < row; col++ {
			matrix[row][col], matrix[col][row] = matrix[col][row], matrix[row][col]
		}
	}
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	i := [][]int{{1,2,3},{4,5,6},{7,8,9}}
	rotate(i)
	fmt.Println(i)
}
