package main

// 76 minimum-window-substring 2023-05-04 19:27:21

// leetcode submit region begin(Prohibit modification and deletion)
func minWindow(s string, t string) string {
	tCount := [128]int{}
	left, cnt, minLeft, minSize := 0, 0, -1, len(s)+1

	for i := range t {
		tCount[t[i]]++
	}

	for i := range s {
		tCount[s[i]]--
		if tCount[s[i]] >= 0 { // 只有t中等待字符会出现大于0的结果，且达到次数后不再累加
			cnt++ // 统计t中字符出现的次数
		}

		for cnt == len(t) { // 当窗口中包含了所有的t字符，就开始收缩左侧边界
			if minSize > i-left+1 {
				minSize = i - left + 1
				minLeft = left
			}

			tCount[s[left]]++

			// 如果最左侧的字符为t中字符，则其值肯定为0
			// 值加一后，循环直到再次遇见该字符，cnt才会等于t长度，再尝试移动左边界
			// 而非t中字符，由于前面的减一操作，最多也只会累加到0
			if tCount[s[left]] > 0 {
				cnt--
			}

			left++
		}
	}

	if minLeft == -1 {
		return ""
	}

	return s[minLeft : minLeft+minSize]

}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	minWindow("ADOBECODEBANC", "ABC")
}
