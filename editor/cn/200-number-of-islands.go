package main

// 200 number-of-islands 2023-05-05 15:23:37

// leetcode submit region begin(Prohibit modification and deletion)
func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	result := 0
	rows, cols := len(grid), len(grid[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				maxIsland(i, j, &visited, grid)
				result++
			}
		}
	}

	return result
}

func maxIsland(x, y int, visited *[][]bool, grid [][]byte) {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) ||
		(*visited)[x][y] || grid[x][y] == '0' {
		return
	}
	(*visited)[x][y] = true

	maxIsland(x+1, y, visited, grid)
	maxIsland(x-1, y, visited, grid)
	maxIsland(x, y+1, visited, grid)
	maxIsland(x, y-1, visited, grid)
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
