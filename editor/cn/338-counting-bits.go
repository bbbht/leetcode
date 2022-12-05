package main

import (
	"fmt"
)

// 338 counting-bits 2022-12-02 15:31:02

// leetcode submit region begin(Prohibit modification and deletion)
func countBits(n int) []int {
	res := make([]int, n+1)
	for i := 1; i <= n; i++ {
		// i&(i-1)对应的值的 1 的数量 + 1
		res[i] = res[i&(i-1)] + 1
	}

	return res
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	for i := 0; i < 16; i++ {
		fmt.Printf("%04b -><- %04b \n", i, i&(i-1))
	}
}
