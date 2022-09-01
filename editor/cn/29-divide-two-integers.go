package main

import (
	"math"
)

// 两数相除 2022-07-26 18:55:31

// leetcode submit region begin(Prohibit modification and deletion)
func divide(dividend int, divisor int) int {
	res := 0
	isPos := (dividend < 0) == (divisor < 0)
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}

	for dividend >= divisor {
		tmp, n := divisor, 1

		// 逐级逼近，除数2倍增长直至大于被除数的余数，然后循环处理余数，直至被除数的余数小于除数
		for dividend >= tmp {
			dividend -= tmp

			res += n

			n <<= 1
			tmp <<= 1
		}

	}

	if !isPos {
		res = -res
	}

	if res > math.MaxInt32 || res < math.MinInt32 {
		return math.MaxInt32
	}

	return res
}

// func main() {
// 	// fmt.Println(divide(2147483647, -3))
// 	fmt.Println(divide(31, 2))
// 	fmt.Println(divide(10, 3))
// }

// leetcode submit region end(Prohibit modification and deletion)
