package main

import (
	"math"
	"sort"
)

// 最接近的三数之和 2022-07-22 11:26:08

// leetcode submit region begin(Prohibit modification and deletion)
// func main() {
// 	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4, 1}, 1))
// }
func threeSumClosest(nums []int, target int) int {
	res, tmp, size := math.MaxInt32, 0, len(nums)
	sort.Ints(nums)
	for i := range nums {
		l, r := i+1, size-1
		for l < r {
			tmp = nums[i] + nums[l] + nums[r]

			if tmp == target { // 相等是最接近
				return tmp
			}

			if abs(res, target) > abs(tmp, target) {
				res = tmp
			}

			if tmp > target {
				r--
			}
			if tmp < target {
				l++
			}
		}
	}

	return res
}

func abs(a, b int) int {
	if a-b > 0 {
		return a - b
	}
	return b - a
}

// leetcode submit region end(Prohibit modification and deletion)
