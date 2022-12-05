package main

// 344 reverse-string 2022-12-02 17:27:22

// leetcode submit region begin(Prohibit modification and deletion)
func reverseString(s []byte) {
	l, r := 0, len(s)-1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
