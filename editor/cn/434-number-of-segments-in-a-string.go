package main

import (
	"strings"
)

// 434 number-of-segments-in-a-string 2022-12-08 16:54:05

// leetcode submit region begin(Prohibit modification and deletion)
func countSegments(s string) int {
	s = strings.Trim(s, " ")
	n, size := 0, len(s)
	if size == 0 {
		return 0
	}
	for i := range s {
		if s[i] == ' ' && i < size-1 && s[i+1] != ' ' {
			n++
		}
	}

	return n + 1
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
