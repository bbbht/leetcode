package main

// 串联所有单词的子串 2022-07-27 11:18:58

// leetcode submit region begin(Prohibit modification and deletion)
func findSubstring(s string, words []string) []int {
	var res []int
	num, wlen := len(words), len(words[0])

	m := make(map[string]int, num)
	for i := range words {
		if _, ok := m[words[i]]; ok {
			m[words[i]] += 1
		} else {
			m[words[i]] = 1
		}
	}

	count, tmpM := 0, make(map[string]int, num)

	for i := 0; i <= len(s)-num*wlen; i++ {
		// 需要重置map的两种情况
		if count > 0 || i == 0 {
			for k := range m {
				tmpM[k] = m[k]
			}
		}

		count = 0

		for j := 0; j < num; j++ {
			start, end := i+j*wlen, i+j*wlen+wlen
			if v, ok := tmpM[s[start:end]]; !ok || v <= 0 {
				break
			}

			tmpM[s[start:end]] -= 1
			count++
		}

		// 匹配到的数量与words相符
		if count == num {
			res = append(res, i)
		}
	}

	return res
}

// func main() {
// 	fmt.Println(findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
// 	fmt.Println(findSubstring("wordgoodgoodgoodbestwordword", []string{"word", "good", "best", "word"}))
// 	fmt.Println(findSubstring("barfoofoobarthefoobarman", []string{"bar", "foo", "the"}))
// }

// leetcode submit region end(Prohibit modification and deletion)
