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

// func letterCombinations(digits string) []string {
// 	if digits == "" {
// 		return nil
// 	}
//
// 	size := 1
// 	for i := range digits {
// 		size *= len(m17[digits[i]])
// 	}
//
// 	tmp := make([][]uint8, size)
// 	res := make([]string, size)
// 	for i := range tmp {
// 		tmp[i] = make([]uint8, len(digits))
// 	}
//
// 	// 一行一行进行填充，规则为循环平铺
// 	cycle := size
// 	for i := range digits {
// 		// 应输出字母
// 		cycle = cycle / len(m17[digits[i]])
// 		for j := range tmp {
// 			index := (j / cycle) % len(m17[digits[i]])
// 			tmp[j][i] = m17[digits[i]][index]
// 		}
// 	}
//
// 	for i := range res {
// 		res[i] = string(tmp[i])
// 	}
//
// 	return res
// }

func letterCombinations(digits string) []string {
	var result []string
	if digits == "" {
		return result
	}

	f17(digits, "", &result)

	return result
}

func f17(digits string, str string, result *[]string) {
	if digits == "" {
		*result = append(*result, str)
		return
	}

	num := digits[0]
	digits = digits[1:]
	for _, v := range m17[num] {
		str += string(v)
		f17(digits, str, result)
		str = str[:len(str)-1]
	}
}

// leetcode submit region end(Prohibit modification and deletion)
