package main
// 219 contains-duplicate-ii 2022-11-30 10:35:50

//leetcode submit region begin(Prohibit modification and deletion)
func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int, len(nums))
	for i := range nums {
		if j, ok := m[nums[i]]; ok && i-j <= k {
			return true
		}
		m[nums[i]] = i
	}
	return false
}
//leetcode submit region end(Prohibit modification and deletion)


func main() {

}