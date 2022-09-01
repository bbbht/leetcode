package main

// 无重复字符的最长子串 2022-07-06 16:55:20

// leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	l, ml := len(s), 0
	m := make(map[uint8]int, 52)

	for i, j := 0, 0; j < l; j++ {
		// 只使用 ASCII 码即可
		v := s[j]
		// k >= i 如果是当前窗口之外的字符，不需要考虑
		if k, ok := m[v]; ok && k >= i {
			// 字符已存在，则直接跳过它后面
			i = k + 1
		}

		// 无论是否存在，存入当前字符及位置
		m[v] = j

		// 最大长度判断
		ml = max3(ml, j+1-i)

		// 剩下的就没必要判断了
		if l-i <= ml {
			break
		}
	}

	return ml
}

func max3(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// leetcode submit region end(Prohibit modification and deletion)
