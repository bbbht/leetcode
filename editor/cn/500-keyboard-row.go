package main

import (
	"strings"
)

// 500 keyboard-row 2022-12-09 10:49:37

// leetcode submit region begin(Prohibit modification and deletion)
var m500 = map[byte]uint8{'q': 1, 'w': 1, 'e': 1, 'r': 1, 't': 1, 'y': 1, 'u': 1, 'i': 1, 'o': 1, 'p': 1, 'a': 2, 's': 2, 'd': 2, 'f': 2, 'g': 2, 'h': 2, 'j': 2, 'k': 2, 'l': 2, 'z': 3, 'x': 3, 'c': 3, 'v': 3, 'b': 3, 'n': 3, 'm': 3}

func findWords(words []string) []string {
	res := make([]string, 0, len(words))
	for i := range words {
		ok, tmp := true, strings.ToLower(words[i])
		for j := range tmp {
			if m500[tmp[j]] != m500[tmp[0]] {
				ok = false
				break
			}
		}
		if ok {
			res = append(res, words[i])
		}

	}

	return res
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
