package main

// 85 maximal-rectangle 2023-05-07 17:09:16

// leetcode submit region begin(Prohibit modification and deletion)
func maximalRectangle(matrix [][]byte) int {
	var max int
	rows := len(matrix)
	if rows == 0 {
		return max
	}

	cols := len(matrix[0])
	nums := make([]int, cols)

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if matrix[row][col] != '0' {
				nums[col] = nums[col] + 1
			} else {
				nums[col] = 0
			}
		}

		// 求出每一行作为底的直方图高度，复用84的方法求最大值
		maxArea := largestRectangleArea84(nums)
		max = max85(max, maxArea)
	}

	return max
}

func largestRectangleArea84(heights []int) int {
	heights = append(heights, 0)
	res, stack := 0, make([]int, 0, len(heights)/8+1)

	for i := range heights {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}

			res = max85(res, height*width)
		}

		// 存储的数据为索引
		stack = append(stack, i)
	}

	return res
}

func max85(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
