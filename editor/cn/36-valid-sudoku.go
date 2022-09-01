package main

// 有效的数独 2022-07-28 15:53:49

// leetcode submit region begin(Prohibit modification and deletion)
func isValidSudoku(board [][]byte) bool {
	rows, cols, grids := [9][9]byte{}, [9][9]byte{}, [9][9]byte{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			b := board[i][j] - '1'

			// 第 i 行，最多 9 个数
			if rows[i][b] != 0 {
				return false
			}
			rows[i][b] = 1

			// 第 j 列，最多 9 个数
			if cols[j][b] != 0 {
				return false
			}
			cols[j][b] = 1

			// 第几个格子，从左到右，从上到下
			g := (i/3)*3 + j/3
			if grids[g][b] != 0 {
				return false
			}
			grids[g][b] = 1
		}
	}

	return true
}

// leetcode submit region end(Prohibit modification and deletion)
