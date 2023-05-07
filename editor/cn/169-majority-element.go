package main

// 169 majority-element 2022-11-29 10:32:27

// leetcode submit region begin(Prohibit modification and deletion)
func majorityElement(nums []int) int {
	// 利用前提，如果某个数多余1半，那么不同于此数字时的抵消操作，最后必定剩下多数那个数字
	res, count := nums[0], 0
	for i := range nums {
		if count == 0 {
			res, count = nums[i], 1
		} else {
			if res == nums[i] {
				count++
			} else {
				count--
			}
		}
	}
	return res
}

// func majorityElement(nums []int) int {
// 	base := len(nums) / 2
// 	m := make(map[int]int, base)
//
// 	for i := range nums {
// 		if n, ok := m[nums[i]]; ok {
// 			if n+1 > base {
// 				return nums[i]
// 			}
// 			m[nums[i]] = n + 1
// 		} else {
// 			m[nums[i]] = 1
// 		}
// 	}
//
// 	return nums[0]
// }

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
