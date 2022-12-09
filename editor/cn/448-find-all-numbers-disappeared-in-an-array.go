package main

// 448 find-all-numbers-disappeared-in-an-array 2022-12-08 16:54:02

// leetcode submit region begin(Prohibit modification and deletion)
func findDisappearedNumbers(nums []int) []int {
	n, res := len(nums), []int{}
	m := make(map[int]struct{})
	for i := range nums {
		m[nums[i]] = struct{}{}
	}
	for i := 1; i <= n; i++ {
		if _, ok := m[i]; !ok {
			res = append(res, i)
		}
	}

	return res
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
