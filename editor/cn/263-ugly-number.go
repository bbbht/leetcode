package main

// 263 ugly-number 2022-12-01 15:55:44

// leetcode submit region begin(Prohibit modification and deletion)
func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	for n%2 == 0 {
		n /= 2
	}
	for n%3 == 0 {
		n /= 3
	}
	for n%5 == 0 {
		n /= 5
	}
	return n == 1
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
