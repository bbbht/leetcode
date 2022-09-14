package main

// 55 jump-game 2022-09-13 16:08:20

// leetcode submit region begin(Prohibit modification and deletion)
func canJump(nums []int) bool {
	cur, index, size := 0, 0, len(nums)

	for cur < size-1 {
		tmp := cur
		for ; index <= tmp; index++ {
			if index+nums[index] > cur {
				cur = index + nums[index]
			}
		}
		if tmp == cur {
			return false
		}
	}

	return true
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
