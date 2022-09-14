package main

// 45 jump-game-ii 2022-08-30 17:00:53

// leetcode submit region begin(Prohibit modification and deletion)
func jump(nums []int) int {
	// index 逐步移动
	// cur 保存当前到达的最远位置
	n, cur, size, index := 0, 0, len(nums), 0
	for cur < size-1 {
		n++
		// 记录当前的位置，让索引从上次的位置走到当前位置来
		pre := cur
		for ; index <= pre; index++ {
			// 范围内可以跳到更远的位置，则记录下来
			if nums[index]+index > cur {
				cur = nums[index] + index
			}
		}
		// 没移动，那就不会再移动了，没有解
		if pre == cur {
			return -1
		}
	}

	return n
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
