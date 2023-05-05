package main

// 55 jump-game 2022-09-13 16:08:20

// leetcode submit region begin(Prohibit modification and deletion)
// func canJump(nums []int) bool {
// 	cur, index, size := 0, 0, len(nums)
//
// 	for cur < size-1 {
// 		tmp := cur
// 		for ; index <= tmp; index++ {
// 			if index+nums[index] > cur {
// 				cur = index + nums[index]
// 			}
// 		}
// 		if tmp == cur {
// 			return false
// 		}
// 	}
//
// 	return true
// }

func canJump(nums []int) bool {
	// 每次计算最大跳跃位置，如果某处小于等于最大位置，则无法到达终点
	if len(nums) == 0 {
		return false
	}
	if len(nums) == 1 {
		return true
	}
	maxRight := nums[0]
	for i := 1; i < len(nums); i++ {
		if i > maxRight {
			return false
		}
		if i+nums[i] > maxRight {
			maxRight = i + nums[i]
		}
		if maxRight >= len(nums)-1 {
			return true
		}
	}

	return false
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
