package main

import (
	"fmt"
)

// 91 decode-ways 2022-11-07 17:36:43

// leetcode submit region begin(Prohibit modification and deletion)
func numDecodings(s string) int {
	size := len(s)
	dp := make([]int, size+1)
	dp[0] = 1

	for i := 1; i <= size; i++ {
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}
		// 前面是两位数的情况
		if i > 1 && s[i-2] != '0' && (s[i-2]-'0')*10+(s[i-1]-'0') <= 26 {
			dp[i] += dp[i-2]
		}
	}

	return dp[size]
}

// func numDecodings(s string) int {
// 	m := make(map[string]int)
// 	return recursion91(s, m)
// }

// func recursion91(s string, m map[string]int) int {
// 	// 前面已经是最后一个划分可能了
// 	if s == "" {
// 		return 1
// 	}
// 	if s[0] == '0' {
// 		return 0
// 	}
//
// 	if n, ok := m[s]; ok {
// 		return n
// 	}
//
// 	// 个位数的可能
// 	r1, r2 := recursion91(s[1:], m), 0
// 	// 判断前两位是否符合条件
// 	if len(s) >= 2 && (s[0] == '1' || (s[0] == '2' && s[1] <= '6')) {
// 		r2 = recursion91(s[2:], m)
// 	}
//
// 	m[s] = r1 + r2
//
// 	return r1 + r2
// }

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(numDecodings("226"))
}
