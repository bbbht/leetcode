package main

// 28 find-the-index-of-the-first-occurrence-in-a-string 2023-05-07 17:43:02

// leetcode submit region begin(Prohibit modification and deletion)
func strStr(haystack string, needle string) int {
	for i := 0; ; i++ {
		for j := 0; ; j++ {
			if j == len(needle) {
				return i
			}

			if i+j == len(haystack) {
				return -1
			}

			if needle[j] != haystack[i+j] {
				break
			}
		}
	}
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
