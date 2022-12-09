package main

// 409 longest-palindrome 2022-12-08 15:23:34

// leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) int {
	mid, res, m := 0, 0, make(map[byte]int)
	for i := range s {
		if _, ok := m[s[i]]; !ok {
			m[s[i]] = 0
		}
		m[s[i]] += 1
	}

	for i := range m {
		res += m[i] & 0xfffffffe
		if mid == 0 && m[i]&1 == 1 {
			mid = 1
		}
	}

	return res + mid
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
