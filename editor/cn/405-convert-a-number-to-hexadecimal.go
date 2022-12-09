package main

import (
	"fmt"
)

// 405 convert-a-number-to-hexadecimal 2022-12-08 14:26:47

// leetcode submit region begin(Prohibit modification and deletion)
var m405 = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}

func toHex(num int) string {
	// 题目说明是32位
	if num < 0 {
		num += 1 << 32
	}

	res := make([]byte, 0, 8)
	for i, mask := range []int{0xf0000000, 0x0f000000, 0x00f00000, 0x000f0000, 0x0000f000, 0x00000f00, 0x000000f0, 0x0000000f} {
		tmp := (num & mask) >> ((7 - i) << 2)
		if tmp != 0 || len(res) > 0 || i == 7 {
			res = append(res, m405[tmp])
		}
	}

	return string(res)
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(toHex(0))
	fmt.Println(toHex(15))
	fmt.Println(toHex(16))
	fmt.Println(toHex(17))
	fmt.Println(toHex(-2))
}
