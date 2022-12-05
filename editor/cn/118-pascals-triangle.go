package main

// 118 pascals-triangle 2022-11-24 10:04:56

// leetcode submit region begin(Prohibit modification and deletion)
func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	result[0] = []int{1}
	for i := 1; i < numRows; i++ {
		result[i] = make([]int, i+1)
		for j := range result[i] {
			if j == 0 || j == i {
				result[i][j] = 1
			} else {
				result[i][j] = result[i-1][j-1] + result[i-1][j]
			}
		}
	}

	return result
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
