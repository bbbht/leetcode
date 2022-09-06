package main

// 275 h-index-ii 2022-09-02 16:43:44

// leetcode submit region begin(Prohibit modification and deletion)
func hIndex(citations []int) int {
	size := len(citations)
	mid, left, right := 0, 0, size-1

	for left <= right {
		mid = (left + right) / 2
		if citations[mid] == size-mid {
			return size - mid
		}
		if citations[mid] > size-mid {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return size - left
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
