package main

import (
	"fmt"
)

// 9 palindrome-number 2023-04-20 14:25:41

// leetcode submit region begin(Prohibit modification and deletion)
func isPalindrome(x int) bool {
	if x == 0 {
		return true
	}
	if x < 0 || x%10 == 0 {
		return false
	}
	nums := make([]int, 0, 32)
	for x > 0 {
		nums = append(nums, x%10)
		x /= 10
	}

	size := len(nums)

	for i, j := 0, size-1; i <= j; i, j = i+1, j-1 {
		if nums[i] != nums[j] {
			return false
		}
	}

	return true
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(isPalindrome(0))
	fmt.Println(isPalindrome(1))
	fmt.Println(isPalindrome(102))
}
