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

	// r为上限，l 跳增，r 尽量递减
	mid, l, r := 0, 0, maxLen
	for r > 0 && l < r {
		mid = (l + r + 1) / 2
		if checkLongest(mid, strs) {
			l = mid
		} else {
			r = mid-1
		}
	}

	return strs[0][0:l]
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
