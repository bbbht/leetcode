package main

// 82 remove-duplicates-from-sorted-list-ii 2022-11-04 16:01:42

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	slow, fast := new(ListNode), head
	slow.Next = head
	newHead := slow

	for fast != nil && fast.Next != nil {
		dup := false
		// 判断当前值是否重复
		for fast.Next != nil && fast.Val == fast.Next.Val {
			dup = true
			fast = fast.Next
		}

		// 如果重复了，那么就要跳过所有重复的节点
		if dup {
			// 如果重复，那么slow就需要变更next了
			// 修改newHead中此处的值
			slow.Next = fast.Next
		} else {
			// 移动slow
			slow = fast
		}

		fast = fast.Next
	}

	return newHead.Next
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
