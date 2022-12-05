package main

// 89 gray-code 2022-11-07 10:34:34

// leetcode submit region begin(Prohibit modification and deletion)
func grayCode(n int) []int {
	result := make([]int, 0, 1<<n)
	result = append(result, 0)
	// 从1开始向后推，相当于前面一半数字前面+0（不变），后面新增一倍数字前面+1（倒叙）
	// 可以自行推算
	for i := 0; i < n; i++ {
		size, add := len(result), 1<<i
		for j := size - 1; j >= 0; j-- {
			result = append(result, result[j]+add)
		}
	}

	return result
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
