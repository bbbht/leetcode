package main

// 496 next-greater-element-i 2022-12-09 10:49:43

// leetcode submit region begin(Prohibit modification and deletion)
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m, tmp := make(map[int]int), make([]int, 0)
	for _, v := range nums2 {
		for len(tmp) > 0 && tmp[len(tmp)-1] < v {
			m[tmp[len(tmp)-1]] = v
			tmp = tmp[:len(tmp)-1]
		}
		tmp = append(tmp, v)
	}
	var ans []int
	for _, v := range nums1 {
		val, ok := m[v]
		if !ok {
			val = -1
		}
		ans = append(ans, val)
	}
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
