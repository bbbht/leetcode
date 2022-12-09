package main

// 392 is-subsequence 2022-12-07 17:32:05

// leetcode submit region begin(Prohibit modification and deletion)
func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}

	j, size := 0, len(t)
	for i := range s {
		for ; j <= size; j++ {
			if j == size {
				return false
			}
			if t[j] == s[i] {
				j++
				break
			}
		}
	}

	return true
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	// j := 0
	// for ; j < 3; j++ {
	// 	if j == 1 {
	// 		break
	// 	}
	// }
	// fmt.Println(j)
}
