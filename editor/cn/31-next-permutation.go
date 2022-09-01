package main

// 下一个排列 2022-07-27 15:53:28

// leetcode submit region begin(Prohibit modification and deletion)
func nextPermutation(nums []int) {
	size := len(nums)
	left, right := -1, size-1
	// 找到第一个不符合递减规律的位置，i-1
	for i := size - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			left = i - 1
			break
		}
	}
	if left == -1 {
		reverse31(0, size-1, nums)
		return
	}

	// 找出 left 之后（递减），最接近且大于它的数字
	for right = size - 1; right > left; {
		if nums[left] < nums[right] {
			break
		}
		right--
	}

	nums[left], nums[right] = nums[right], nums[left]

	reverse31(left+1, size-1, nums)
}

func reverse31(left, right int, nums []int) {
	for i, j := left, right; i < j; {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

// func main() {
// 	l := []int{5, 4, 3, 2, 1}
// 	nextPermutation(l)
// 	fmt.Println(l)
// }

// leetcode submit region end(Prohibit modification and deletion)
