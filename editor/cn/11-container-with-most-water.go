package main

// 盛最多水的容器 2022-07-18 17:21:11

// leetcode submit region begin(Prohibit modification and deletion)
func maxArea(height []int) int {
	res := 0
	if len(height) <= 1 {
		return res
	}
	l, r := 0, len(height)-1
	for l < r {
		res = max11(res, min11(height[l], height[r])*(r-l))
		if height[l] < height[r] {
			l++
			continue
		}
		r--
	}

	return res
}

func max11(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min11(a, b int) int {
	if a > b {
		return b
	}

	return a
}

// func main()  {
// 	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
// }
// leetcode submit region end(Prohibit modification and deletion)
