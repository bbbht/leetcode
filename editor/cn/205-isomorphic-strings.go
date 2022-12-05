package main

import (
	"fmt"
)

// 205 isomorphic-strings 2022-11-29 16:57:55

// leetcode submit region begin(Prohibit modification and deletion)
func isIsomorphic(s string, t string) bool {
	m1 := make(map[byte]int)
	m2 := make(map[byte]int)
	for i := range s {
		m1[s[i]] = i
		m2[t[i]] = i
	}
	for i := range s {
		if m1[s[i]] != m2[t[i]] {
			return false
		}
	}

	return true
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(isIsomorphic("badc", "baba"))
}
