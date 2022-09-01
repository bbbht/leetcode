package main

// 括号生成 2022-07-22 18:27:30

// leetcode submit region begin(Prohibit modification and deletion)
func generateParenthesis(n int) []string {
	res := make([]string, 0)
	if n == 0 {
		return res
	}

	res = generate22("", n, n, res)

	return res
}

func generate22(s string, left, right int, res []string) []string {
	if left == 0 && right == 0 {
		res = append(res, s)
		return res
	}

	if left > 0 {
		res = generate22(s+"(", left-1, right, res)
	}

	if right > 0 && left < right {
		res = generate22(s+")", left, right-1, res)
	}

	return res
}
//
// func main()  {
// 	fmt.Println(generateParenthesis(3))
// }

// leetcode submit region end(Prohibit modification and deletion)
