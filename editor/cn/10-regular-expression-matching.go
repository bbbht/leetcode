package main

import (
	"fmt"
)

// 正则表达式匹配 2022-07-18 17:16:12

// leetcode submit region begin(Prohibit modification and deletion)
func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	if len(p) == 1 {
		return len(s) == 1 && (s[0] == p[0] || p[0] == '.')
	}

	// 如果非 * 则第一个字符必定要匹配
	if p[1] != '*' {
		if len(s) == 0 {
			return false
		}
		return (s[0] == p[0] || p[0] == '.') && isMatch(s[1:], p[1:])
	}

	// 当 * 匹配为0的情况，即p的前2个字符不匹配 s 的任何值而是由后续的规则匹配
	for len(s) > 0 && (s[0] == p[0] || p[0] == '.') {
		if isMatch(s, p[2:]) {
			return true
		}

		// 去掉匹配的 s 首字母，因为* 可以匹配任意个，然循环直到 p 的前两个规则无法匹配，跳出循环
		s = s[1:]
	}

	return isMatch(s, p[2:])
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(isMatch("a", ".*..a*"))
}
