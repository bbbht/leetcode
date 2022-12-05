package main

import (
	"fmt"
)

// 202 happy-number 2022-11-29 14:31:45

// leetcode submit region begin(Prohibit modification and deletion)
func isHappy(n int) bool {
	m := make(map[int]struct{})

	for {
		sum := 0
		for n > 0 {
			sum += (n % 10) * (n % 10)
			n /= 10
		}
		if sum == 1 {
			return true
		}
		if _, ok := m[sum]; ok {
			return false
		}
		m[sum] = struct{}{}

		n = sum
	}
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(isHappy(1))
	fmt.Println(isHappy(0))
	fmt.Println(isHappy(19))
}
