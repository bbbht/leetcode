package main

// 283 move-zeroes 2022-12-01 16:46:32

// leetcode submit region begin(Prohibit modification and deletion)
func moveZeroes(nums []int) {
	size := len(nums)
	for i, j := 0, 0; i < size; i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
