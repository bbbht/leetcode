package main

// 最长有效括号 2022-07-27 18:11:12

// leetcode submit region begin(Prohibit modification and deletion)
func longestValidParentheses(s string) int {
	size, max := len(s), 0
	if size == 0 {
		return 0
	}

	for i := 0; i < size; {
		n := 0
		if max >= size-i {
			break
		}

		for j := i; j < size; j++ {
			switch s[j] {
			case '(':
				n += 1
			case ')':
				n -= 1
			}

			if n < 0 {
				i = j
				break
			}

			if j > i && n == 0 {
				max = max32(j-i+1, max)
			}
		}
		i++
	}

	return max
}

func max32(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func main() {
// 	fmt.Println(longestValidParentheses("()"))
// 	fmt.Println(longestValidParentheses(")()()())((()())()"))
// 	fmt.Println(longestValidParentheses("))))()()()()"))
// }

// leetcode submit region end(Prohibit modification and deletion)
