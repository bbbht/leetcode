package main

// 搜索旋转排序数组 2022-07-28 11:08:50

// leetcode submit region begin(Prohibit modification and deletion)
func search(nums []int, target int) int {
	size := len(nums)
	if size == 0 {
		return -1
	}
	if target == nums[0] {
		return 0
	}

	left, right, mid := 0, size-1, 0
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[left] <= nums[mid] { // mid 未超过翻转点
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
				continue
			}
			left = mid + 1
		} else { // mid 超过了翻转点，较小的递增区间
			if nums[right] >= target && target > nums[mid] {
				left = mid + 1
				continue
			}
			right = mid - 1
		}
	}

	return -1
}

// func main() {
// 	fmt.Println(search([]int{6, 7, 8, 9, 10, 11, 12, 13, 0, 1, 2, 3, 4, 5}, 4))
// }

// leetcode submit region end(Prohibit modification and deletion)
