package main

// 169 majority-element 2022-11-29 10:32:27

// leetcode submit region begin(Prohibit modification and deletion)
func majorityElement(nums []int) int {
	base := len(nums) / 2
	m := make(map[int]int, base)

	for i := range nums {
		if n, ok := m[nums[i]]; ok {
			if n+1 > base {
				return nums[i]
			}
			m[nums[i]] = n + 1
		} else {
			m[nums[i]] = 1
		}
	}

	return nums[0]
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
