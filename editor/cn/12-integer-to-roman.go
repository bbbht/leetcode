package main

// 整数转罗马数字 2022-07-18 17:42:24

// leetcode submit region begin(Prohibit modification and deletion)
var (
	romanM = []string{"", "M", "MM", "MMM"}
	romanC = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	romanX = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	romanI = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
)

func intToRoman(num int) string {
	if num < 0 {
		return ""
	}

	// 题目标注，1-3999，所以直接排列组合即可
	return romanM[num/1000] + romanC[(num%1000)/100] + romanX[(num%100)/10] + romanI[num%10]
}

// leetcode submit region end(Prohibit modification and deletion)
