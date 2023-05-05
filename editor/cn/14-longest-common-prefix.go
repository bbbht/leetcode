package main

// 最长公共前缀 2022-07-18 18:33:58

// leetcode submit region begin(Prohibit modification and deletion)
func longestCommonPrefix(strs []string) string {
	res := ""
	if len(strs) == 0 {
		return res
	}

	maxLen := len(strs[0])
	for i := range strs {
		if len(strs[i]) < maxLen {
			maxLen = len(strs[i])
		}
	}

	// r为上限，left 跳增，right 尽量递减
	mid, left, right := 0, 0, maxLen
	for right > 0 && left < right {
		mid = (left + right + 1) / 2
		if checkLongest(mid, strs) {
			left = mid
		} else {
			right = mid - 1
		}
	}

	return strs[0][0:left]
}

func checkLongest(lp int, strs []string) bool {
	for i := range strs {
		if strs[i][0:lp] != strs[0][0:lp] {
			return false
		}

	}
	return true
}

// func main()  {
// 	fmt.Println(longestCommonPrefix([]string{"floflower", "floflow", "flofloight"}))
// }
// leetcode submit region end(Prohibit modification and deletion)
