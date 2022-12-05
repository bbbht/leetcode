package main

import (
	"strings"
)

// 290 word-pattern 2022-12-01 18:32:36

// leetcode submit region begin(Prohibit modification and deletion)
func wordPattern(pattern string, s string) bool {
	sm := make(map[string]byte)
	pm := make(map[byte]string)

	sArr := strings.Fields(s)

	if len(pattern) != len(sArr) {
		return false
	}

	for i := range sArr {
		_, sOk := sm[sArr[i]]
		_, pOk := pm[pattern[i]]
		if !sOk && !pOk {
			sm[sArr[i]] = pattern[i]
			pm[pattern[i]] = sArr[i]
		} else if pOk && sOk {
			if sm[sArr[i]] != pattern[i] || pm[pattern[i]] != sArr[i] {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
