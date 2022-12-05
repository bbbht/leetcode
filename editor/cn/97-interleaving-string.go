package main

// 97 interleaving-string 2022-11-21 15:54:54

// leetcode submit region begin(Prohibit modification and deletion)
func isInterleave(s1 string, s2 string, s3 string) bool {
	l1, l2 := len(s1), len(s2)

	if l1+l2 != len(s3) {
		return false
	}

	dp := make([]bool, len(s2)+1)
	for i := 0; i <= l1; i++ {
		for j := 0; j <= l2; j++ {
			if i == 0 && j == 0 {
				dp[j] = true
			} else if i == 0 {
				dp[j] = dp[j-1] && s2[j-1] == s3[j-1]
			} else if j == 0 {
				dp[j] = dp[j] && s1[i-1] == s3[i-1]
			} else {
				dp[j] = (dp[j] && s1[i-1] == s3[i+j-1]) ||
					(dp[j-1] && s2[j-1] == s3[i+j-1])
			}
		}
	}

	return dp[l2]
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
