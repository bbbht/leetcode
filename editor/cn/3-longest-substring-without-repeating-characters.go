package main

// 无重复字符的最长子串 2022-07-06 16:55:20

// leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	size, ml := len(s), 0
	m := make(map[uint8]int, 52)

	for left, right := 0, 0; right < size; right++ {
		// 只使用 ASCII 码即可
		val := s[right]
		// index >= left 如果是当前窗口之外的字符，不需要考虑
		if index, ok := m[val]; ok && index >= left {
			// 字符已存在，则直接跳过它后面
			left = index + 1
		}

		// 无论是否存在，存入当前字符及位置
		m[val] = right

		// 最大长度判断
		ml = max3(ml, right+1-left)

		// 剩下的就没必要判断了
		if size-left <= ml {
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
