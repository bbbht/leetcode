package main

import (
	"fmt"
)

// 374 guess-number-higher-or-lower 2022-09-06 16:24:49

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	left, right, mid := 0, n, 0
	for left <= right {
		mid = (left + right) / 2
		switch guess(mid) {
		case 0:
			return mid
		case -1:
			right = mid - 1
		case 1:
			left = mid + 1
		}
	}

	return -1
}

// leetcode submit region end(Prohibit modification and deletion)

func guess(num int) int {
	n := 2
	if num > n {
		return -1
	}
	if num < n {
		return 1
	}

	return 0
}

func main() {
	fmt.Println(guessNumber(1))
}
