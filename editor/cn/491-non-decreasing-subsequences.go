package main

// 491 non-decreasing-subsequences 2023-05-05 16:11:30

// leetcode submit region begin(Prohibit modification and deletion)
func findSubsequences(nums []int) [][]int {
	var res [][]int
	var path []int

	bt491(nums, 0, path, &res)

	return res
}

func bt491(nums []int, start int, path []int, res *[][]int) {
	if len(path) >= 2 {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
	}

	var used [201]bool
	for i := start; i < len(nums); i++ {
		if len(path) > 0 && nums[i] < path[len(path)-1] || used[nums[i]+100] {
			continue
		}

		used[nums[i]+100] = true

		path = append(path, nums[i])

		bt491(nums, i+1, path, res)

		path = path[:len(path)-1]
	}
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
