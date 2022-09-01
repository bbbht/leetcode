package main

// K 个一组翻转链表 2022-07-26 14:12:34

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	// 判断是否满足k，不足k则直接返回
	size := 0
	cur, next := head, head
	newHead := new(ListNode)
	tmp := newHead

	for next != nil {
		next = next.Next
		size++

		// 遍历达到 k，执行一次反转
		if size == k {
			size = 0
			node := reverseKNodeList(cur, k)

			// 反转后的这段链表，头变成尾，尾变成头
			tmp.Next = node
			tmp = cur

			// 重新定位
			cur = next
		}
	}

	return newHead.Next
}

// 就地逆置法，长度不足不处理
func reverseKNodeList(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	tmp, size := head, 0
	for tmp != nil && size < k {
		tmp = tmp.Next
		size++
	}

	if size < k {
		return head
	}

	cur, next := head, head.Next
	for ; next != nil && k > 1; k-- {
		cur.Next = next.Next
		next.Next = head

		head = next
		next = cur.Next
	}

	return head
}

// leetcode submit region end(Prohibit modification and deletion)
