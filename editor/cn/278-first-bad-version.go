package main

import (
	"fmt"
)

// 278 first-bad-version 2022-09-05 15:31:12

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	mid, left, right := 0, 0, n
	for left < right {
		mid = (left + right) / 2
		if isBadVersion(mid + 1) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return right + 1
}

// leetcode submit region end(Prohibit modification and deletion)
func isBadVersion(version int) bool {
	if version >= 4 {
		return true
	}
	return false
}

func main() {
	fmt.Println(firstBadVersion(5))
}
