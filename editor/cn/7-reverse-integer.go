package main

import (
	"fmt"
	"math"
)

// 7 reverse-integer 2023-04-20 14:08:34

// leetcode submit region begin(Prohibit modification and deletion)
func reverse(x int) int {
	res := 0
	// 低位到高位
	for x != 0 {
		res = res*10 + x%10
		x /= 10
	}

	if res < math.MinInt32 || res > math.MaxInt32 {
		return 0
	}

	return res
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(reverse(9))
	fmt.Println(reverse(10))
	fmt.Println(reverse(-123))
}
