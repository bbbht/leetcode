package main

import (
	"fmt"
)

// 401 binary-watch 2022-12-08 11:06:26

// leetcode submit region begin(Prohibit modification and deletion)
var dict = []int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4, 1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5, 1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4}

func readBinaryWatch(turnedOn int) []string {
	if turnedOn >= 9 || turnedOn < 0 {
		return nil
	}
	var res []string
	for i := 0; i < 12; i++ {
		for j := 0; j < 60; j++ {
			if turnedOn == dict[i]+dict[j] {
				time := fmt.Sprintf("%d:%02d", i, j)
				res = append(res, time)
			}
		}
	}

	return res
}

// leetcode submit region end(Prohibit modification and deletion)

func countBits401(n int) []int {
	res := make([]int, n+1)
	for i := 1; i <= n; i++ {
		res[i] = res[i&(i-1)] + 1
	}

	return res
}

func main() {
	fmt.Printf("%#v", countBits401(60))
}
