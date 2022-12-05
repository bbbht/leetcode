package main

// 92 reverse-linked-list-ii 2022-11-09 16:20:10

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	pre := dummy

	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	cur := pre.Next

	for i := left; i < right; i++ {
		tmp := cur.Next
		cur.Next = tmp.Next
		tmp.Next = pre.Next
		pre.Next = tmp
	}

	return dummy.Next
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {
	reverseBetween(createNode(8), 3, 7)
}
