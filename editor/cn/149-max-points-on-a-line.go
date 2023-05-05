package main

import (
	"fmt"
)

// 149 max-points-on-a-line 2023-04-10 13:25:02

// leetcode submit region begin(Prohibit modification and deletion)
func maxPoints(points [][]int) int {
	l := len(points)
	if l <= 2 {
		return l
	}
	max := 0

	for i := 0; i < l; i++ {
		same := 1 // 一样的点计数
		for j := i + 1; j < l; j++ {
			count := 0 // 在一条线上的计数
			if points[i][0] == points[j][0] && points[i][1] == points[j][1] {
				same++
				continue
			}
			count++
			xDiff, yDiff := points[i][0]-points[j][0], points[i][1]-points[j][1]
			for k := j + 1; k < l; k++ { // 其它点是否在 i,j 直线上
				// 将除法比较结果转为乘法比较结果，防止精度问题
				if xDiff*(points[i][1]-points[k][1]) == yDiff*(points[i][0]-points[k][0]) {
					count++
				}
			}
			if max < same+count {
				max = same + count
			}
		}
	}

	return max
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Println(maxPoints([][]int{
		{1, 1},
		{2, 2},
		{3, 3},
	}))
}
