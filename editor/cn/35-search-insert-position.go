package main
// 搜索插入位置 2022-07-28 15:50:46

//leetcode submit region begin(Prohibit modification and deletion)
func searchInsert(nums []int, target int) int {
	left, right, mid := 0, len(nums)-1, 0

	for left <= right {
		mid = (left + right) / 2
		if target == nums[mid] {
			return mid
		}
		if target < nums[mid] {
			if mid == 0 || target > nums[mid-1] {
				return mid
			}
			right = mid - 1
		} else {
			if mid == right || target < nums[mid+1] {
				return mid + 1
			}
			left = mid + 1
		}
	}

	return mid
}
//leetcode submit region end(Prohibit modification and deletion)
