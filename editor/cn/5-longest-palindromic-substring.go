package main

import (
	"fmt"
	"strings"
)

// 最长回文子串 2022-07-12 18:18:43

// leetcode submit region begin(Prohibit modification and deletion)
// 超时
// func longestPalindrome(s string) string {
// 	size := len(s)
// 	if size <= 1 {
// 		return s
// 	}
//
// 	m := make(map[int]map[int]bool, size)
// 	for i := 0; i < size; i++ {
// 		if _, ok := m[i]; !ok {
// 			m[i] = make(map[int]bool, size-i)
// 		}
// 	}
// 	max, res := 0, s[0:1]
//
// 	for r := 1; r < size; r++ {
// 		for l := 0; l < r; l++ {
// 			m[l][r] = s[l] == s[r] && (r-l <= 2 || m[l+1][r-1])
// 			if m[l][r] {
// 				if max < r - l + 1 {
// 					max = r-l+1
// 					res = s[l:r+1]
// 				}
// 				if max == size {
// 					return s
// 				}
// 			}
// 		}
// 	}
//
// 	return res
// }

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	return manacher(s)
}

// 这个算法其实就是中心扩散法的改进
// 原始算法需要奇偶，通过插入间隔符使得新字符串肯定只包含奇数的回文串，从而扩散的次数就是字符串长度
// 然后对新字符串使用中心扩散法即可
func manacher(s string) string {
	mstr := "#" + strings.Join(strings.Split(s, ""), "#") + "#"
	length, max, start := len(mstr), 0, 0
	for i := 0; i < length; i++ {
		left, right, step := i-1, i+1, 0

		for {
			if !(left >= 0 && right < length && mstr[left] == mstr[right]) {
				break
			}
			left--
			right++
			step++
		}

		if step > max {
			max = step
			start = (i - max) / 2
		}
	}
	return s[start : start+max]
}
// -a-b-a-d-a-d-d-a-
func main()  {
	fmt.Println(longestPalindrome("a"))
}
// leetcode submit region end(Prohibit modification and deletion)
