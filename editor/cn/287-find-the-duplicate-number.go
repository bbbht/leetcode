package main

import (
	"fmt"
)

// 287 find-the-duplicate-number 2022-09-06 10:13:33

// leetcode submit region begin(Prohibit modification and deletion)
func findDuplicate(nums []int) int {
	size := (len(nums) + 7) / 8
	bits := make([]byte, size)

	for i := range nums {
		// 取整定位到对应的 byte 索引，取余定位到具体 位索引
		if bits[nums[i]/8]&(1<<(nums[i]%8)) != 0 {
			return nums[i]
		}
		bits[nums[i]/8] |= 1 << (nums[i] % 8)
		// fmt.Printf("%08b \n", bits)
	}

	return -1
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(findDuplicate([]int{1,3,4,2,5,6,7,8,9,9}))
}
