package main

import (
	"strconv"
)

// 412 fizz-buzz 2022-12-08 15:35:04

// leetcode submit region begin(Prohibit modification and deletion)
func fizzBuzz(n int) []string {
	tmp, res := "", make([]string, 0, n)
	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				tmp = "FizzBuzz"
			} else {
				tmp = "Fizz"
			}
		} else if i%5 == 0 {
			tmp = "Buzz"
		} else {
			tmp = strconv.Itoa(i)
		}
		res = append(res, tmp)
	}
	return res
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
