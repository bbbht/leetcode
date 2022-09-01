package main

// Z 字形变换 2022-07-15 17:30:55

// leetcode submit region begin(Prohibit modification and deletion)
// 通过利用重复的规律，判断每个字符所在行，然后再按行拼接
func convert(s string, numRows int) string {
	// 最小重复单元的容量
	if numRows <= 1 {
		return s
	}

	length := len(s)

	size := 2*numRows-2

	ns := make([][]uint8, numRows)
	for i := range ns {
		ns[i] = make([]uint8, 0, length/size * 2)
	}

	for i := 0; i < length; i++ {
		// 每个循环中的索引
		row := i % size
		if row > numRows - 1 {
			row = size - row
		}

		ns[row] = append(ns[row], s[i])
	}

	res := ""
	for i := range ns {
		res += string(ns[i])
	}
	return res
}

// func convert(s string, numRows int) string {
// 	// 最小重复单元的容量
// 	if numRows <= 1 {
// 		return s
// 	}
//
// 	length := len(s)
//
// 	size := 2*numRows - 2
//
// 	res := make([]uint8, 0, length)
//
// 	// 计算公式，使用间隔规律
// 	// 第一行 循环容量 size
// 	// 第二行，size-2/2
// 	// 第三行，siez-4/4
// 	// 直到 2/size-2
// 	// 最后一行 size
// 	for i := 0; i < numRows; i++ {
// 		for j := range s {
// 			// 计算
// 			if  {
// 				res = append(res,s[j])
// 			}
// 		}
// 	}
//
// 	return string(res)
// }

// func main()  {
// 	fmt.Println(convert("ABCDEF", 5))
// }
// leetcode submit region end(Prohibit modification and deletion)
