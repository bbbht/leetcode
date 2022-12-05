package main

// 136 single-number 2022-11-28 14:50:41

// leetcode submit region begin(Prohibit modification and deletion)
func singleNumber(nums []int) int {
	res := 0
	// 异或，两次出现的异或结果为 0
	for _, v := range nums {
		res = res ^ v
	}

	return res
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
