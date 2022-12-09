package main

import (
	"math"
)

// 492 construct-the-rectangle 2022-12-09 15:45:11

// leetcode submit region begin(Prohibit modification and deletion)
func constructRectangle(area int) []int {
	w := int(math.Sqrt(float64(area)))

	for ; w > 0; w-- {
		if area%w == 0 {
			return []int{area / w, w}
		}
	}

	return nil
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
