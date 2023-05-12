package main

// 238 product-of-array-except-self 2023-05-08 19:19:14

// leetcode submit region begin(Prohibit modification and deletion)
func productExceptSelf(nums []int) []int {
	size := len(nums)

	fnums, bnums, res := make([]int, size), make([]int, size), make([]int, size)
	fnums[0], bnums[size-1] = 1, 1
	for i := 0; i < size-1; i++ {
		fnums[i+1] = nums[i] * fnums[i]
	}
	for i := size - 1; i > 0; i-- {
		bnums[i-1] = nums[i] * bnums[i]
	}

	// 不能用乘法，
	// 分别构建两个数组，一个从前往后累乘，一个从后往前，则两个组合就可以求出任一位置的结果
	for i := 0; i < size; i++ {
		res[i] = fnums[i] * bnums[i]
	}

	return res
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
