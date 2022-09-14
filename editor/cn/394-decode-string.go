package main

import (
	"strings"
)

// 394 decode-string 2022-09-08 17:15:06

// leetcode submit region begin(Prohibit modification and deletion)
func decodeString(s string) string {
	var strStack []byte // 字符栈
	var numStack []int  // 数字栈
	num := 0
	for i := range s {
		// 数字范围是 1-300，即最多三位
		if s[i] <= '9' && s[i] >= '0' {
			num = num*10 + int(s[i]-'0')
			continue
		}

		// 数字阶段结束，压入完整数字，并重置
		if num > 0 {
			numStack = append(numStack, num)
			num = 0
		}

		if (s[i] <= 'z' && s[i] >= 'a') || s[i] == '[' {
			strStack = append(strStack, s[i])
			continue
		}

		// 剩下的只有一种情况了，]，意味着要开始出栈以拼接字符串了
		var charArr []byte
		// 字节拼接，由于出入栈关系，需要翻转
		for strStack[len(strStack)-1] != '[' {
			charArr = append(charArr, strStack[len(strStack)-1])
			strStack = strStack[:len(strStack)-1]
		}
		for j := 0; j < len(charArr)/2; j++ {
			charArr[j], charArr[len(charArr)-j-1] = charArr[len(charArr)-j-1], charArr[j]
		}
		str := string(charArr)

		// 移除 [
		strStack = strStack[:len(strStack)-1]

		// 获取重复次数，由于 [ 前肯定是数字，所以肯定有对应的数字可用
		str = strings.Repeat(str, numStack[len(numStack)-1])

		// 弹出数字栈数字
		numStack = numStack[:len(numStack)-1]

		// 解码操作后，将阶段性的结果入栈
		for j := range str {
			strStack = append(strStack, str[j])
		}
	}
	return string(strStack)
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
