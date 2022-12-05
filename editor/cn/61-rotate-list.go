package main

// 61 rotate-list 2022-10-08 16:57:15

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	tmp, length := head, 1
	for tmp.Next != nil {
		length++
		tmp = tmp.Next
	}

	// 收尾连接
	tmp.Next = head
	// 需要移动的步数，k可能超过长度
	step := length - k%length

	for i := 0; i < step; i++ {
		tmp = tmp.Next
	}

	head = tmp.Next
	tmp.Next = nil

	return head
}

// leetcode submit region end(Prohibit modification and deletion)

func main() {

}
