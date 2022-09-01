package main

// 电话号码的字母组合 2022-07-22 11:49:58

// leetcode submit region begin(Prohibit modification and deletion)

// func main() {
// 	fmt.Println(letterCombinations("79"))
// }

var m17 = map[uint8][]uint8{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}

	size := 1
	for i := range digits {
		size *= len(m17[digits[i]])
	}

	tmp := make([][]uint8, size)
	res := make([]string, size)
	for i := range tmp {
		tmp[i] = make([]uint8, len(digits))
	}

	// 一行一行进行填充，规则为循环平铺
	cycle := size
	for i := range digits {
		// 应输出字母
		cycle = cycle / len(m17[digits[i]])
		for j := range tmp {
			index := (j / cycle) % len(m17[digits[i]])
			tmp[j][i] = m17[digits[i]][index]
		}
	}

	for i := range res {
		res[i] = string(tmp[i])
	}

	return res
}

// leetcode submit region end(Prohibit modification and deletion)
