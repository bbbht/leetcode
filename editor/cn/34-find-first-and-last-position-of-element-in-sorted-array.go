package main

// 在排序数组中查找元素的第一个和最后一个位置 2022-07-28 13:46:29

// leetcode submit region begin(Prohibit modification and deletion)
func searchRange(nums []int, target int) []int {
	res, size := []int{-1, -1}, len(nums)

	if size == 0 {
		return res
	}
	res[0] = binarySearch34Left(nums, target)
	if res[0] == -1 {
		return res
	}
	res[1] = binarySearch34Right(nums, target)

	return res
}

func binarySearch34Left(nums []int, target int) int {
	border := -1
	for l, r := 0, len(nums)-1; l <= r; {
		mid := (l + r) / 2

		if nums[mid] >= target {
			border = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	if border >= 0 && nums[border] == target {
		return border
	}

	return -1
}
func binarySearch34Right(nums []int, target int) int {
	border := -1
	for l, r := 0, len(nums)-1; l <= r; {
		mid := (l + r) / 2

		if nums[mid] > target {
			r = mid - 1
		} else {
			border = mid
			l = mid + 1
		}
	}

	if border >= 0 && nums[border] == target {
		return border
	}

	return -1
}

// func main() {
// 	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
// 	fmt.Println(searchRange([]int{1, 2, 3, 4, 5, 7, 7, 8, 8, 8, 8}, 8))
// 	fmt.Println(searchRange([]int{1}, 8))
// 	fmt.Println(searchRange([]int{1}, 1))
// }

// leetcode submit region end(Prohibit modification and deletion)
