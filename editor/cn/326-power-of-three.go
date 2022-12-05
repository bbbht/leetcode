package main

// 326 power-of-three 2022-12-02 15:27:01

// leetcode submit region begin(Prohibit modification and deletion)
func isPowerOfThree(n int) bool {
	// 1162261467 is 3^19
	return n > 0 && 1162261467%n == 0
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
