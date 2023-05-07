package main

import (
	"fmt"
)

// 移除元素 2022-07-25 17:12:21

// leetcode submit region begin(Prohibit modification and deletion)
func removeElement(nums []int, val int) int {
	num := 0
	// 直接将不需要删除的值前移即可
	for _, item := range nums {
		if item != val {
			nums[num] = item
			num++
		}
	}

	return num
}

// func removeElement(nums []int, val int) int {
// 	cnt := 0
// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] != val {
// 			if i != cnt {
// 				nums[i], nums[cnt] = nums[cnt], nums[i]
// 			}
// 			cnt++
// 		}
// 	}
//
// 	return cnt
// }

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}

	fmt.Println(removeElement(nums, 2), nums)
}
