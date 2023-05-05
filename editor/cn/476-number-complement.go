package main

import (
	"fmt"
)

// 476 number-complement 2022-12-12 10:53:43

// leetcode submit region begin(Prohibit modification and deletion)
func findComplement(num int) int {
	res := 0
	for n := 0; num > 0; n++ {
		if num&1 == 0 {
			res += 1 << n
		}
		num >>= 1
	}

	return res
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(findComplement(5))
	fmt.Println(findComplement(10))
	fmt.Println(findComplement(99))
}
