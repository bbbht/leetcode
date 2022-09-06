package main

// 350 intersection-of-two-arrays-ii 2022-09-06 13:57:19

// leetcode submit region begin(Prohibit modification and deletion)
func intersect(nums1 []int, nums2 []int) []int {
	var res []int

	m := make(map[int]int, len(nums1))
	for _, num := range nums1 {
		if _, ok := m[num]; ok {
			m[num] += 1
		} else {
			m[num] = 1
		}
	}

	for _, num := range nums2 {
		if c, ok := m[num]; ok && c > 0 {
			m[num] -= 1
			res = append(res, num)
		}
	}

	return res
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
