package main

// 罗马数字转整数 2022-07-18 18:23:07

// leetcode submit region begin(Prohibit modification and deletion)
var m13 = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	res, last := 0, 0

	for i := len(s) - 1; i >= 0; i-- {
		temp := m13[s[i]]

		sign := 1
		if temp < last {
			sign = -1
		}

		res += sign * temp

		last = temp
	}

	return res
}

// leetcode submit region end(Prohibit modification and deletion)
