package main
// 231 power-of-two 2022-11-30 15:56:43

//leetcode submit region begin(Prohibit modification and deletion)
func isPowerOfTwo(n int) bool {
	return n > 0 && n & (n-1) == 0
}
//leetcode submit region end(Prohibit modification and deletion)


func main() {

}