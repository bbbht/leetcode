package main

// 203 remove-linked-list-elements 2022-11-29 15:35:08

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	cur := dummy
	for head != nil {
		if head.Val == val {
			cur.Next, head = head.Next, head.Next
		} else {
			cur, head = head, head.Next
		}
	}

	return dummy.Next
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
