package main

// 73 set-matrix-zeroes 2022-10-18 16:32:04

// leetcode submit region begin(Prohibit modification and deletion)
func setZeroes(matrix [][]int) {
	// 某列需要清零则为负，列为正
	rows, cols := len(matrix), len(matrix[0])
	rowBM, colBM := newBitMap73(rows), newBitMap73(cols)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == 0 {
				rowBM.add(i)
				colBM.add(j)
				// fmt.Printf("%d 行 %d 列 \n", i, j)
			}
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if rowBM.isExist(i) || colBM.isExist(j) {
				matrix[i][j] = 0
			}
		}
	}
}

type bitMap73 struct {
	bits []byte
	cap  int
}

func newBitMap73(cap int) *bitMap73 {
	bits := make([]byte, (cap>>3)+1)
	return &bitMap73{bits: bits, cap: cap}
}

func (b *bitMap73) add(num int) {
	b.bits[num>>3] |= 1 << (num & 0x07)
}

func (b *bitMap73) isExist(num int) bool {
	return b.bits[num>>3]&(1<<(num&0x07)) != 0
}

func (b *bitMap73) remove(num int) {
	b.bits[num>>3] &= ^(1 << (num & 0x07))
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
