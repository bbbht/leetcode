package main

// 合并K个升序链表 2022-07-26 10:34:27

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	return merge23(0, len(lists)-1, lists)
}

func merge23(l, r int, lists []*ListNode) *ListNode {
	if l == r {
		return lists[l]
	}
	if l > r {
		return nil
	}

	mid := (l + r) / 2

	// 分治，通过不断的拆分数组，实现两两合并，直到合二为一
	return mergeTwoLists23(merge23(l, mid, lists), merge23(mid+1, r, lists))
}

func mergeTwoLists23(l1, l2 *ListNode) *ListNode {
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
