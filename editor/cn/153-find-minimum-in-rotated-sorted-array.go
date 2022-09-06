package main

// 153 find-minimum-in-rotated-sorted-array 2022-09-01 11:09:14

// leetcode submit region begin(Prohibit modification and deletion)
func findMin(nums []int) int {
	size := len(nums)
	if nums[0] < nums[size-1] {
		return nums[0]
	}
	mid, left, right := 0, 0, size-1
	for left < right-1 {
		mid = (left + right) / 2
		if nums[mid] > nums[0] {
			left = mid
		} else {
			right = mid
		}
	}

	return nums[right]
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
