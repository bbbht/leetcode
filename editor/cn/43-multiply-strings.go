package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 43 multiply-strings 2022-08-01 11:46:26

// leetcode submit region begin(Prohibit modification and deletion)
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	len1, len2 := len(num1), len(num2)

	result := make([]uint8, len1+len2)

	for i := len2 - 1; i >= 0; i-- {
		for j := len1 - 1; j >= 0; j-- {
			v := (num2[i] - '0') * (num1[j] - '0')
			idx := i + j + 1
			result[idx] += v
			// fmt.Println(i, j, idx, v)

			for ; result[idx] > 9; idx-- {
				result[idx-1] += result[idx] / 10
				result[idx] %= 10
			}
		}
	}

	res := ""
	for i := range result {
		res += strconv.Itoa(int(result[i]))
	}

	return strings.TrimLeft(res, "0")
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	n1, n2 := 990, 999999
	fmt.Println(multiply(strconv.Itoa(n1), strconv.Itoa(n2)), n1*n2)
}
