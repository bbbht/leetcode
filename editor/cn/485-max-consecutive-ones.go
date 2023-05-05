package main

// 485 max-consecutive-ones 2022-12-12 10:03:48

// leetcode submit region begin(Prohibit modification and deletion)
func findMaxConsecutiveOnes(nums []int) int {
	n, max := 0, 0
	for i := range nums {
		if nums[i] == 0 {
			n = 0
			continue
		}

		n += 1
		if n > max {
			max = n
		}
	}

	return max
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
