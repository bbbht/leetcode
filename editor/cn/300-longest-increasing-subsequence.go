package main

// 300 longest-increasing-subsequence 2022-09-06 13:35:25

// leetcode submit region begin(Prohibit modification and deletion)
// func lengthOfLIS(nums []int) int {
// 	res, dp := 1, make([]int, len(nums))
// 	for i := range nums {
// 		dp[i] = 1
// 		for j := 0; j < i; j++ {
// 			if nums[i] <= nums[j] {
// 				continue
// 			}
//			// 如果 i 处值大于 j 处，那么 i处值 一定能加入 j 处值形成子序列
// 			dp[i] = max300(dp[i], dp[j]+1)
// 			res = max300(dp[i], res)
// 		}
// 	}
//
// 	return res
// }
//
// func max300(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

func lengthOfLIS(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}

	res := []int{nums[0]}
	for i := range nums {
		// 如果当前值大于队列末尾的值，则追加
		if nums[i] > res[len(res)-1] {
			res = append(res, nums[i])
			continue
		}
		// 如果当前值小于队列末尾的值，则通过二分查找方法找到第一个大于等于该值的位置进行替换
		// 参考了一篇找不到的帖子思路，初始为动态规划
		// 此方法下，临时队列重要的不算绝对正确的子序列顺序，而是最终队列的长度，可以找例子试一下
		if nums[i] < res[len(res)-1] {
			left, right, mid := 0, len(res)-1, 0
			for left <= right {
				mid = (left + right) / 2
				if res[mid] >= nums[i] {
					right = mid - 1
				} else {
					left = mid + 1
				}
			}

			// 替换后并不影响队列的递增
			res[left] = nums[i]
		}
	}
	return len(res)
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
