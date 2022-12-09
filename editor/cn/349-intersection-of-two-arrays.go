package main

// 349 intersection-of-two-arrays 2022-12-07 14:20:54

// leetcode submit region begin(Prohibit modification and deletion)
func intersection(nums1 []int, nums2 []int) []int {
	m, res := make(map[int]bool, len(nums1)), make([]int, 0, min349(len(nums1), len(nums2)))
	for i := range nums1 {
		m[nums1[i]] = false
	}
	for i := range nums2 {
		if _, ok := m[nums2[i]]; ok {
			if !m[nums2[i]] {
				res = append(res, nums2[i])
				m[nums2[i]] = true
			}
		}
	}
	return res
}

func min349(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
