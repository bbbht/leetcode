package main

// 75 sort-colors 2022-10-18 17:12:34

// leetcode submit region begin(Prohibit modification and deletion)
func sortColors(nums []int) {
	left, right := 0, len(nums)-1
	for i := left; i <= right; i++ {
		if nums[i] == 0 {
			// 交换，然后左侧向右移动，因为它当前及之前的位置已经是 0
			// 且被交换到i的值肯定是0或1，2已被交换到右侧（i>=left）
			nums[i], nums[left] = nums[left], nums[i]
			left++
		}
		if nums[i] == 2 {
			// 交换，然后右侧侧向左移动,i保持不动，因为右侧交换过来的可能为 0/1/2
			nums[i], nums[right] = nums[right], nums[i]
			right--
			i--
		}
	}
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
