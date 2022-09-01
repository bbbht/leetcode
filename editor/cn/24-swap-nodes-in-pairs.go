package main

// 两两交换链表中的节点 2022-07-26 11:04:36

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	newHead := new(ListNode)
	newHead.Next = head

	pre, cur := newHead, head

	// 迭代修改
	for cur != nil && cur.Next != nil {
		//
		next := cur.Next

		cur.Next = cur.Next.Next
		next.Next = cur
		pre.Next = next

		pre = cur
		cur = cur.Next
	}

	return newHead.Next
}

// leetcode submit region end(Prohibit modification and deletion)
