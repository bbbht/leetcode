package main

// 删除链表的倒数第 N 个结点 2022-07-22 15:57:34

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fast, slow := head, head

	// 先让一个指针往前n步，然后要移除的指针开始同步走，直到先走的指针到头
	i := 1
	for ; fast.Next != nil; i++ {
		fast = fast.Next
		if i > n {
			slow = slow.Next
		}
	}

	if i <= n {
		return head.Next
	}

	slow.Next = slow.Next.Next

	return head
}

// leetcode submit region end(Prohibit modification and deletion)
