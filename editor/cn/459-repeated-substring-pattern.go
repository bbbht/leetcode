package main

import (
	"fmt"
)

// 459 repeated-substring-pattern 2022-12-08 17:40:30

// leetcode submit region begin(Prohibit modification and deletion)
func repeatedSubstringPattern(s string) bool {
	size := len(s)
	for i := 1; i <= size/2; i++ {
		// 重复的开始
		for s[i] != s[0] {
			i++
			if i > size/2 {
				return false
			}
		}
		// 重复的部分药相等
		if isRepeat459(s, i) {
			return true
		}
	}

	return false
}

func isRepeat459(s string, subLen int) bool {
	if len(s)%subLen != 0 {
		return false
	}
	n := len(s) / subLen
	for j := 2; j <= n; j++ {
		start, end := subLen*(j-1), subLen*j
		if s[start:end] != s[:subLen] {
			return false
		}
	}

	return true
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(repeatedSubstringPattern("bb"))
}
