package main

// 合并两个有序链表 2022-07-25 16:48:24

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	tmp := head
	for l1 != nil || l2 != nil {
		if l1 == nil {
			tmp.Next = l2
			return head.Next
		}

		if l2 == nil {
			tmp.Next = l1
			return head.Next
		}

		if l1.Val < l2.Val {
			tmp.Next = l1
			tmp = tmp.Next
			l1 = l1.Next
		} else {
			tmp.Next = l2
			tmp = tmp.Next
			l2 = l2.Next
		}
	}

	return head.Next
}

// leetcode submit region end(Prohibit modification and deletion)
