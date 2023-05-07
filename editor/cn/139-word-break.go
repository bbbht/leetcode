package main

import (
	"fmt"
)

// 139 word-break 2023-05-06 20:46:26

// leetcode submit region begin(Prohibit modification and deletion)
func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true

	dict := make(map[string]bool, len(wordDict))
	for i := range wordDict {
		dict[wordDict[i]] = true
	}

	for i := 0; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && dict[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}
