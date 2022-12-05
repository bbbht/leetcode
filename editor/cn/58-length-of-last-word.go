package main

import (
	"strings"
)

// 58 length-of-last-word 2022-09-14 13:59:46

//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLastWord(s string) int {
	s = strings.TrimRight(s, " ")
	index := len(s) - 1
	for ; index >= 0; index-- {
		if s[index] == ' ' {
			break
		}
	}
	return len(s) - index - 1
}
//leetcode submit region end(Prohibit modification and deletion)


func main() {

}