package main

import (
	"sort"
)

// 242 valid-anagram 2022-12-01 15:22:49

// leetcode submit region begin(Prohibit modification and deletion)
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	bs, bt := []byte(s), []byte(t)
	sort.Slice(bs, func(i, j int) bool {
		return bs[i] < bs[j]
	})
	sort.Slice(bt, func(i, j int) bool {
		return bt[i] < bt[j]
	})
	for i := range bs {
		if bs[i] != bt[i] {
			return false
		}
	}

	return true
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
