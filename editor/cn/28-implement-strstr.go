package main

// 实现 strStr() 2022-07-25 17:12:57

// leetcode submit region begin(Prohibit modification and deletion)
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	hLen, nLen := len(haystack), len(needle)
	for index := 0; index+nLen <= hLen; index++ {
		if needle == haystack[index:index+nLen] {
			return index
		}
	}

	return -1
}

// leetcode submit region end(Prohibit modification and deletion)
