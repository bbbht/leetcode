package main

import (
	"fmt"
	"strconv"
)

// 228 summary-ranges 2022-11-30 11:41:18

// leetcode submit region begin(Prohibit modification and deletion)
func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}

	low, high, size, res := 0, 0, len(nums), make([]string, 0)

	for high < size {
		for high < size-1 && nums[high]+1 == nums[high+1] {
			high++
		}
		if low == high {
			res = append(res, strconv.Itoa(nums[low]))
		} else if low < high {
			res = append(res, strconv.Itoa(nums[low]) + "->" + strconv.Itoa(nums[high]))
		}

		high++
		low = high
	}

	return res
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(summaryRanges([]int{1}))
	fmt.Println(summaryRanges([]int{1, 2, 3}))
	fmt.Println(summaryRanges([]int{1, 2, 4, 6}))
}
