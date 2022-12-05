package main

// 342 power-of-four 2022-12-02 16:48:15

// leetcode submit region begin(Prohibit modification and deletion)
func isPowerOfFour(n int) bool {
	return n > 0 && (n&(n-1) == 0 && (n-1)%3 == 0)
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
