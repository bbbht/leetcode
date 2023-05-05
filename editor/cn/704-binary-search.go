package main

// 704 binary-search 2023-04-20 16:06:55

// leetcode submit region begin(Prohibit modification and deletion)
func search(nums []int, target int) int {
	left, right, mid := 0, len(nums)-1, 0
	for left <= right {
		mid = left + (right-left)>>1

		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			return mid
		}
	}

	return -1
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
