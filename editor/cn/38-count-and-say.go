package main

import (
	"bytes"
	"strconv"
)

// 外观数列 2022-07-28 16:40:57

// leetcode submit region begin(Prohibit modification and deletion)
func countAndSay(n int) string {
	s := "1"
	for i := 2; i <= n; i++ {
		tag, num, buffer := s[0], 0, bytes.Buffer{}
		for j := 0; j <= len(s); j++ {
			if j == len(s) || s[j] != tag {
				buffer.WriteString(strconv.Itoa(num))
				buffer.WriteString(string(tag))
				if j == len(s) {
					break
				}
				tag = s[j]
				num = 1
			} else {
				num++
			}
		}
		s = buffer.String()
	}
	return s
}

// leetcode submit region end(Prohibit modification and deletion)
