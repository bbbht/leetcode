package main

// 167 two-sum-ii-input-array-is-sorted 2022-09-01 17:09:20

// leetcode submit region begin(Prohibit modification and deletion)
func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		if numbers[left]+numbers[right] == target {
			return []int{left + 1, right + 1}
		}
		if numbers[left]+numbers[right] > target {
			right--
		} else {
			left++
		}
	}

	return nil
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
