package main

// 88 merge-sorted-array 2023-05-07 16:53:50

// leetcode submit region begin(Prohibit modification and deletion)
func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := m + n - 1; m > 0 && n > 0; i-- {
		if nums1[m-1] >= nums2[n-1] {
			nums1[i] = nums1[m-1]
			m--
		} else {
			nums1[i] = nums2[n-1]
			n--
		}
	}

	// 如果nums2没有被放置到nums1种，那直接挨个放置即可
	// nums1如果有剩余不需要处理，已经在位置
	for ; n > 0; n-- {
		nums1[n-1] = nums2[n-1]
	}
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
