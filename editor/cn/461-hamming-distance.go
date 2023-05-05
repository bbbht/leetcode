package main

// 461 hamming-distance 2022-12-12 13:52:09

// leetcode submit region begin(Prohibit modification and deletion)
func hammingDistance(x int, y int) int {
	n := 0
	num := x ^ y
	for ; num > 0; num &= num - 1 {
		n++
	}
	return n
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
