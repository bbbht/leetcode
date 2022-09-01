package main

// 缺失的第一个正数 2022-07-29 16:10:14

// leetcode submit region begin(Prohibit modification and deletion)
func firstMissingPositive(nums []int) int {
	size := len(nums)
	for i := 0; i < size; i++ {
		// 排序，nums[i]应该是数字 i+1
		// 如果不是，且数值在索引范围，则将其替换到正确位置，不在索引范围，则可能缺失，向后继续遍历也许能将其归位
		for nums[i] != i+1 && nums[i] > 0 && nums[i] < size {
			if nums[i] == nums[nums[i]-1] {
				break
			}
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := range nums {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	// 全都在，那么就是缺失下一个
	return size + 1
}
//
// func main() {
// 	fmt.Println(firstMissingPositive([]int{1}))
// 	fmt.Println(firstMissingPositive([]int{1, 1}))
// 	fmt.Println(firstMissingPositive([]int{2, 2, 2}))
// }

// leetcode submit region end(Prohibit modification and deletion)
