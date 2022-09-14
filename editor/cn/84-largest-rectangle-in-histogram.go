package main

// 84 largest-rectangle-in-histogram 2022-09-09 16:47:54

// leetcode submit region begin(Prohibit modification and deletion)
func largestRectangleArea(heights []int) int {
	heights = append(heights, 0)
	res, stack := 0, make([]int, 0, len(heights)/8+1)

	for i := range heights {
		// 如果当前值小于等于栈顶的值，那么就触发计算，计算后移除栈顶，直至栈顶值小于当前值
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]

			// 如果栈已清空，说明当前值小于前面所有值，宽度可以是索引值
			// 如果非空，则需要计算一下宽度，存储的索引并不是连续的，而是大于等于当前值的，所以高度可以直接使用当前值
			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}

			res = max84(res, height*width)
		}

		// 存储的数据为索引
		stack = append(stack, i)
	}

	return res
}

func max84(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
