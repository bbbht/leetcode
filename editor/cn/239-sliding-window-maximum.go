package main

import (
	"fmt"
)

// 239 sliding-window-maximum 2023-05-06 18:03:50

// leetcode submit region begin(Prohibit modification and deletion)
func maxSlidingWindow(nums []int, k int) []int {
	queue := make([]int, 0, k)
	res := make([]int, 0, len(nums)-k+1)

	// 通过单调递减的双端队列来实现，原理是如果一个元素a比它前面的值b小，那么只要b元素还在队列中，那么a就不可能是最大值
	for i := range nums {
		// 如果当前队头的下标已经超出当前窗口的范围，移除掉它
		if len(queue) > 0 && i-queue[0] >= k {
			queue = queue[1:]
		}

		// 如果当前遍历到的下标对应的值大于队尾下标对应的值，那么就移除队尾，确保对头下标对应的值最大，总体是一个单调递减队列
		for len(queue) > 0 && nums[queue[len(queue)-1]] <= nums[i] {
			queue = queue[:len(queue)-1]
		}

		queue = append(queue, i)

		// 达到窗口长度后才开始输出窗口最大值结果
		if i+1 >= k {
			res = append(res, nums[queue[0]])
		}
	}

	return res
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(maxSlidingWindow([]int{4, 3, -1, -3, 5, 3, 6, 7}, 3))
}
