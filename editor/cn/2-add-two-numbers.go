package main

// 两数相加 2022-07-06 15:38:42

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	tmp := head
	mark := 0
	for !(l1 == nil && l2 == nil) {
		a, b := 0, 0
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}

		tmp.Next = &ListNode{Val: a + b + mark}

		tmp = tmp.Next

		mark = tmp.Val / 10
		if mark > 0 {
			tmp.Next = &ListNode{Val: mark}
			tmp.Val = tmp.Val % 10
		}
	}

	return head.Next
}

// leetcode submit region end(Prohibit modification and deletion)
