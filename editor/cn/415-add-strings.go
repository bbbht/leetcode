package main

import (
	"fmt"
)

// 415 add-strings 2022-12-08 16:24:01

// leetcode submit region begin(Prohibit modification and deletion)
func addStrings(num1 string, num2 string) string {
	len1, len2 := len(num1), len(num2)
	size := len1
	if len2 > size {
		size = len2
	}

	res := make([]byte, size+1)

	for i := 0; i < len1; i++ {
		res[i+size+1-len1] = num1[i] - '0'
	}

	for i := 0; i < len2; i++ {
		res[i+size+1-len2] += num2[i] - '0'
	}

	for i := size; i >= 0; i-- {
		if res[i] >= 10 {
			res[i-1] += 1
			res[i] -= 10
		}
		res[i] += '0'
	}

	if res[0] == '0' {
		return string(res[1:])
	}

	return string(res)
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(addStrings("11", "123"))
}
