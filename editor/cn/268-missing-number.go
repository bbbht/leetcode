package main

// 268 missing-number 2022-12-01 16:18:41

// leetcode submit region begin(Prohibit modification and deletion)
func missingNumber(nums []int) int {
	sum, size := 0, len(nums)
	// 理论求和，然后减去实际和
	for i := range nums {
		sum += nums[i]
	}

	return size*(size+1)/2 - sum
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
