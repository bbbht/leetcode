package main

// 寻找两个正序数组的中位数 2022-07-07 15:25:29

// leetcode submit region begin(Prohibit modification and deletion)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	total := len(nums1) + len(nums2)
	if total == 0 {
		return 0
	}
	if total%2 == 0 {
		return (float64(getKthElement(total/2, nums1, nums2)) + float64(getKthElement(total/2+1, nums1, nums2))) / 2
	}
	return float64(getKthElement(total/2+1, nums1, nums2))
}

// 获取指定第k小的数字
func getKthElement(k int, nums1, nums2 []int) int {
	len1, len2 := len(nums1), len(nums2)
	// 数组可用的所以起始位置
	idx1, idx2 := 0, 0

	for {
		// 判断是否有数组越界了
		// 前置，避免空数组
		if idx1 == len1 {
			return nums2[idx2+k-1]
		}
		if idx2 == len2 {
			return nums1[idx1+k-1]
		}

		if k == 1 {
			return min(nums1[idx1], nums2[idx2])
		}

		// 计算出本次判断中的起始位置
		i1 := min(len1, idx1+k/2) - 1
		i2 := min(len2, idx2+k/2) - 1
		// 比较两个位点值的大小
		if nums1[i1] <= nums2[i2] {
			// k值减去此轮循环中被忽略掉的数量
			k = k - (i1 - idx1 + 1)
			// 移动数组起始位置，调整可用范围
			idx1 = i1 + 1
		} else {
			k = k - (i2 - idx2 + 1)
			idx2 = i2 + 1
		}
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

// func main()  {
// 	fmt.Println(findMedianSortedArrays([]int{}, []int{1}))
// }
// leetcode submit region end(Prohibit modification and deletion)
