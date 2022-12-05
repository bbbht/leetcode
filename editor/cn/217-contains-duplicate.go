package main

// 217 contains-duplicate 2022-11-30 09:54:06

// leetcode submit region begin(Prohibit modification and deletion)
func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{}, len(nums))
	for i := range nums {
		if _, ok := m[nums[i]]; ok {
			return true
		} else {
			m[nums[i]] = struct{}{}
		}
	}

	return false
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
