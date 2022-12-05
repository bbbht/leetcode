package main
// 191 number-of-1-bits 2022-11-29 11:40:04

//leetcode submit region begin(Prohibit modification and deletion)
func hammingWeight(num uint32) int {
	n := 0
	// n & (n-1) 将最低位的 1 置零
	for ; num > 0; num &= num - 1 {
		n++
	}

	return n
}
//leetcode submit region end(Prohibit modification and deletion)


func main() {

}