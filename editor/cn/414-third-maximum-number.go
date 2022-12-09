package main

import (
	"math"
)

// 414 third-maximum-number 2022-12-08 15:46:08

// leetcode submit region begin(Prohibit modification and deletion)
func thirdMax(nums []int) int {
	a, b, c := math.MinInt64, math.MinInt64, math.MinInt64
	for _, num := range nums {
		if num > a {
			a, b, c = num, a, b
		} else if a > num && num > b {
			b, c = num, b
		} else if b > num && num > c {
			c = num
		}
	}
	if c == math.MinInt64 {
		return a
	}

	return c
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
