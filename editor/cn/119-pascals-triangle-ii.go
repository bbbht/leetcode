package main

import (
	"fmt"
)

// 119 pascals-triangle-ii 2022-11-24 10:15:31

// leetcode submit region begin(Prohibit modification and deletion)
func getRow(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	row[0] = 1
	for i := 0; i < rowIndex; i++ {
		// 倒序
		for j := i; j >= 1; j-- {
			row[j] = row[j] + row[j-1]
		}
		row[i+1] = 1
	}

	return row
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(getRow(3))
}
