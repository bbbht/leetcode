package main

// 387 first-unique-character-in-a-string 2022-12-07 17:01:48

// leetcode submit region begin(Prohibit modification and deletion)
func firstUniqChar(s string) int {
	m := make(map[byte]int, len(s))
	for i := range s {
		if _, ok := m[s[i]]; ok {
			m[s[i]] += 1
		} else {
			m[s[i]] = 1
		}
	}

	for i := range s {
		if n, _ := m[s[i]]; n == 1 {
			return i
		}
	}

	return -1
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
