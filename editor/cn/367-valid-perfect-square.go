package main

// 367 valid-perfect-square 2022-09-07 09:53:53

// leetcode submit region begin(Prohibit modification and deletion)
func isPerfectSquare(num int) bool {
	switch num % 10 {
	case 2, 3, 7, 8:
		return false
	}
	left, right, mid := 0, num, 0
	for left <= right {
		mid = (left + right) / 2
		tmp := mid * mid
		if tmp == num {
			return true
		}
		if tmp > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return false
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
