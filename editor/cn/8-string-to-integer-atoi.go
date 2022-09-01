package main

import (
	"math"
	"strings"
)

// 字符串转换整数 (atoi) 2022-07-18 16:21:18

// leetcode submit region begin(Prohibit modification and deletion)
func myAtoi(s string) int {
	res, isNegative := 0, false
	s = strings.TrimLeft(s, " ")
	if len(s) == 0 {
		return res
	}

	if s[0] == '-' {
		isNegative = true
	}

	for i := range s {
		if '0' > s[i] || s[i] > '9' {
			if i == 0 && (s[i] == '+' || s[i] == '-') {
				continue
			}
			break
		}
		res = res*10 + int(s[i]-'0')
		if res > math.MaxInt32 {
			if isNegative {
				res = -math.MinInt32
			} else {
				res = math.MaxInt32
			}
			break
		}
	}

	if isNegative {
		return -res
	}
	return res
}

// func main() {
// 	fmt.Println(myAtoi("--12"))
// }

// leetcode submit region end(Prohibit modification and deletion)
