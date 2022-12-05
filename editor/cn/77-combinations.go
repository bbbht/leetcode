package main

import (
	"fmt"
)

// 77 combinations 2022-11-03 15:49:01

// leetcode submit region begin(Prohibit modification and deletion)
func combine(n int, k int) [][]int {
	arr, result := make([]int, 0, k), make([][]int, 0, c77(n, k))

	dfs77(n, k, 1, arr, &result)

	return result
}

func dfs77(n, k, start int, arr []int, result *[][]int) {
	if len(arr) == k {
		tmp := make([]int, k)
		copy(tmp, arr)
		*result = append(*result, tmp)

		return
	}

	for i := start; i <= n; i++ {
		arr = append(arr, i)
		dfs77(n,k,i+1, arr, result)
		arr = arr[:len(arr)-1]
	}
}

func c77(n, k int) int {
	if n == k {
		return 1
	}

	return factorial77(n) / (factorial77(k) * factorial77(n-k))
}

func factorial77(n int) (res int) {
	res = 1
	for i := 2; i <= n; i++ {
		res *= i
	}

	return
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(combine(3,1))
}
