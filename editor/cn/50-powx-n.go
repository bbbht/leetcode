package main

import (
	"fmt"
)

// 50 powx-n 2022-08-31 16:15:41

// leetcode submit region begin(Prohibit modification and deletion)
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n < 0 {
		x = 1 / x
		n = -n
	}

	if n%2 != 0 {
		return x * myPow(x, n-1)
	}
	return myPow(x*x, n/2)
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(myPow(2, 5))
}
