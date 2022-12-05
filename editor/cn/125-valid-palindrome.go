package main

import (
	"fmt"
	"strings"
)

// 125 valid-palindrome 2022-11-25 16:59:28

// leetcode submit region begin(Prohibit modification and deletion)
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	size := len(s)
	l, r := 0, size-1
	for l < r {
		for !isValid125(s[l]) && l < r  {
			l++
		}
		for !isValid125(s[r]) && r > l {
			r--
		}
		if l > r || s[l] != s[r] {
			return false
		}
		l++
		r--
	}

	return true
}

func isValid125(b byte) bool {
	return b >= '0' && b <= '9' || b >= 'a' && b <= 'z'
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	// fmt.Println('a', 'z', 'A', 'Z', '0', '9')
	// fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome("  "))
}
