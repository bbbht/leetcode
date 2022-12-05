package main

// 258 add-digits 2022-12-01 15:33:07

// leetcode submit region begin(Prohibit modification and deletion)
func addDigits(num int) int {
	if num <= 0 {
		return num
	}
	// 数字相加，每9个数一个循环，演算下
	return (num-1)%9 +1

	// for {
	// 	sum := 0
	// 	for num > 0 {
	// 		sum += num % 10
	// 		num /= 10
	// 	}
	// 	if sum <= 9 {
	// 		return sum
	// 	}
	// 	num = sum
	// }
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
