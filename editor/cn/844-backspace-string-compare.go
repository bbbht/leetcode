package main

import (
	"fmt"
)

// 844 backspace-string-compare 2022-09-08 16:23:26

// leetcode submit region begin(Prohibit modification and deletion)
func backspaceCompare(s string, t string) bool {
	si, ti := len(s)-1, len(t)-1
	for {
		for back := 0; si >= 0 && (s[si] == '#' || back > 0); si-- {
			if s[si] == '#' {
				back++
			} else {
				back--
			}
		}
		for back := 0; ti >= 0 && (t[ti] == '#' || back > 0); ti-- {
			if t[ti] == '#' {
				back++
			} else {
				back--
			}
		}

		if si < 0 || ti < 0 {
			break
		}

		if s[si] != t[ti] {
			return false
		}

		si--
		ti--
	}

	return ti == si
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(backspaceCompare("xywrrmp", "xywrrmu#p"))
}
