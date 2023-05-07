package main

// 189 rotate-array 2023-05-07 16:41:20

// leetcode submit region begin(Prohibit modification and deletion)
func rotate(nums []int, k int) {
	k = k % len(nums)

	// 翻转后，再份两部分反转回去，即可实现
	reverse189(nums)
	reverse189(nums[:k])
	reverse189(nums[k:])
}

func reverse189(nums []int) {
	size := len(nums)
	for i := 0; i < size/2; i++ {
		nums[i], nums[size-i-1] = nums[size-i-1], nums[i]
	}
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
