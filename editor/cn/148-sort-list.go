package main

// 148 sort-list 2023-05-05 17:54:42

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	size := 0
	tmp := head
	for tmp != nil {
		size++
		tmp = tmp.Next
	}
	if size <= 1 {
		return head
	}

	// 通过快慢指针，找到链表的中间位置，一分为二
	fast, slow, pre := head, head, head
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	pre.Next = nil // 截断前一段链表

	return merge148(sortList(head), sortList(slow))
}

func merge148(l1, l2 *ListNode) *ListNode {
	var dummy = new(ListNode)
	cur := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}

		cur = cur.Next
	}

	if l1 == nil {
		cur.Next = l2
	}
	if l2 == nil {
		cur.Next = l1
	}

	return dummy.Next
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
